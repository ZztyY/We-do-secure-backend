package pol_model

import "We-do-secure/database"

func InitMigrate() {
	database.DB.AutoMigrate(
		&Pol{}, &PolHome{}, &PolVeh{},
	)
}

