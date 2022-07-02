package db

import (
	"github.com/Juminiy/myping/co"
	"github.com/Juminiy/myping/db"
)

func init() {
	config.AppConfig()
	db.RedisConnect()
}
