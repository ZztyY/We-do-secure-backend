package vehicle_model

import "We-do-secure/database"

type VehDri struct {
	ID 			uint  `gorm:"primary_key" json:"id"`
	VID			uint  `json:"vid"`
	DID			uint  `json:"did"`
}

func CreateVehDri(vehDri *VehDri) {
	err := database.DB.Create(vehDri).Error
	if err != nil {
		panic(err)
	}
}

func FindVehDriListByVId(vId uint) []VehDri {
	var list []VehDri
	database.DB.Model(&VehDri{}).Where("v_id = ?", vId).Find(&list)
	return list
}
