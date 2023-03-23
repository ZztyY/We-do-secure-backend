package interfaces

import (
	"We-do-secure/interfaces/cus"
	"We-do-secure/interfaces/home"
	"We-do-secure/interfaces/invoice"
	"We-do-secure/interfaces/pol"
	"We-do-secure/interfaces/user"
	"We-do-secure/interfaces/vehicle"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	group := r.Group("/WDS")

	group.POST("/login", user.UserLogin)
	group.POST("/register", user.UserRegister)
	group.GET("/reset/password", user.ResetPassword)
	group.GET("/user/list", user.UserList)
	group.GET("/user/get", user.GetUser)

	group.POST("/home/add", home.AddHome)
	group.GET("/home/list", home.HomeList)
	group.GET("/home/get", home.GetHome)
	group.GET("/home/user/list", home.UserHomeList)

	group.POST("/cus/edit", cus.EditCus)
	group.GET("/cus/get", cus.GetCus)

	group.POST("/vehicle/edit", vehicle.EditVehicle)
	group.GET("/vehicle/get", vehicle.GetVehicle)
	// user related vehicles list
	group.GET("/vehicle/user/list", vehicle.UserVehicleList)

	group.POST("/driver/edit", vehicle.EditDriver)
	group.GET("/driver/get", vehicle.GetDriver)
	// vehicle related drivers list
	group.GET("/driver/vehicle/list", vehicle.VehicleDriverList)

	group.GET("/pol/user/list", pol.UserPolList)
	group.POST("/pol/edit", pol.EditPol)
	group.GET("/pol/get", pol.GetPol)

	group.GET("/invoice/user/list", invoice.UserInvoiceList)
	group.GET("/invoice/get", invoice.GetInvoice)
	group.POST("/user/update/payment", invoice.UpdatePaymentMethod)
	group.POST("/invoice/make/payment", invoice.MakePayment)
	group.GET("/payment/list", invoice.PaymentList)
}
