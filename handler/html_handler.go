package handler

import (
	"net/http"

	db "github.com/Juminiy/myping-api-view/db"
	"github.com/gin-gonic/gin"
)

const (
	CASUAL_KEY = "time.Date(2022, time.July, 2, 11, 15, 3, 44557661, time.Local)"
)

func HistoryCMDHandler(ctx *gin.Context) {
	val := db.RedisClient.Get(db.GlobalCentext, db.PING_RECORD_DEFAULT_PREFIX+CASUAL_KEY).Val()
	ctx.HTML(
		http.StatusOK,
		"history_cmd.tmpl",
		gin.H{
			"Code":   http.StatusOK,
			"Record": val,
		},
	)
}

func DashboardHandler(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"dashboard.html",
		gin.H{},
	)
}
