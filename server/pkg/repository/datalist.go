package repository

import (
	"log"

	"github.com/go-redis/redis"
)

type DataListRedis struct {
	client *redis.Client
}

func NewDataListRedis(db *redis.Client) *DataListRedis {
	return &DataListRedis{client: db}
}

func (dl *DataListRedis) GetKey(key string) (string, error) {
	val, err := dl.client.Get("name").Result()
	if err != nil {
		return "", err
	}
	log.Println("###############")
	log.Println(val)
	return val, nil
}
