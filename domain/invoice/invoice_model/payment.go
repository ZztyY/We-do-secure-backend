package invoice_model

import (
	"We-do-secure/database"
	"We-do-secure/util"
)

type Payment struct {
	PNo     uint                `gorm:"primary_key" json:"pno"`
	PDate   util.JSONDetailTime `json:"pdate"`
	PMethod string              `gorm:"type:varchar(6)" json:"pmethod"`
	PAmount float64				`gotm:"type:decimal(7,2)" json:"pamount"`
	IID     uint                `json:"iid"`
}

func CreatePayment(payment *Payment) {
	err := database.DB.Create(payment).Error
	if err != nil {
		panic(err)
	}
}

func UpdatePayment(payment *Payment) {
	err := database.DB.Model(&Payment{}).Save(payment).Error
	if err != nil {
		panic(err)
	}
}

func GetPayment(pno uint) *Payment {
	var payment Payment
	database.DB.First(&payment, pno)
	if payment.PNo == 0 {
		return nil
	}
	return &payment
}

func GetPaymentByIId(iId uint) ([]Payment, int) {
	var list []Payment
	var count int
	db := database.DB.Model(&Invoice{}).Where("i_id = ?", iId)
	db.Count(&count)
	db.Find(&list)
	return list, count
}
