package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nktauserum/crawler-service/pkg/db"
)

/*
Алгоритм обработки запроса:
1. Валидируем входящие данные
2. Создаём новый запрос и присваиваем ему ID
3. Заносим его в очередь в базе данных
4. Процесс-воркер забирает запрос и обрабатывает
*/

type request struct {
	URL         string `json:"url"`
	IncludeHTML bool   `json:"include_html"`
}

// type response struct {
// 	common.Page
// 	Time string `json:"time"`
// }

func Crawl(c *gin.Context) {
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

	// Валидация данных
	if request.URL == "" {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "empty URL field in request",
		})
		return
	}

	s := db.GetStorage()

	uuid, err := db.NewTask(s, request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": uuid,
	})
}

/*
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
*/
