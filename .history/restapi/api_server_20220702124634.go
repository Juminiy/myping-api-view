package restapi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	config "github.com/Juminiy/myping-api-view/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

var (
	errgp errgroup.Group
)

func init() {
	config.AppConfig()
}

func APIEngine() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	apiEngine := gin.New()
	apiEngine.Use(gin.Recovery())
	apiEngine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "success",
				"data":    viper.GetString("web.restapi.version"),
			},
		)
	})

	v1 := apiEngine.Group("/v1")
	{
		v1.
	}
	return apiEngine
}

func APIServer() {
	fmt.Println("😜 API Server is listening port :", viper.GetString("web.api.listen"))
	apiServer := &http.Server{
		Addr:         viper.GetString("web.api.listen"),
		Handler:      APIEngine(),
		ReadTimeout:  time.Second * time.Duration(viper.GetInt("web.api.timeout.read")),
		WriteTimeout: time.Second * time.Duration(viper.GetInt("web.api.timeout.write")),
	}
	errgp.Go(func() error {
		return apiServer.ListenAndServe()
	})

	if err := errgp.Wait(); err != nil {
		log.Fatal(err)
	}
}
