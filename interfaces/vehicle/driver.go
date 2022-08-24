package vehicle

import (
	"We-do-secure/domain/vehicle/vehicle_model"
	"We-do-secure/interfaces/errorcode"
	"We-do-secure/interfaces/response"
	"We-do-secure/util"
	"github.com/gin-gonic/gin"
)

func EditDriver(c *gin.Context) {
	dId := c.PostForm("did")
	fName := c.PostForm("fName")
	lName := c.PostForm("lName")
	lNum := c.PostForm("lNum")
	birth := c.PostForm("birth")
	vId := c.PostForm("vid")

	if fName == "" || lName == "" || lNum == "" || birth == "" || vId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	if vehicle_model.GetVehicle(util.StrToUInt(vId)) == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "vehicle doesn't exist", nil)
		return
	}

	if dId == "" {
		driver := vehicle_model.Driver{}
		driver.FName = fName
		driver.LName = lName
		driver.LNum = lNum
		driver.Birth = util.JSONDetailTime{Time: util.StrToTimeTime(birth)}
		vehicle_model.CreateDriver(&driver)

		vehDri := vehicle_model.VehDri{}
		vehDri.VID = util.StrToUInt(vId)
		vehDri.DID = driver.DID
		vehicle_model.CreateVehDri(&vehDri)
		response.SendSuccess(c, driver)
	} else if vehicle_model.GetDriver(util.StrToUInt(dId)) == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "Invalid DID", nil)
		return
	} else {
		driver := vehicle_model.GetDriver(util.StrToUInt(dId))
		driver.FName = fName
		driver.LName = lName
		driver.LNum = lNum
		driver.Birth = util.JSONDetailTime{Time: util.StrToTimeTime(birth)}
		vehicle_model.UpdateDriver(driver)
		response.SendSuccess(c, driver)
	}
}

func GetDriver(c *gin.Context) {
	dId := c.DefaultQuery("did", "")
	if dId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "please input driver id", nil)
		return
	}
	driver := vehicle_model.GetDriver(util.StrToUInt(dId))
	if driver == nil {
		response.SendError(c, errorcode.CODE_INVALID_INPUT, "driver doesn't exist", nil)
		return
	}
	response.SendSuccess(c, driver)
}

func VehicleDriverList(c *gin.Context) {
	vId := c.DefaultQuery("vid", "")
	if vId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "please input vehicle id", nil)
		return
	}
	if vehicle_model.GetVehicle(util.StrToUInt(vId)) == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "vehicle doesn't exist", nil)
		return
	}

	vehDriList := vehicle_model.FindVehDriListByVId(util.StrToUInt(vId))

	var dIdList []uint
	for _, v := range vehDriList {
		dIdList = append(dIdList, v.DID)
	}

	driverList := vehicle_model.FindDriverListByDIdList(dIdList)
	response.SendSuccess(c, driverList)
}
