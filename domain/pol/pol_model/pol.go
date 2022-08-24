package pol_model

import (
	"We-do-secure/database"
	"We-do-secure/util"
)

type Pol struct {
	PID     uint                `gorm:"primary_key" json:"pid"`
	SDate   util.JSONDetailTime `json:"start_date"`
	EDate   util.JSONDetailTime `json:"end_date"`
	PAmount float64             `gorm:"type:decimal(7,2)" json:"pamount"`
	PStatus string              `gorm:"type:varchar(1)" json:"pstatus"`
	PType   string              `gorm:"type:varchar(1)" json:"ptype"`
	CID     uint                `json:"cid"`
}

func CreatePol(pol *Pol) {
	err := database.DB.Create(pol).Error
	if err != nil {
		panic(err)
	}
}

func UpdatePol(pol *Pol) {
	err := database.DB.Model(&Pol{}).Save(pol).Error
	if err != nil {
		panic(err)
	}
}

func GetPol(pid uint) *Pol {
	var pol Pol
	database.DB.First(&pol, pid)
	if pol.PID == 0 {
		return nil
	}
	return &pol
}

func FindPolList(offset int, limit int, cId uint) ([]Pol, int) {
	var list []Pol
	var count int
	db := database.DB.Model(&Pol{}).Where("c_id = ?", cId)
	db.Count(&count)
	db.Offset(offset).Limit(limit).Find(&list)
	return list, count
}

func FindPolListByFilter(filter map[string]interface{}, offset int, limit int) ([]Pol, int) {
	var list []Pol
	var count int

	db := database.DB.Model(&Pol{}).Where(filter)
	db.Count(&count)
	db.Offset(offset).Limit(limit).Find(&list)
	return list, count
}
