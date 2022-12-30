package repository

import (
	"context"
	"server/pkg/logging"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func NewRedisDB(v *viper.Viper) (*redis.Client, error) {
	logger := logging.GetLooger()

	ctx := context.Background()

	db := redis.NewClient(&redis.Options{
		Addr:     v.GetString("db.host"),
		Password: v.GetString("db.password"),
		DB:       v.GetInt("db.database"),
	})

	_, err := db.Ping(ctx).Result()
	if err != nil {
		logger.Fatal(err)
		return db, err
	}

	return db, nil
}
