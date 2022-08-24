package vehicle_model

import "We-do-secure/database"

type Vehicle struct {
	VID             uint                `gorm:"primary_key" json:"vid"`
	VIN      		string              `gorm:"type:varchar(17);unique" json:"vin"`
	VMMYear         int                 `json:"vmmyear"`
	VStatus         string         		`gorm:"type:varchar(1)" json:"vstatus"`
	UID             uint                `json:"uid"`
}

func CreateVehicle(vehicle *Vehicle) {
	err := database.DB.Create(vehicle).Error
	if err != nil {
		panic(err)
	}
}

func UpdateVehicle(vehicle *Vehicle) {
	err := database.DB.Model(&Vehicle{}).Save(vehicle).Error
	if err != nil {
		panic(err)
	}
}

func GetVehicle(vid uint) *Vehicle {
	var vehicle Vehicle
	database.DB.First(&vehicle, vid)
	if vehicle.VID == 0 {
		return nil
	}
	return &vehicle
}

func GetVehicleByUid(uid uint) *Vehicle {
	var vehicle Vehicle
	database.DB.Where("uid = ?", uid).First(&vehicle)
	if vehicle.VID == 0 {
		return nil
	}
	return &vehicle
}

func FindUserVehicleList(offset int, limit int, uId uint) ([]Vehicle, int) {
	var list []Vehicle
	var count int
	db := database.DB.Model(&Vehicle{}).Where("uid = ?", uId)
	db.Count(&count)
	db.Offset(offset).Limit(limit).Find(&list)
	return list, count
}
