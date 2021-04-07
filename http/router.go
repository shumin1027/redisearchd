package http

import (
	"context"
	"fmt"
	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.xtc.home/xtc/redisearchd/conn"
	_ "gitlab.xtc.home/xtc/redisearchd/docs"
	"gitlab.xtc.home/xtc/redisearchd/internal/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var app *fiber.App
var routers = make(map[string]Router)

type Router interface {
	Route()
}

func init() {
	cfg := fiber.Config{
		JSONEncoder: json.Marshal,
	}

	app = fiber.New(cfg)
	Route(app)
}

// 注册路由
func Route(a *fiber.App) *fiber.App {

	// Recover Middleware
	app.Use(recover.New())

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
			"name":    "redisearchd",
			"version": "v1.0",
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
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	fmt.Println("Closeing redis conn...")
	conn.Close()

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
