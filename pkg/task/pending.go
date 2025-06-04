package task

import (
	"database/sql"
	"log"

	"github.com/nktauserum/crawler-service/common"
	"github.com/nktauserum/crawler-service/pkg/db"
)

func PendingTask(s *db.Storage) (*common.Task, error) {
	var task common.Task

	err := s.DB.QueryRow(
		"SELECT uuid, url FROM crawl_tasks WHERE status = 'pending'",
	).Scan(&task.UUID, &task.URL)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Ошибка при поиске свободной задачи: %v", err)
		}
		return nil, err
	}

	err = db.UpdateStatus(s, task.UUID, "progress")
	if err != nil {
		return nil, err
	}

	return &task, nil
}
