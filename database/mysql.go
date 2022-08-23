package database

import (
	"We-do-secure/env"
	"We-do-secure/logger"
	"sync"


	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func getDbInstance() {
	cfg := env.APP
	dbConfig := cfg.MysqlUser + ":" + cfg.MysqlPassword + "@(" + cfg.MysqlUrl + ")/" + cfg.MysqlDatabase + "?charset=utf8mb4&parseTime=true&loc=Local"
	connect, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		logger.LogError("error_mysql_connect_fail", err.Error())
	}
	DB = connect
}

func init() {
	var once sync.Once
	once.Do(getDbInstance)
}
