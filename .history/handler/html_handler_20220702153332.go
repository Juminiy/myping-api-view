package handler

import (
	"net/http"
	db "github.com/Juminiy/myping-api-view/db"
	"github.com/gin-gonic/gin"
)

func HistoryCMDHandler(ctx *gin.Context) {
	db.RedisClient.Get(db.GlobalCentext,PING_RECORD_DEFAULT_PREFIX)
	ctx.HTML(
		http.StatusOK, 
		"history_cmd.tmpl",
		gin.H{
			"" : "",
		},
	)
}