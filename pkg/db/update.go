package db

func UpdateText(s *Storage, uuid, text string) error {
	_, err := s.DB.Exec(
		"UPDATE crawl_tasks SET result = ?, status = ?, completed_at = CURRENT_TIMESTAMP WHERE uuid = ?",
		text, "completed", uuid,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStatus(s *Storage, uuid, status string) error {
	_, err := s.DB.Exec(
		"UPDATE crawl_tasks SET status = ? WHERE uuid = ?",
		status, uuid,
	)
	if err != nil {
		return err
	}

	return nil
}
