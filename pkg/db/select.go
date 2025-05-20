package db

import (
	"database/sql"
	"log"

	"github.com/nktauserum/crawler-service/common"
)

func SelectTask(s *Storage, uuid string) (*common.Task, error) {
	var task common.Task
	var result sql.NullString

	err := s.DB.QueryRow(
		"SELECT uuid, url, status, result FROM crawl_tasks WHERE uuid = ?",
		uuid,
	).Scan(&task.UUID, &task.URL, &task.Status, &result)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Ошибка при поиске выражения %s: %v", uuid, err)
		}
		return nil, err
	}

	task.Result = result.String
	return &task, nil
}
