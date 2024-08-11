package mysql

import (
	"fmt"
	"github.com/cat9host/gin-air-boilerplate/internal/config"
	"github.com/cat9host/gin-air-boilerplate/internal/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

// DB is single instance of our database connection pool.
var db *gorm.DB
var mutex = &sync.Mutex{}

// Connects to DB
func openConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Error(
			fmt.Sprintf("error has thrown while connecting to database: %v", err),
			"DB",
		)
		time.Sleep(3 * time.Second)
		return openConnection(dsn)
	} else {
		sqlDB, err := db.DB()

		if err != nil {
			log.Error(
				fmt.Sprintf("error has thrown while connecting to database: %v", err),
				"DB",
			)
		}

		sqlDB.SetMaxOpenConns(30)
		sqlDB.SetMaxIdleConns(30)
		sqlDB.SetConnMaxLifetime(10 * time.Second)
		sqlDB.SetConnMaxIdleTime(10 * time.Second)
	}

	return db
}

func connect() *gorm.DB {
	return openConnection(config.MySqlDSN)
}

func PingDB() error {
	if err := GetDBConnection().Raw("SELECT 1").Error; err != nil {
		return err
	}

	return nil
}

// GetDBConnection returns database connection instance
func GetDBConnection() *gorm.DB {
	mutex.Lock()
	if db == nil {
		db = connect()
	}
	mutex.Unlock()

	return db
}
