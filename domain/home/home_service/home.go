package home_service

import "We-do-secure/domain/home/home_model"

func FindHomeList(offset int, limit int) ([]home_model.Home, int) {
	return home_model.FindHomeList(offset, limit)
}
