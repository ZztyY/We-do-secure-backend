package invoice_service

import (
	"We-do-secure/domain/invoice/invoice_model"
)

func DueAmountLeft(iId uint) float64 {
	invoice := invoice_model.GetInvoice(iId)
	list, count := invoice_model.GetPaymentByIId(iId)
	if count == 0 {
		return invoice.IAmount
	} else {
		var payedAmount float64
		for _, v := range list {
			payedAmount += v.PAmount
		}
		return invoice.IAmount - payedAmount
	}
}
