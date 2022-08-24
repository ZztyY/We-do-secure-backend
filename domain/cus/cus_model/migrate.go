package cus_model

import "We-do-secure/database"

func InitMigrate() {
	database.DB.AutoMigrate(
		&Cus{},
	)
}
