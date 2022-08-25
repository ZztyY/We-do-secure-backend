package invoice_model

import "We-do-secure/database"

func InitMigrate() {
	database.DB.AutoMigrate(
		&Invoice{}, &Payment{},
	)
}
