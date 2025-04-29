package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nktauserum/crawler-service/internal/handlers"
)

type application struct {
	port int
}

func NewApplication(port int) *application {
	return &application{port: port}
}

func (app *application) Run() error {
	router := gin.Default()

	// Simple health check
	router.GET("/", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	authorized := router.Group("/")
	authorized.Use(handlers.CheckAPIToken())
	{
		authorized.POST("/crawl", handlers.Crawl)
	}

	return router.Run(fmt.Sprintf(":%d", app.port))
}
