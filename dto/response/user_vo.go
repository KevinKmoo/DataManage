package response

import "github.com/kmookay/MyBittDataManage/model"

type LoginVo struct {
	model.User
	Roles string `json:"roles"`
}
