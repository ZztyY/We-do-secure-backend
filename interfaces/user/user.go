package user

import (
	"We-do-secure/domain/user/user_model"
	"We-do-secure/domain/user/user_service"
	"We-do-secure/interfaces/errorcode"
	"We-do-secure/interfaces/response"
	"We-do-secure/util"
	"github.com/gin-gonic/gin"
	"time"
)

func UserLogin(c *gin.Context) {
	userName := c.PostForm("user_name")
	password := c.PostForm("password")

	if userName == "" || password == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "please input username or password", nil)
		return
	}
	user := user_service.UserLogin(userName, password)
	if user == nil {
		response.SendError(c, errorcode.CODE_AUTH_CHECK_TOKEN_FAIL, "wrong username or password", nil)
		return
	}
	user.Token = util.RandomStr(10)
	user.TokenExpiredAt = util.JSONDetailTime{Time: time.Now().Add(time.Hour * 24 * 10)}
	user_model.UpdateUser(user)
	response.SendSuccess(c, user)
}

func GetUser(c *gin.Context) {
	uId := c.DefaultQuery("uid", "")
	if uId == "" {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "missing params", nil)
		return
	}

	if user_model.GetUser(util.StrToUInt(uId)) == nil {
		response.SendError(c, errorcode.CODE_PARAMS_INVALID, "no such user", nil)
		return
	}

	user := user_model.GetUser(util.StrToUInt(uId))

	response.SendSuccess(c, user)
}

func UserRegister(c *gin.Context) {
	username := c.PostForm("user_name")
	password := c.PostForm("password")
	if username == "" || password == "" {
		response.SendParamsError(c, nil)
		return
	}
	filter := make(map[string]interface{})
	filter["user_name"] = username
	if user_model.FindUser(filter) != nil {
		response.SendError(c, errorcode.CODE_USER_NAME_EXIST, "username exists!", nil)
		return
	}
	if len(password) < 6 {
		response.SendError(c, errorcode.CODE_USER_PASSWORD_FORMAT, "password too short!", nil)
		return
	}
	user := user_model.User{}
	user.Password = password
	user.UserName = username
	user.Token = util.RandomStr(10)
	user.TokenExpiredAt = util.JSONDetailTime{Time: time.Now().Add(time.Hour * 24 * 10)}
	user_model.CreateUser(&user)
	response.SendSuccess(c, user)
}

func ResetPassword(c *gin.Context) {
	userId := c.DefaultQuery("user_id", "")
	password := c.DefaultQuery("password", "")
	if userId == "" {
		response.SendParamsError(c, nil)
		return
	}
	if len(password) < 6 {
		response.SendError(c, errorcode.CODE_USER_PASSWORD_FORMAT, "password too short!", nil)
		return
	}
	user := user_model.GetUser(util.StrToUInt(userId))
	if password == user.Password {
		response.SendError(c, errorcode.CODE_USER_PASSWORD_FORMAT, "password can not be same as before!", nil)
		return
	}
	user.Password = password
	user_model.UpdateUser(user)
	response.SendSuccess(c, "success")
}

func UserList(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	list, count := user_service.FindUserList(util.StrToInt(offset), util.StrToInt(limit))
	response.SendSuccess(c, map[string]interface{}{"count": count, "list": list})
}
