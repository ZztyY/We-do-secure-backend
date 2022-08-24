package home_model

import "We-do-secure/database"

func InitMigrate() {
	database.DB.AutoMigrate(
		&Home{},
	)
}
