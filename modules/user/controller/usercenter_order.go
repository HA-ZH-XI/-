package web

import (
	"ksd-social-api/modules/user/model/context"
)

/**
 * 文章订单
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyUserBuyBbsOrderPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyUserBuyBbsOrderPage(userCenterContext)
	web.Ok(p)
}

/**
 * vip订单
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyUserBuyVipOrderPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyUserBuyVipOrderPage(userCenterContext)
	web.Ok(p)
}

/**
 * 课程订单
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web UserCenterController) FindMyCoursesOrderPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyCoursesOrderPage(userCenterContext)
	web.Ok(p)
}
