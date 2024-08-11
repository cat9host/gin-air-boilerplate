package hc

import (
	"github.com/cat9host/gin-air-boilerplate/internal/db/mysql"
	"github.com/cat9host/gin-air-boilerplate/internal/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

func hcResponse(sql bool) *interfaces.HCResponse {
	MySql := interfaces.ResultOk
	if !sql {
		MySql = interfaces.ResultErr
	}
	return &interfaces.HCResponse{
		Status: sql,
		MySql:  MySql,
	}
}

func checkSql() error {
	mysql.GetDBConnection()

	return mysql.PingDB()
}

func HealthCheckHandle(c *gin.Context) {
	dbError := checkSql()

	c.JSON(
		http.StatusOK,
		hcResponse(
			dbError == nil,
		),
	)
}
