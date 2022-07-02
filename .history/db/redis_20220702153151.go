package db

import (
	"context"

	"github.com/Juminiy/myping-api-view/config"
	"github.com/Juminiy/myping/db"
	"github.com/go-redis/redis/v9"
)

const (
	
)

var (
	RedisClient   redis.Client
	GlobalCentext context.Context
)

func init() {
	config.AppConfig()
	db.RedisConnect()
	RedisClient = *db.RedisClient
	GlobalCentext = db.Ctx
}