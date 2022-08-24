package cus_model

import (
	"We-do-secure/database"
)

type Cus struct {
	CID             uint                `gorm:"primary_key" json:"cid"`
	FName      		string              `gorm:"type:varchar(30)" json:"fName"`
	LName       	string              `gorm:"type:varchar(30)" json:"lName"`
	Street         	string              `gorm:"type:varchar(30)" json:"street"`
	City         	string              `gorm:"type:varchar(30)" json:"city"`
	State         	string              `gorm:"type:varchar(30)" json:"state"`
	Zipcode         int              	`gorm:"type:int(5)" json:"zipcode"`
	Gender          string              `gorm:"type:varchar(1)" json:"gender"`
	MarStatus       string              `gorm:"type:varchar(1)" json:"mar_status"`
	UID             uint                `json:"uid"`
}

func CreateCus(cus *Cus) {
	err := database.DB.Create(cus).Error
	if err != nil {
		panic(err)
	}
}

func UpdateCus(cus *Cus) {
	err := database.DB.Model(&Cus{}).Save(cus).Error
	if err != nil {
		panic(err)
	}
}

func GetCus(cid uint) *Cus {
	var cus Cus
	database.DB.First(&cus, cid)
	return &cus
}

func GetCusByUid(uid uint) *Cus {
	var cus Cus
	database.DB.Where("uid = ?", uid).First(&cus)
	return &cus
}
