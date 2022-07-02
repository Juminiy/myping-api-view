package handler

import (
	"net/http"
	db "github.com/Juminiy/myping/db"
	"github.com/gin-gonic/gin"
)

func HistoryCMDHandler(ctx *gin.Context) {
	
	ctx.HTML(
		http.StatusOK, 
		"history_cmd.tmpl",
		gin.H{
			"" : "",
		},
	)
}