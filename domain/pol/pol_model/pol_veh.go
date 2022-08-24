package pol_model

import "We-do-secure/database"

type PolVeh struct {
	ID  uint `gorm:"primary_key" json:"id"`
	PID uint `json:"pid"`
	VID uint `json:"vid"`
}

func CreatePolVeh(polVeh *PolVeh) {
	err := database.DB.Create(polVeh).Error
	if err != nil {
		panic(err)
	}
}

func FindPolVeh(filter map[string]interface{}) *PolVeh {
	var polVeh PolVeh
	database.DB.Where(filter).Last(&polVeh)
	if polVeh.ID == 0 {
		return nil
	}
	return &polVeh
}
