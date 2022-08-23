package user_model

import (
	"We-do-secure/database"
	"We-do-secure/util"
)

type User struct {
	ID             uint                `gorm:"primary_key" json:"id"`
	UserName       string              `gorm:"type:varchar(30);unique" json:"user_name"`
	Password       string              `gorm:"type:varchar(30)" json:"-"`
	Token          string              `gorm:"type:varchar(30)" json:"token"`
	TokenExpiredAt util.JSONDetailTime `json:"token_expired_at"`
	CreatedAt      util.JSONTime       `json:"created_at"`
	UpdatedAt      util.JSONTime       `json:"updated_at"`
	DeletedAt      *util.JSONTime      `json:"deleted_at"`
}

func CreateUser(user *User) {
	err := database.DB.Create(user).Error
	if err != nil {
		panic(err)
	}
}

func UpdateUser(user *User) {
	err := database.DB.Model(&User{}).Save(user).Error
	if err != nil {
		panic(err)
	}
}

func GetUser(id uint) *User {
	var user User
	database.DB.First(&user, id)
	return &user
}

func FindUser(filter map[string]interface{}) *User {
	var user User
	database.DB.Where(filter).Last(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func FindUserList(offset int, limit int) ([]User, int) {
	var count int
	var list []User
	db := database.DB.Model(&User{})
	db.Offset(offset).Limit(limit).Find(&list)
	db.Count(&count)
	return list, count
}
