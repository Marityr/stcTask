package repository

import (
	"stcTask/server/pkg/logging"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func NewRedisDB(v *viper.Viper) (*redis.Client, error) {
	logger := logging.GetLooger()

	db := redis.NewClient(&redis.Options{
		Addr:     v.GetString("db.host"),
		Password: v.GetString("db.pasword"),
		DB:       v.GetInt("db.database"),
	})

	_, err := db.Ping().Result()
	if err != nil {
		logger.Info(err)
	}

	return db, nil
}
