package vehicle_model

import (
	"We-do-secure/database"
	"We-do-secure/util"
)

type Driver struct {
	DID           	uint  				`gorm:"primary_key" json:"did"`
	FName           string  			`gorm:"type:varchar(30)" json:"fName"`
	LName           string  			`gorm:"type:varchar(30)" json:"lName"`
	LNum			string				`gorm:"type:varchar(15)" json:"lNum"`
	Birth           util.JSONDetailTime `json:"birth"`
}

func CreateDriver(driver *Driver) {
	err := database.DB.Create(driver).Error
	if err != nil {
		panic(err)
	}
}

func UpdateDriver(driver *Driver) {
	err := database.DB.Model(&Driver{}).Save(driver).Error
	if err != nil {
		panic(err)
	}
}

func GetDriver(did uint) *Driver {
	var driver Driver
	database.DB.First(&driver, did)
	if driver.DID == 0 {
		return nil
	}
	return &driver
}

func FindDriverListByDIdList(dIdList []uint) []Driver {
	var list []Driver
	database.DB.Model(&Driver{}).Where("d_id in (?)", dIdList).Find(&list)
	return list
}
