package db

import "github.com/google/uuid"

func NewTask(s *Storage, url string) (string, error) {
	id := uuid.New().String()

	_, err := s.DB.Exec("INSERT INTO crawl_tasks (uuid, url, status) VALUES (?, ?, ?)", id, url, "pending")
	if err != nil {
		return "", err
	}

	return id, nil
}
