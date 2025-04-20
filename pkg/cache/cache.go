package cache

import (
	"sync"
	"time"

	"github.com/nktauserum/crawler-service/common"
)

var (
	once  sync.Once
	cache Cache

	ttl = 5 * time.Minute
)

type Cache struct {
	memory map[string]page
	mu     sync.Mutex
}

type page struct {
	common.Page
	TTL time.Time
}

func NewCache() *Cache {
	once.Do(func() {
		cache = Cache{memory: make(map[string]page)}
	})

	return &cache
}

// Получаем запись из кэша. Возвращает страницу и флаг наличия
func (c *Cache) Get(url string) (page, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	page_struct, exists := c.memory[url]
	if !exists {
		return page{}, exists
	}

	if time.Now().After(page_struct.TTL) {
		return page{}, false
	}

	return page_struct, exists
}

func (c *Cache) Set(url string, new_page common.Page) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.memory[url] = page{Page: new_page, TTL: time.Now().Add(ttl)}
}
