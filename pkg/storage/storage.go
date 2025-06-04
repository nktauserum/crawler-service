package storage

import (
	"fmt"
	"github.com/nktauserum/crawler-service/common"
	"sync"
)

var (
	ErrTaskNotExists = fmt.Errorf("task with given uuid is not exists")
	ErrEmptyIDField  = fmt.Errorf("field id is empty")
)

var (
	once       sync.Once
	inMstorage *InMemoryStorage
)

// Хранит в себе задачи, в данный момент обрабатывающиеся
// Может отдавать задачу по ID: статус, результат (если есть)
// Различные горутины могут перезаписывать задачи с определенным статусом
type Storage interface {
	Get(string) (common.Task, error)
	Set(common.Task) error
}

type InMemoryStorage struct {
	s  map[string]common.Task
	mu sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{s: make(map[string]common.Task)}
}

func GetInMemoryStorage() *InMemoryStorage {
	once.Do(func() {
		inMstorage = NewInMemoryStorage()
	})

	return inMstorage
}

func (i *InMemoryStorage) Get(uuid string) (common.Task, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	task, exists := i.s[uuid]
	if !exists {
		return common.Task{}, ErrTaskNotExists
	}

	return task, nil
}

func (i *InMemoryStorage) Set(task common.Task) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if task.ID == "" {
		return ErrEmptyIDField
	}

	i.s[task.ID] = task

	return nil
}
