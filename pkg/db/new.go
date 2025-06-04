package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var storage *Storage
var once sync.Once

func GetStorage() *Storage {
	var err error
	once.Do(func() {
		q, err_in := NewStorage("sqlite.db")
		if err_in != nil {
			err = err_in
		}

		storage = q
	})

	if err != nil {
		panic("error setting up queue: " + err.Error())
	}

	return storage
}

type Storage struct {
	DB *sql.DB
}

// NewQueue создает новую очередь с SQLite
func NewStorage(dbPath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия базы данных: %w", err)
	}

	db.Exec("PRAGMA journal_mode = WAL;")

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	q := &Storage{DB: db}
	if err := q.initTables(); err != nil {
		return nil, fmt.Errorf("ошибка инициализации таблиц: %w", err)
	}

	return q, nil
}

// Close закрывает соединение с базой данных
func (q *Storage) Close() error {
	return q.DB.Close()
}

// initTables создает необходимые таблицы, если они не существуют
func (q *Storage) initTables() error {
	// Создаем таблицу для задач
	_, err := q.DB.Exec(`
		CREATE TABLE IF NOT EXISTS crawl_tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL,
			status TEXT NOT NULL,
			result TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			completed_at TIMESTAMP
		)
	`)

	return err
}
