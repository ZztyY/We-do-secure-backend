package user_service

import "We-do-secure/domain/user/user_model"

func UserLogin(userName string, password string) *user_model.User {
	filter := make(map[string]interface{})
	filter["user_name"] = userName
	filter["password"] = password
	return user_model.FindUser(filter)
}

func FindUserList(offset int, limit int) ([]user_model.User, int) {
	return user_model.FindUserList(offset, limit)
}
