package web

import (
	"ksd-social-api/commons/base/controller"
	"ksd-social-api/modules/user/service"
)

type UserVipCodeController struct {
	controller.BaseController
}

var userVipCodeService = service.UserVipCodeService{}

/**
 * 创建兑换码
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (web *UserVipCodeController) SaveUserVipCode() {
	userVipCodeService.SaveUserVipCode()
	web.Ok("success")
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 查询课程和code是否存在
 */
func (web *UserVipCodeController) DuihuanUserVIP() {
	code := web.GetString("code")
	if len(code) == 0 {
		web.FailCodeMsg(999, "请输入兑换码！")
		return
	}
	userVipCode := userVipCodeService.DuihuanUserVIP(code, web.GetUserId(), web.GetSystemId())
	// 如果还没兑换，就开始兑换，生成课程订单，同时删除订单
	if !userVipCode {
		web.FailCodeMsg(999, "该兑换码不存在或者已被使用！")
		return
	}
	web.Ok("success")
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 查询课程和code是否存在
 */
func (web *UserVipCodeController) QueryInfoByCode() {
	code := web.GetString("code")
	if len(code) == 0 {
		web.FailCodeMsg(999, "请输入兑换码！")
		return
	}
	userVipCode := userVipService.GetUserVipByCode(code, web.GetSystemId())
	// 如果还没兑换，就开始兑换，生成课程订单，同时删除订单
	if nil == userVipCode {
		web.FailCodeMsg(999, "没有该兑换码信息")
		return
	}
	web.Ok(userVipCode)
}
