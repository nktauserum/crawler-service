package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nktauserum/crawler-service/pkg/db"
)

func Result(c *gin.Context) {
	uuid := c.Query("id")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "needs url query id",
		})
		return
	}

	s := db.GetStorage()

	task, err := db.SelectTask(s, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}
