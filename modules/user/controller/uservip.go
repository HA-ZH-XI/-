package web

import (
	"ksd-social-api/commons/base/controller"
	"ksd-social-api/modules/user/service"
)

/**
 * 用户
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserVipController struct {
	controller.BaseController
}

var userVipService = service.UserVipService{}

/**
 * 查看用户vip列表
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (web *UserVipController) FindUserVipCardList() {
	list := userVipService.FindUserVipCardList(web.GetSystemId())
	web.Ok(list)
}
