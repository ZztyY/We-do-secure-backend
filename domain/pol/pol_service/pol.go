package pol_service

import (
	"We-do-secure/domain/pol/pol_model"
)

func GetPolHomeByHId(hId uint) *pol_model.PolHome {
	filter := make(map[string]interface{})
	filter["h_id"] = hId
	return pol_model.FindPolHome(filter)
}

func GetPolVehByVId(vId uint) *pol_model.PolVeh {
	filter := make(map[string]interface{})
	filter["v_id"] = vId
	return pol_model.FindPolVeh(filter)
}
