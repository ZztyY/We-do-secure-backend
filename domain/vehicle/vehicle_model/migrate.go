package vehicle_model

import "We-do-secure/database"

func InitMigrate() {
	database.DB.AutoMigrate(
		&Vehicle{}, &Driver{}, &VehDri{},
	)
}
