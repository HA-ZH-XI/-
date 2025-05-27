package web

import (
	"ksd-social-api/commons/base/controller"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/service"
)

/**
 * 用户关注
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserFocusController struct {
	controller.BaseController
}

var userFocusService = service.UserFocusService{}

/**
 * 关注用户
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (web *UserFocusController) SaveCancelUserFocus() {
	userFocusContext := context.UserFocusContext{}
	web.BindJSON(&userFocusContext)
	// 登录用户---关注用户
	userFocusContext.UserId = web.GetUserId()
	userFocusContext.Uuid = web.GetUuid()
	userFocusContext.Avatar = web.GetUserAvatar()
	userFocusContext.Nickname = web.GetUserName()
	userFocusContext.SystemId = web.GetSystemId()
	if userFocusContext.FocusId.Uint64() == userFocusContext.UserId {
		web.FailCodeMsg(601, "自己不能关注自己！")
		return
	}

	// 关注，取消 和 再关注
	focus := userFocusService.SaveCancelUserFocus(userFocusContext)
	web.Ok(focus)
}

/**
 * 删除关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (web *UserFocusController) CancelUserFocus() {
	userFocusDelContext := context.UserFocusDelContext{}
	web.BindJSON(&userFocusDelContext)
	userFocusDelContext.UserId = web.GetUserId()
	userFocusDelContext.SystemId = web.GetSystemId()
	focus := userFocusService.CancelUserFocus(userFocusDelContext)
	if !focus {
		web.FailCodeMsg(1001, "取消关注失败!")
		return
	}
	web.Ok(focus)
}

/**
 * 移除粉丝
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (web *UserFocusController) CancelUserFansFocus() {
	userFocusDelFansContext := context.UserFocusDelFansContext{}
	web.BindJSON(&userFocusDelFansContext)
	userFocusDelFansContext.UserId = web.GetUserId()
	userFocusDelFansContext.SystemId = web.GetSystemId()
	focus := userFocusService.DelUserFansFocus(userFocusDelFansContext)
	if !focus {
		web.FailCodeMsg(1001, "取消关注失败!")
		return
	}
	web.Ok(focus)
}

/**
 * 删除关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (web *UserFocusController) DelUserFocus() {
	userFocusDelContext := context.UserFocusDelContext{}
	web.BindJSON(&userFocusDelContext)
	userFocusDelContext.UserId = web.GetUserId()
	userFocusDelContext.SystemId = web.GetSystemId()
	focus := userFocusService.DelUserFocus(userFocusDelContext)
	if !focus {
		web.FailCodeMsg(1001, "删除关注失败!")
		return
	}
	web.Ok(focus)
}
