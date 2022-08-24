package home_model

import (
	"We-do-secure/database"
	"We-do-secure/util"
)

type Home struct {
	HID     uint                `gorm:"primary_key" json:"hid"`
	PURDate util.JSONDetailTime `json:"pur_date"`
	PURVal  float64             `gorm:"type:decimal(10,2)" json:"pur_val"`
	HArea   float64             `gorm:"type:decimal(6,2)" json:"harea"`
	HType   string              `gorm:"type:varchar(1)" json:"htype"`
	HAfn    int                 `gorm:"type:tinyint(4)" json:"hafn"`
	HHss    int                 `gorm:"type:tinyint(4)" json:"hhss"`
	Hsp     string              `gorm:"type:varchar(1)" json:"hsp"`
	Hbm     int                 `gorm:"type:tinyint(4)" json:"hbm"`
	UID     uint				`json:"uid"`
}

func CreateHome(home *Home) {
	err := database.DB.Create(home).Error
	if err != nil {
		panic(err)
	}
}

func UpdateHome(home *Home) {
	err := database.DB.Model(&Home{}).Save(home).Error
	if err != nil {
		panic(err)
	}
}

func GetHome(id uint) *Home {
	var home Home
	database.DB.First(&home, id)
	if home.HID == 0 {
		return nil
	}
	return &home
}

func FindHome(filter map[string]interface{}) *Home {
	var home Home
	database.DB.Where(filter).Last(&home)
	if home.HID == 0 {
		return nil
	}
	return &home
}

func FindHomeList(offset int, limit int) ([]Home, int) {
	var count int
	var list []Home
	db := database.DB.Model(&Home{})
	db.Offset(offset).Limit(limit).Find(&list)
	db.Count(&count)
	return list, count
}

func FindUserHomeList(offset int, limit int, uId uint) ([]Home, int) {
	var list []Home
	var count int
	db := database.DB.Model(&Home{}).Where("uid = ?", uId)
	db.Count(&count)
	db.Offset(offset).Limit(limit).Find(&list)
	return list, count
}
