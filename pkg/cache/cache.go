package cache

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/nktauserum/crawler-service/common"
)

type CacheInterface interface {
	Get(url string) (page, bool)
	Set(url string, new_page common.Page)
}

var (
	once  sync.Once
	cache CacheInterface

	// TTL полгода
	ttl = 6 * 30 * (24 * time.Hour)
)

type Cache struct {
	memory map[string]page
	mu     sync.Mutex
}

type page struct {
	common.Page
	TTL time.Time
}

func NewCache() CacheInterface {
	once.Do(func() {
		_ = godotenv.Load(".env")
		cacheType := os.Getenv("CACHE_TYPE")

		if cacheType == "redis" {
			redisAddr := os.Getenv("REDIS_ADDR")
			redisPassword := os.Getenv("REDIS_PASSWORD")

			if redisAddr == "" {
				redisAddr = "localhost:6379" // по умолчанию
			}

			cache = NewRedisCache(redisAddr, redisPassword, 0)
		} else {
			log.Println("Using in-memory cache by default")
			cache = &Cache{memory: make(map[string]page)}
		}
	})

	return cache
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
