package handler

import (
	"github.com/gin-gonic/gin"
	"http"
)
func VersionHandler(ctx * gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    viper.GetString("web.api.version"),
		},
	)
}