package user_model

import "We-do-secure/database"

func InitMigrate() {
	database.DB.AutoMigrate(
		&User{},&Password_His{},
	)
}
