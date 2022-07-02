package handler

import (
	"net/http"
	"github.com/Ju"
	"github.com/gin-gonic/gin"
)

func HistoryCMDHandler(ctx *gin.Context) {
	RedisClient.
	ctx.HTML(
		http.StatusOK, 
		"history_cmd.tmpl",
		gin.H{
			"" : "",
		},
	)
}