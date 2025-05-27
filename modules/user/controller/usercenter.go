package web

import (
	"ksd-social-api/commons/base/controller"
	service2 "ksd-social-api/modules/msg/service"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/service"
)

/**
 * 用户
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserCenterController struct {
	controller.BaseController
}

var userCenterService = service.UserCenterService{}
var messageMeService = service2.MessageMeService{}

/**
 * 查询用户中心首页总览的数据信息
 * @author feige
 * @date 2024-01-08
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) CountUserRelationState() {
	web.Ok(userCenterService.CountUserRelationState(web.GetUserId(), web.GetSystemId()))
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  统计用户的每日发文章数量，学习课程，新增粉丝
 */
func (web *UserCenterController) CountUserModelState() {
	userStateModelContext := context.UserStateModelContext{}
	web.BindJSON(&userStateModelContext)
	userStateModelContext.UserId = web.GetUserId()
	userStateModelContext.SystemId = web.GetSystemId()
	web.Ok(userCenterService.CountUserModelState(userStateModelContext))
}
