package app

import (
	"fmt"

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

	router.POST("/crawl", handlers.Crawl)

	return router.Run(fmt.Sprintf(":%d", app.port))
}
