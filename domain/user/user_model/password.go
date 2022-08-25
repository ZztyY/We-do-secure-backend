package user_model

import (
	"We-do-secure/database"
	"We-do-secure/util"
)

type Password_His struct {
	Pwid    uint                `gorm:"primary_key" json:"pwid"`
	Up_Date util.JSONDetailTime `json:"up_date"`
	New     string              `gorm:"type:varchar(30)" json:"new"`
	Old     string              `gorm:"type:varchar(30)" json:"old"`
	UID     uint                `json:"uid"`
}

func CreatePassword_His(password *Password_His) {
	err := database.DB.Create(password).Error
	if err != nil {
		panic(err)
	}
}

func GetPassword_His(pwid uint) *Password_His {
	var password Password_His
	database.DB.First(&password, pwid)
	if password.Pwid == 0 {
		return nil
	}
	return &password
}

func GetPassword_HisByUid(uId uint) ([]Password_His, int) {
	var list []Password_His
	var count int
	db := database.DB.Model(&User{}).Where("uid = ?", uId)
	db.Count(&count)
	db.Find(&list)
	return list, count
}
