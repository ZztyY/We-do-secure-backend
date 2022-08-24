package cus

import (
	"We-do-secure/domain/cus/cus_model"
	"We-do-secure/interfaces/errorcode"
	"We-do-secure/interfaces/response"
	"We-do-secure/util"
	"github.com/gin-gonic/gin"
)

func EditCus(c *gin.Context) {
	fName := c.PostForm("fname")
	lName := c.PostForm("lname")
	street := c.PostForm("street")
	city := c.PostForm("city")
	state := c.PostForm("state")
	zipcode := c.PostForm("zipcode")
	gender := c.PostForm("gender")
	marStatus := c.PostForm("mar_status")
	uId := c.PostForm("uid")

	if fName == "" || lName == "" || street == "" || city == "" || state == "" || zipcode == "" || marStatus == "" || uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	if cus_model.GetCusByUid(util.StrToUInt(uId)) == nil {
		cus := cus_model.Cus{}
		cus.UID = util.StrToUInt(uId)
		cus.FName = fName
		cus.LName = lName
		cus.Street = street
		cus.City = city
		cus.State = state
		cus.Zipcode = util.StrToInt(zipcode)
		cus.Gender = gender
		cus.MarStatus = marStatus
		cus_model.CreateCus(&cus)
		response.SendSuccess(c, cus)
	} else {
		cus := cus_model.GetCusByUid(util.StrToUInt(uId))
		cus.FName = fName
		cus.LName = lName
		cus.Street = street
		cus.City = city
		cus.State = state
		cus.Zipcode = util.StrToInt(zipcode)
		cus.Gender = gender
		cus.MarStatus = marStatus
		cus_model.UpdateCus(cus)
		response.SendSuccess(c, cus)
	}
}

func GetCus(c *gin.Context) {
	uId := c.DefaultQuery("uid", "")

	if uId == "" || cus_model.GetCusByUid(util.StrToUInt(uId)) == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "error or missing params", nil)
		return
	}

	cus := cus_model.GetCusByUid(util.StrToUInt(uId))
	response.SendSuccess(c, cus)
}
