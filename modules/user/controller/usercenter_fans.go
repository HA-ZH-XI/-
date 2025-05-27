package web

import "ksd-social-api/modules/user/model/context"

/**
* 粉丝列表
* @author feige
* @date 2023-11-03
* @version 1.0
* @desc
 */
func (web *UserCenterController) FindUserFocusPage() {
	userFocusPageContext := context.UserFocusPageContext{}
	web.BindJSON(&userFocusPageContext)
	userFocusPageContext.UserId = web.GetUserId()
	userFocusPageContext.SystemId = web.GetSystemId()
	web.Ok(userFocusService.FindUserFocusFansPage(userFocusPageContext))
}

/**
* 关注列表
* @author feige
* @date 2023-11-03
* @version 1.0
* @desc
 */
func (web *UserCenterController) FindUserFocusGzPage() {
	userFocusPageContext := context.UserFocusPageContext{}
	web.BindJSON(&userFocusPageContext)
	userFocusPageContext.UserId = web.GetUserId()
	userFocusPageContext.SystemId = web.GetSystemId()
	web.Ok(userFocusService.FindUserFocusGzPage(userFocusPageContext))
}
