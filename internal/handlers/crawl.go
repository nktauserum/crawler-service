package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nktauserum/crawler-service/common"
	"github.com/nktauserum/crawler-service/pkg/crawler"
)

type request struct {
	URL         string `json:"url"`
	IncludeHTML bool   `json:"include_html"`
}

type response struct {
	common.Page
	Time string `json:"time"`
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

	page, err := crawler.GetContent(context.Background(), request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !request.IncludeHTML {
		page.HTML = ""
	}

	crawler_response := response{
		Page: page,
		Time: fmt.Sprint(time.Since(start_time).Seconds()),
	}

	c.JSON(http.StatusOK, crawler_response)
}
