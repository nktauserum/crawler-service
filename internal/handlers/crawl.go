package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nktauserum/crawler-service/common"
	"github.com/nktauserum/crawler-service/internal/worker"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	URL         string `json:"url"`
	IncludeHTML bool   `json:"include_html"`
}

func Crawl(c *gin.Context) {
	start_time := time.Now()

	request_body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	request := new(request)
	err = json.Unmarshal(request_body, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if request.URL == "" {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "empty URL field in request",
		})
		return
	}

	id := uuid.New().String()
	task := common.Task{ID: id, URL: request.URL, Status: "pending", Time: time.Since(start_time).String()}
	go worker.Process(task)

	c.JSON(http.StatusOK, task)
}
