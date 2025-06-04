package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nktauserum/crawler-service/common"
	storage "github.com/nktauserum/crawler-service/pkg/storage"
	"io"
	"net/http"
)

func Task(c *gin.Context) {
	memory := storage.GetInMemoryStorage()

	request_body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	request := new(common.GetTaskRequest)
	err = json.Unmarshal(request_body, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := memory.Get(request.ID)
	if err != nil {
		if errors.Is(err, storage.ErrTaskNotExists) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("task with ID %s not exists", request.ID),
			})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}
