package db

import (
	"github.com/Juminiy/myping/cofig"
	"github.com/Juminiy/myping/db"
)

func init() {
	config.AppConfig()
	db.RedisConnect()
}