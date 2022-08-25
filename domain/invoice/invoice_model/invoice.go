package invoice_model

import (
	"We-do-secure/database"
	"We-do-secure/util"
)

type Invoice struct {
	IID      uint                `gorm:"primary_key" json:"iid"`
	IDueDate util.JSONDetailTime `json:"iduedate"`
	IAmount  float64             `gorm:"type:decimal(7,2)" json:"iamount"`
	PID      uint                `json:"pid"`
}

func CreateInvoice(invoice *Invoice) {
	err := database.DB.Create(invoice).Error
	if err != nil {
		panic(err)
	}
}

func UpdateInvoice(invoice *Invoice) {
	err := database.DB.Model(&Invoice{}).Save(invoice).Error
	if err != nil {
		panic(err)
	}
}

func GetInvoice(iid uint) *Invoice {
	var invoice Invoice
	database.DB.First(&invoice, iid)
	if invoice.IID == 0 {
		return nil
	}
	return &invoice
}

func GetInvoiceByPId(pId uint) *Invoice {
	var invoice Invoice
	database.DB.Model(&Invoice{}).Where("p_id = ?", pId).First(&invoice)
	if invoice.IID == 0 {
		return nil
	}
	return &invoice
}

func FindInvoiceListByPIdList(offset int, limit int, pIdList []uint) ([]Invoice, int) {
	var list []Invoice
	var count int
	db := database.DB.Model(&Invoice{}).Where("p_id in (?)", pIdList)
	db.Count(&count)
	db.Offset(offset).Limit(limit).Find(&list)
	return list, count
}
