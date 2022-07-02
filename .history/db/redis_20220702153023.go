package db

import (
	"github.com/Juminiy/myping-api-view/config"
	"github.com/Juminiy/myping/db"
	"github.com/go-redis/redis/v9"
)

var (
	RedisClient redis.Client
	Global
)

func init() {
	config.AppConfig()
	db.RedisConnect()
	RedisClient = *db.RedisClient

}
