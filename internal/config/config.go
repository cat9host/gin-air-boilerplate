package config

import (
	"fmt"
	"github.com/cat9host/gin-air-boilerplate/internal/log"
	"os"
)

var AppPort string
var PromPort string
var HCPort string
var AppSecretKey string
var MySqlDSN string

func cfgPanic() {
	log.FatalError("Necessary environment variables not fount", "CONFIG")
}

func logMissedVariable(property string) {
	log.Warn(fmt.Sprintf("%s not presented in environment", property), "CONFIG")
}

func Configure() {
	var isSetAppPort, isSetPromPort, isSetHCPort bool

	// some of the variables we can cover with default
	AppPort, isSetAppPort = os.LookupEnv("PORT")
	if !isSetAppPort {
		AppPort = "25100"
	}
	PromPort, isSetPromPort = os.LookupEnv("METRICS_PORT")
	if !isSetPromPort {
		PromPort = "10001"
	}
	HCPort, isSetHCPort = os.LookupEnv("HC_PORT")
	if !isSetHCPort {
		HCPort = "10002"
	}

	MySqlDSN = os.Getenv("MYSQL_DSN")

	AppSecretKey = os.Getenv("SECRET_KEY")

	if MySqlDSN == "" {
		logMissedVariable("MYSQL_DSN")
	}

	if AppSecretKey == "" {
		logMissedVariable("SECRET_KEY")
		cfgPanic()
	}
}
