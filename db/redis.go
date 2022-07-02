package db

import (
	"github.com/Juminiy/myping-api-view/config"
	"github.com/Juminiy/myping/db"
)

func init() {
	config.AppConfig()
	db.RedisConnect()
}
