package worker

import (
	"context"
	"github.com/nktauserum/crawler-service/pkg/storage"
	"time"

	"github.com/nktauserum/crawler-service/common"
	"github.com/nktauserum/crawler-service/pkg/crawler"
)

func Process(task common.Task) {
	start_time := time.Now()

	page, err := crawler.GetContent(context.Background(), task.URL)
	if err != nil {
		task.Status = "failed"
	} else {
		task.Status = "done"
		task.Result = page
	}

	memory := storage.GetInMemoryStorage()
	task.Time = time.Since(start_time).String()
	_ = memory.Set(task)
}
