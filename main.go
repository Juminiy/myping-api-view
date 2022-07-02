package main

import (
	"github.com/Juminiy/myping-api-view/config"
	"github.com/Juminiy/myping-api-view/restapi"
)

func init() {
	config.AppConfig()
}

func main() {

	// Web API
	restapi.APIServer()

}
