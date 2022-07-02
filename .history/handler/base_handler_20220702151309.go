package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func VersionHandler(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    viper.GetString("web.api.version"),
		},
	)
}
