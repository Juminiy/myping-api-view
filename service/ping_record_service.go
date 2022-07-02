package service

import (
	"github.com/Juminiy/myping-api-view/db"
)

func PingRecordKV(keys bool) []string {
	resultList := make([]string, 0)
	iter := db.RedisClient.Scan(
		db.GlobalCentext,
		db.REDIS_PING_RECORD_CURSOR,
		db.PING_RECORD_DEFAULT_PREFIX+"*",
		db.REDIS_PING_RECORD_COUNT,
	).Iterator()
	for iter.Next(db.GlobalCentext) {
		if keys {
			resultList = append(resultList, iter.Val())
		} else {
			val := db.RedisClient.Get(db.GlobalCentext, iter.Val()).Val()
			resultList = append(resultList, val)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return resultList
}
