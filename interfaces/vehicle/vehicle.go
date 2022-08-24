package vehicle

import (
	"We-do-secure/domain/vehicle/vehicle_model"
	"We-do-secure/interfaces/errorcode"
	"We-do-secure/interfaces/response"
	"We-do-secure/util"
	"github.com/gin-gonic/gin"
)

func EditVehicle(c *gin.Context) {
	vId := c.PostForm("vid")
	vIn := c.PostForm("vin")
	vMMYear := c.PostForm("vmmyear")
	vStatus := c.PostForm("vstatus")
	uId := c.PostForm("uid")

	if vIn == "" || vMMYear == "" || vStatus == "" || uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	if vId == "" {
		vehicle := vehicle_model.Vehicle{}
		vehicle.VIN = vIn
		vehicle.VMMYear = util.StrToInt(vMMYear)
		vehicle.VStatus = vStatus
		vehicle.UID = util.StrToUInt(uId)
		vehicle_model.CreateVehicle(&vehicle)
		response.SendSuccess(c, vehicle)
	} else if vehicle_model.GetVehicle(util.StrToUInt(vId)) == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "Invalid VID", nil)
		return
	} else {
		vehicle := vehicle_model.GetVehicle(util.StrToUInt(vId))

		vehicle.VIN = vIn
		vehicle.VMMYear = util.StrToInt(vMMYear)
		vehicle.VStatus = vStatus
		vehicle.UID = util.StrToUInt(uId)

		vehicle_model.UpdateVehicle(vehicle)
		response.SendSuccess(c, vehicle)
	}
}

func GetVehicle(c *gin.Context) {
	vId := c.DefaultQuery("vid", "")
	if vId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "please input vehicle id", nil)
		return
	}
	vehicle := vehicle_model.GetVehicle(util.StrToUInt(vId))
	if vehicle == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "vehicle doesn't exist", nil)
		return
	}
	response.SendSuccess(c, vehicle)
}

func UserVehicleList(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	uId := c.DefaultQuery("uid", "")
	if uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}
	list, count := vehicle_model.FindUserVehicleList(util.StrToInt(offset), util.StrToInt(limit), util.StrToUInt(uId))
	response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
}
