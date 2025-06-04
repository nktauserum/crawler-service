package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nktauserum/crawler-service/common"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) Get(url string) (page, bool) {
	val, err := r.client.Get(r.ctx, url).Result()
	if err != nil {
		return page{}, false
	}

	var p page
	err = json.Unmarshal([]byte(val), &p)
	if err != nil {
		return page{}, false
	}

	return p, true
}

func (r *RedisCache) Set(url string, new_page common.Page) {
	p := page{
		Page: new_page,
		TTL:  time.Now().Add(ttl),
	}

	data, err := json.Marshal(p)
	if err != nil {
		return
	}

	r.client.Set(r.ctx, url, data, ttl)
}
