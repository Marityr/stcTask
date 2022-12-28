package repository

import "github.com/go-redis/redis"

type DataList interface {
	GetKey(key string) (string, error)
}

type Repository struct {
	DataList
}

func NewReposiroty(db *redis.Client) *Repository {
	return &Repository{
		DataList: NewDataListRedis(db),
	}
}
