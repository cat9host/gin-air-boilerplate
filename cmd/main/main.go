package main

import (
	"fmt"
	"github.com/cat9host/gin-air-boilerplate/internal/app"
	"github.com/cat9host/gin-air-boilerplate/internal/config"
	"github.com/cat9host/gin-air-boilerplate/internal/db/mysql"
	"github.com/cat9host/gin-air-boilerplate/internal/log"
)

func main() {
	//configure
	config.Configure()
	//init logger
	log.InitializeLogger()
	routerMain, routerProm, routerHC := app.SetupRouter(true)

	mysql.GetDBConnection()

	go func() {
		log.Info(fmt.Sprintf("Metrics started at port %s", config.PromPort), "APP")
		err := routerProm.Run(fmt.Sprintf(":%s", config.PromPort))
		if err != nil {
			return
		}
	}()
	go func() {
		log.Info(fmt.Sprintf("HealthCheck started at port %s", config.HCPort), "APP")
		err := routerHC.Run(fmt.Sprintf(":%s", config.HCPort))
		if err != nil {
			return
		}
	}()

	log.Info(fmt.Sprintf("Application started at port %s", config.AppPort), "APP")
	err := routerMain.Run(fmt.Sprintf(":%s", config.AppPort))

	if err != nil {
		panic(err)
	}
}
