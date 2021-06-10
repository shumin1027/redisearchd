package http

import (
	"gitlab.xtc.home/xtc/redisearchd/conn/search"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	self "gitlab.xtc.home/xtc/redisearchd/app"
	_ "gitlab.xtc.home/xtc/redisearchd/docs"
	"gitlab.xtc.home/xtc/redisearchd/pkg/json"
)

var app *fiber.App
var routers = make(map[string]Router)

type Router interface {
	Route()
}

func init() {
	cfg := fiber.Config{
		Prefork:     true,
		JSONEncoder: json.Marshal,
		ReadTimeout: 10 * time.Second,
	}

	app = fiber.New(cfg)

	Route(app)
}

// 注册路由
func Route(a *fiber.App) *fiber.App {

	// Recover Middleware
	app.Use(recover.New())

	// Log Middleware
	app.Use(logger.New(logger.Config{
		Format:       "${time} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "2006/01/02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
	}))

	// Duration Middleware
	a.Use(func(c *fiber.Ctx) error {
		var start, stop time.Time
		start = time.Now()
		if chainErr := c.Next(); chainErr != nil {
			return chainErr
		}
		stop = time.Now()
		duration := stop.Sub(start).Round(time.Nanosecond)
		c.Set("X-Request-Duration", duration.String())
		return nil
	})

	// swagger api docs url: /swagger/index.html
	app.Get("/swagger/*", swagger.Handler) // default

	// version
	a.Get("/version", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"name":    self.Name,
			"version": self.Version,
		})
	})

	// ping
	a.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("pong")
	})

	// stack
	a.Get("/stack", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(app.Stack())
	})

	routers["doc"] = NewDocRouter(a.Group("/docs"))
	routers["index"] = NewIndexRouter(a.Group("/indexes"))
	routers["search"] = NewSearchRouter(a.Group("/search"))

	for _, v := range routers {
		v.Route()
	}

	return app
}

// RediSearchs API
// @title RediSearchs API
// @contact.name shumin
// @contact.email shumin@compubiq.com
// @version 1.0
// @Description RediSearchs API
// @host localhost:8080
// @BasePath /
func Start(addr string) {
	go func() {
		// service connections
		if err := app.Listen(addr); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	graceful(app)
}

func graceful(app *fiber.App) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server ...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("server shutdown:", err)
	}

	log.Println("running cleanup tasks...")
	// Your cleanup tasks go heremak
	log.Println("closeing redis conn...")
	if err := search.Close(); err != nil {
		log.Fatal("close redis conn:", err)
	}

	log.Println("server exited")
}
