package web

import "ksd-social-api/modules/user/model/context"

/**
 * @desc 我的消息
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) FindUserMeMessageRecords() {
	messageContext := context.MessageContext{}
	web.BindJSON(&messageContext)
	messageContext.SystemId = web.GetSystemId()
	messageContext.UserId = web.GetUserId()
	page, _ := userCenterService.FindMessageMePage(messageContext)
	web.Ok(page)
}

/**
 * @desc 系统消息
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) FindUserSystemMessageRecords() {
	messageContext := context.MessagePageContext{}
	web.BindJSON(&messageContext)
	messageContext.SystemId = web.GetSystemId()
	page, _ := userCenterService.FindMessageSystemPage(messageContext)
	web.Ok(page)
}

/**
 * 根据ID删除消息
 * @author feige
 * @date 2023-11-03
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) DelMessageMeById() {
	id, _ := web.GetInt("id")
	web.Ok(messageMeService.DelMessageMeById(id, web.GetUserId(), web.GetSystemId()))
}

/**
 * 文章评论已读和未读
 * @author feige
 * @date 2023-11-03
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) UpdateMessageMeMarkById() {
	id, _ := web.GetInt("id")
	web.Ok(messageMeService.UpdateMessageMeMarkById(id, web.GetUserId(), web.GetSystemId()))
}

/**
 * 一键已读
 * @author feige
 * @date 2023-11-03
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) UpdateMessageMeMarkByUserId() {
	web.Ok(messageMeService.UpdateMessageMeMarkByUserId(web.GetUserId(), web.GetSystemId()))
}
