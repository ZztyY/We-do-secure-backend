package pol

import (
	"We-do-secure/domain/cus/cus_model"
	"We-do-secure/domain/invoice/invoice_model"
	"We-do-secure/domain/pol/pol_model"
	"We-do-secure/domain/pol/pol_service"
	"We-do-secure/interfaces/errorcode"
	"We-do-secure/interfaces/response"
	"We-do-secure/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func UserPolList(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	uId := c.DefaultQuery("uid", "")
	pType := c.DefaultQuery("ptype", "")
	if uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}
	cus := cus_model.GetCusByUid(util.StrToUInt(uId))
	var list []pol_model.Pol
	var count int
	if cus == nil {
		response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
	} else {
		if pType == "" {
			list, count = pol_model.FindPolList(util.StrToInt(offset), util.StrToInt(limit), cus.CID)
			response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
		} else {
			filter := make(map[string]interface{})
			filter["c_id"] = cus.CID
			filter["p_type"] = pType
			list, count = pol_model.FindPolListByFilter(filter, util.StrToInt(offset), util.StrToInt(limit))
			response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
		}
	}
}

func EditPol(c *gin.Context) {
	pId := c.PostForm("pid")
	sDate := c.PostForm("start_date")
	months := c.PostForm("months")
	pAmount := c.PostForm("pamount")
	pType := c.PostForm("ptype")
	uId := c.PostForm("uid")

	if sDate == "" || months == "" || pAmount == "" || pType == "" || uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	cus := cus_model.GetCusByUid(util.StrToUInt(uId))
	if cus == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "please fill profile", nil)
		return
	}

	if pId == "" {
		var vId string
		var hId string
		if pType == "A" {
			vId = c.PostForm("vid")
			if vId == "" {
				response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
				return
			}
			if pol_service.GetPolVehByVId(util.StrToUInt(vId)) != nil {
				response.SendError(c, errorcode.CODE_PARAMS_INVALID, "policy exists", nil)
				return
			}
		} else {
			hId = c.PostForm("hid")
			if hId == "" {
				response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
				return
			}
			if pol_service.GetPolHomeByHId(util.StrToUInt(hId)) != nil {
				response.SendError(c, errorcode.CODE_PARAMS_INVALID, "policy exists", nil)
				return
			}
		}

		pol := pol_model.Pol{}

		pol.SDate = util.JSONDetailTime{Time: util.StrToTimeTime(sDate)}
		pol.EDate = util.JSONDetailTime{Time: util.StrToTimeTime(sDate).AddDate(0, util.StrToInt(months), 0)}
		pol.PAmount, _ = strconv.ParseFloat(pAmount, 64)
		if pol.EDate.Before(time.Now()) {
			pol.PStatus = "P"
		} else {
			pol.PStatus = "C"
		}
		pol.PType = pType
		pol.CID = cus.CID

		pol_model.CreatePol(&pol)

		invoice := invoice_model.Invoice{}
		invoice.PID = pol.PID
		invoice.IDueDate = util.JSONDetailTime{Time: util.StrToTimeTime(sDate).AddDate(0, 1, 0)}
		invoice.IAmount = pol.PAmount

		invoice_model.CreateInvoice(&invoice)

		if pType == "A" {
			polVeh := pol_model.PolVeh{}
			polVeh.VID = util.StrToUInt(vId)
			polVeh.PID = pol.PID

			pol_model.CreatePolVeh(&polVeh)
		} else {
			polHome := pol_model.PolHome{}
			polHome.HID = util.StrToUInt(hId)
			polHome.PID = pol.PID

			pol_model.CreatePolHome(&polHome)
		}
		response.SendSuccess(c, pol)
	} else if pol_model.GetPol(util.StrToUInt(pId)) == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "Invalid PID", nil)
		return
	} else {
		pol := pol_model.GetPol(util.StrToUInt(pId))

		pol.SDate = util.JSONDetailTime{Time: util.StrToTimeTime(sDate)}
		pol.EDate = util.JSONDetailTime{Time: util.StrToTimeTime(sDate).AddDate(0, util.StrToInt(months), 0)}
		pol.PAmount, _ = strconv.ParseFloat(pAmount, 64)
		if pol.EDate.Before(time.Now()) {
			pol.PStatus = "P"
		} else {
			pol.PStatus = "C"
		}
		pol.CID = cus.CID

		pol_model.UpdatePol(pol)

		invoice := invoice_model.GetInvoiceByPId(pol.PID)
		if invoice == nil {
			invoice := invoice_model.Invoice{}
			invoice.IAmount = pol.PAmount
			invoice.IDueDate = util.JSONDetailTime{Time: util.StrToTimeTime(sDate).AddDate(0, 1, 0)}
			invoice.PID = pol.PID

			invoice_model.CreateInvoice(&invoice)
		} else {
			invoice.IAmount = pol.PAmount
			invoice.IDueDate = util.JSONDetailTime{Time: util.StrToTimeTime(sDate).AddDate(0, 1, 0)}
			invoice_model.UpdateInvoice(invoice)
		}
		response.SendSuccess(c, pol)
	}
}

func GetPol(c *gin.Context) {
	pId := c.DefaultQuery("pid", "")
	if pId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "please input policy id", nil)
		return
	}
	pol := pol_model.GetPol(util.StrToUInt(pId))
	if pol == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "policy doesn't exist", nil)
		return
	}
	response.SendSuccess(c, pol)
}
