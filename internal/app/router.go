package app

import (
	"github.com/cat9host/gin-air-boilerplate/internal/hc"
	"github.com/cat9host/gin-air-boilerplate/internal/interfaces"
	"github.com/cat9host/gin-air-boilerplate/internal/log"
	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"net/http"
	"strings"
	"time"
)

func defineStatic(routerMain *gin.Engine, routerProm *gin.Engine) {
	// for prevent annoying error in logs on any request
	routerMain.StaticFile("/favicon.ico", "./static/favicon.ico")
	routerProm.StaticFile("/favicon.ico", "./static/favicon.ico")
}

func defineRoutes(routerMain *gin.Engine, routerProm *gin.Engine) {
	defineStatic(routerMain, routerProm)

	routerMain.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// debug group
	debugGroup := routerMain.Group("/debug")
	debugGroup.Use(AuthMiddleware)

	pprof.RouteRegister(debugGroup, "pprof")

	if gin.Mode() == gin.DebugMode {
		debugGroup.GET("/routes", func(c *gin.Context) {
			var routes []interfaces.RouteInterface
			for _, item := range routerMain.Routes() {
				if strings.Contains(item.Path, "api") || strings.Contains(item.Path, "debug") {
					routes = append(routes, interfaces.RouteInterface{
						Method: item.Method,
						Path:   item.Path,
					})
				}
			}
			c.JSON(http.StatusOK, routes)
		})
	}

	//api routes
	//apiGroup := routerMain.Group("api")
	//apiGroup.Use(AuthMiddleware)
	//v1
	//v1 := apiGroup.Group("v1")
}

func defineHealthCheck(router *gin.Engine) *gin.Engine {
	router.GET("/hc", hc.HealthCheckHandle)

	return router
}

func SetupRouter(withLogger bool) (*gin.Engine, *gin.Engine, *gin.Engine) {
	routerMain := gin.New()
	routerProm := gin.New()
	routerHC := gin.New()

	if withLogger {
		// default logger
		routerMain.Use(ginzap.Ginzap(log.Logger, time.RFC3339, true))

		// Logs all panic to error log
		routerMain.Use(ginzap.RecoveryWithZap(log.Logger, true))
	}

	p := ginprometheus.NewPrometheus("gin-air-boilerplate", nil)
	routerMain.Use(p.HandlerFunc())
	p.SetMetricsPath(routerProm)

	defineRoutes(routerMain, routerProm)
	defineHealthCheck(routerHC)

	return routerMain, routerProm, routerHC
}
