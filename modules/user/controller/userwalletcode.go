package web

import (
	"ksd-social-api/commons/base/controller"
	"ksd-social-api/modules/user/service"
)

type UserWalletCodeController struct {
	controller.BaseController
}

var userWalletCodeService = service.UserWalletCodeService{}

/**
 * 创建兑换码
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (web *UserWalletCodeController) SaveUserWalletCode() {
	userWalletCodeService.SaveUserWalletCode()
	web.Ok("success")
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 查询课程和code是否存在
 */
func (web *UserWalletCodeController) DuihuanUserWallet() {
	code := web.GetString("code")
	if len(code) == 0 {
		web.FailCodeMsg(999, "请输入兑换码！")
		return
	}
	userVipCode := userWalletCodeService.DuihuanUserWallet(code, web.GetUserId(), web.GetSystemId())
	// 如果还没兑换，就开始兑换，生成课程订单，同时删除订单
	if !userVipCode {
		web.FailCodeMsg(999, "该兑换码不存在或者已被使用！")
		return
	}
	web.Ok("success")
}
