package task

import "github.com/nktauserum/crawler-service/pkg/db"

func CompleteTask(s *db.Storage, uuid, text string) error {
	_, err := s.DB.Exec(`
 		UPDATE crawl_tasks 
 		SET status = 'completed',
 		    completed_at = CURRENT_TIMESTAMP,
			result = ?
 		WHERE uuid = ?`, text, uuid)
	return err
}
