package handler

import (
	"net/http"
	"strings"

	db "github.com/Juminiy/myping-api-view/db"
	myping_service "github.com/Juminiy/myping-api-view/service"
	"github.com/gin-gonic/gin"
)

// time.Date(2022, time.July, 2, 11, 15, 3, 44557661, time.Local)
func PingSearchHandler(ctx *gin.Context) {
	year := ctx.Query("year")
	mon := ctx.Query("mon")
	day := ctx.Query("day")
	hour := ctx.Query("hour")
	min := ctx.Query("min")
	sec := ctx.Query("sec")
	stamp := ctx.Query("stamp")

	noneVersionKey := strings.Join([]string{"time.Date(" + year, "time." + mon, day, hour, min, sec, stamp, "time.Local)"}, ", ")
	//fmt.Println(noneVersionKey)
	val := db.RedisClient.Get(db.GlobalCentext, db.PING_RECORD_DEFAULT_PREFIX+noneVersionKey).Val()

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    val,
		},
	)
}

func PingHistoryRecordKey(ctx *gin.Context) {
	recordKeys := myping_service.PingRecordKV(true)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    recordKeys,
		},
	)
}
func PingHistoryRecordVal(ctx *gin.Context) {
	recordVals := myping_service.PingRecordKV(false)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    recordVals,
		},
	)
}

func PingHistoryArrTimestamps(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H {
			"code":    http.StatusOK,
			"message": "success",
			"data":    myping_service.PingTimestampArr(),
		},
	)
}

func PingHistoryArrRecords(ctx *gin.Context) {
	ky := ctx.Query("propertyKey")
	ctx.JSON(
		http.StatusOK,
		gin.H {
			"code":    http.StatusOK,
			"message": "success",
			"data":    myping_service.PingRecordArr(ky),
		},
	)
}