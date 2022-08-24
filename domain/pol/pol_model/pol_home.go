package pol_model

import "We-do-secure/database"

type PolHome struct {
	ID  uint `gorm:"primary_key" json:"id"`
	PID uint `json:"pid"`
	HID uint `json:"hid"`
}

func CreatePolHome(polHome *PolHome) {
	err := database.DB.Create(polHome).Error
	if err != nil {
		panic(err)
	}
}

func FindPolHome(filter map[string]interface{}) *PolHome {
	var polHome PolHome
	database.DB.Where(filter).Last(&polHome)
	if polHome.ID == 0 {
		return nil
	}
	return &polHome
}
