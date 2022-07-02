package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/Juminiy/myping-api-view/config"
)

func init () {
	config.AppConfig()
}

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
