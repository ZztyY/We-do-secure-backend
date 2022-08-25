package invoice

import (
	"We-do-secure/domain/cus/cus_model"
	"We-do-secure/domain/invoice/invoice_model"
	"We-do-secure/domain/invoice/invoice_service"
	"We-do-secure/domain/pol/pol_model"
	"We-do-secure/domain/user/user_model"
	"We-do-secure/interfaces/errorcode"
	"We-do-secure/interfaces/response"
	"We-do-secure/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PayedInvoice struct {
	*invoice_model.Invoice
	AmountLeft 					float64  `json:"amount_left"`
	Status  					string   `json:"status"`
}

func UserInvoiceList(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	uId := c.DefaultQuery("uid", "")

	if uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}
	cus := cus_model.GetCusByUid(util.StrToUInt(uId))

	var list []pol_model.Pol
	if cus == nil {
		response.SendSuccess(c, map[string]interface{}{"count": 0, "list": list})
	} else {
		list, _ = pol_model.FindPolList(0, 100, cus.CID)
		var pIdList []uint
		for _, v := range list {
			pIdList = append(pIdList, v.PID)
		}

		invoiceList, count := invoice_model.FindInvoiceListByPIdList(util.StrToInt(offset), util.StrToInt(limit), pIdList)
		response.SendSuccess(c, map[string]interface{}{"count": count, "list": invoiceList})
	}
}

func GetInvoice(c *gin.Context) {
	iId := c.DefaultQuery("iid", "")
	if iId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	if invoice_model.GetInvoice(util.StrToUInt(iId)) == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invoice doesn't exist", nil)
		return
	}

	invoice := invoice_model.GetInvoice(util.StrToUInt(iId))



	payedInvoice := PayedInvoice{}
	payedInvoice.Invoice = invoice
	payedInvoice.AmountLeft = invoice_service.DueAmountLeft(invoice.IID)
	if invoice.IDueDate.Before(time.Now()) {
		payedInvoice.Status = "P"
	} else {
		payedInvoice.Status = "C"
	}
	response.SendSuccess(c, payedInvoice)
}

func UpdatePaymentMethod(c *gin.Context) {
	uId := c.PostForm("uid")
	pMethod := c.PostForm("p_method")
	pAccountNum := c.PostForm("p_account_num")

	if uId == "" || pMethod == "" || pAccountNum == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	user := user_model.GetUser(util.StrToUInt(uId))

	if user == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invalid user", nil)
		return
	}
	if pMethod != "Credit" && pMethod != "PayPal" && pMethod != "Debit" && pMethod != "Check" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invalid method", nil)
		return
	}

	user.PMethod = pMethod
	user.PAccountNum = pAccountNum
	user_model.UpdateUser(user)
	response.SendSuccess(c, "success")
}

func MakePayment(c *gin.Context) {
	uId := c.PostForm("uid")
	iId := c.PostForm("iid")
	pAmount := c.PostForm("pAmount")
	if uId == "" || iId == "" || pAmount == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	user := user_model.GetUser(util.StrToUInt(uId))
	invoice := invoice_model.GetInvoice(util.StrToUInt(iId))

	if user == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invalid user", nil)
		return
	}
	if user.PMethod != "Credit" && user.PMethod != "Debit" && user.PMethod != "PayPal" && user.PMethod != "Check" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invalid pay method", nil)
		return
	}

	if invoice == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invalid invoice", nil)
		return
	}
	if invoice.IDueDate.Before(time.Now()) {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invoice has due", nil)
		return
	}
	p, _ := strconv.ParseFloat(pAmount, 64)
	if p > invoice_service.DueAmountLeft(invoice.IID) {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "pay amount too large", nil)
		return
	}


	payment := invoice_model.Payment{}
	payment.IID = invoice.IID
	payment.PMethod = user.PMethod
	payment.PAmount = p
	payment.PDate = util.JSONDetailTime{Time: time.Now()}
	invoice_model.CreatePayment(&payment)
	response.SendSuccess(c, payment)
}

func PaymentList(c *gin.Context) {
	iId := c.DefaultQuery("iid", "")
	if iId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}
	invoice := invoice_model.GetInvoice(util.StrToUInt(iId))
	if invoice == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "invalid invoice", nil)
		return
	}
	list, count := invoice_model.GetPaymentByIId(invoice.IID)
	response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
}
