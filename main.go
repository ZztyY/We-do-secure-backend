package main

import (
	"We-do-secure/domain/cus/cus_model"
	"We-do-secure/domain/home/home_model"
	"We-do-secure/domain/pol/pol_model"
	"We-do-secure/domain/user/user_model"
	"We-do-secure/domain/vehicle/vehicle_model"
	"We-do-secure/env"
	"We-do-secure/interfaces"
	"We-do-secure/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {

	// 日志信息
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	if env.APP.Env != "prod" {
		logrus.SetOutput(os.Stdout)
	} else {
		f, _ := os.Create(os.TempDir() + "/gin." + util.GetTodayDate() + ".log")
		logrus.SetOutput(f)
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
	r := gin.Default()

	// set app mode
	if env.APP.AppMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// init mysql
	InitMigrate()

	// init api router
	interfaces.InitRouter(r)

	//start service
	err := r.Run(":8088")
	if err != nil {
		logrus.Fatal("start service error:", err)
	}
}

func InitMigrate() {
	// 基础服务
	user_model.InitMigrate()
	home_model.InitMigrate()
	cus_model.InitMigrate()
	vehicle_model.InitMigrate()
	pol_model.InitMigrate()
}
