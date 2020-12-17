package http

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.xtc.home/xtc/redisearchd/docs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var server *http.Server
var engine *gin.Engine
var routers = make(map[string]Router)

type Router interface {
	Route()
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.Default()
	Route(engine)
}

// 注册路由
func Route(e *gin.Engine) *gin.Engine {
	// Duration Middleware
	e.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		duration := end.Sub(start)
		c.Writer.Header().Add("X-Request-Duration", duration.String())
	})

	// swagger api docs url: /swagger/index.html
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// version
	e.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "redisearchd",
			"version": "v1.0",
		})
	})

	// ping
	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	routers["index"] = NewIndexRouter(e.Group("/indexes"))
	routers["search"] = NewSearchRouter(e.Group("/search"))
	routers["doc"] = NewDocRouter(e.Group("/docs"))

	for _, v := range routers {
		v.Route()
	}

	return e
}

// RediSearchs API
// @title RediSearchs API
// @contact.name shumin
// @contact.email shumin@compubiq.com
// @version 1.0
// @Description RediSearchs API
// @host localhost:8080
// @BasePath /api/v1
func Start(addr string) {
	server = &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	log.Printf("listen: %s\n", addr)
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	graceful(server)
}

func graceful(server *http.Server) {
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
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
