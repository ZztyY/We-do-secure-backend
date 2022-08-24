package home

import (
	"We-do-secure/domain/home/home_model"
	"We-do-secure/domain/home/home_service"
	"We-do-secure/interfaces/errorcode"
	"We-do-secure/interfaces/response"
	"We-do-secure/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddHome(c *gin.Context) {
	id := c.PostForm("hid")
	purDate := c.PostForm("pur_date")
	purVal := c.PostForm("pur_val")
	hArea := c.PostForm("harea")
	hType := c.PostForm("htype")
	hAfn := c.PostForm("hafn")
	hHss := c.PostForm("hhss")
	Hsp := c.PostForm("hsp")
	Hbm := c.PostForm("hbm")
	uId := c.PostForm("uid")

	if purDate == "" || purVal == "" || hArea == "" || hType == "" || hAfn == "" || hHss == "" || Hbm == "" || uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	if id == "" {
		home := home_model.Home{}
		home.PURDate = util.JSONDetailTime{Time: util.StrToTimeTime(purDate)}
		home.PURVal, _ = strconv.ParseFloat(purVal, 64)
		home.HArea, _ = strconv.ParseFloat(hArea, 64)
		home.HType = hType
		home.HAfn = util.StrToInt(hAfn)
		home.HHss = util.StrToInt(hHss)
		home.Hsp = Hsp
		home.Hbm = util.StrToInt(Hbm)
		home.UID = util.StrToUInt(uId)
		home_model.CreateHome(&home)
		response.SendSuccess(c, home)
	} else if home_model.GetHome(util.StrToUInt(id)) == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "Invalid ID", nil)
		return
	} else {
		home := home_model.GetHome(util.StrToUInt(id))

		home.PURDate = util.JSONDetailTime{Time: util.StrToTimeTime(purDate)}
		home.PURVal, _ = strconv.ParseFloat(purVal, 64)
		home.HArea, _ = strconv.ParseFloat(hArea, 64)
		home.HType = hType
		home.HAfn = util.StrToInt(hAfn)
		home.HHss = util.StrToInt(hHss)
		home.Hsp = Hsp
		home.Hbm = util.StrToInt(Hbm)
		home.UID = util.StrToUInt(uId)

		home_model.UpdateHome(home)
		response.SendSuccess(c, home)
	}
}

func HomeList(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	list, count := home_service.FindHomeList(util.StrToInt(offset), util.StrToInt(limit))
	response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
}

func GetHome(c *gin.Context) {
	id := c.DefaultQuery("hid", "")
	if id == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "please input home id", nil)
		return
	}
	home := home_model.GetHome(util.StrToUInt(id))
	if home == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "home doesn't exist", nil)
		return
	}
	response.SendSuccess(c, home)
}

func UserHomeList(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	uId := c.DefaultQuery("uid", "")
	if uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}
	list, count := home_model.FindUserHomeList(util.StrToInt(offset), util.StrToInt(limit), util.StrToUInt(uId))
	response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
}
