package interfaces

import (
	"We-do-secure/interfaces/home"
	"We-do-secure/interfaces/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	group := r.Group("/WDS")

	group.POST("/login", user.UserLogin)
	group.POST("/register", user.UserRegister)
	group.GET("/reset/password", user.ResetPassword)
	group.GET("/user/list", user.UserList)
	group.POST("/home/add", home.AddHome)
	group.GET("/home/list", home.HomeList)
	group.GET("/home/get", home.GetHome)
	group.GET("/home/user/list", home.UserHomeList)
}
