package service

import (
	"ksd-social-api/commons/page"
	"ksd-social-api/modules/user/model/context"
)

/**
 * 我的消息
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMessageMePage(ctx context.MessageContext) (p *page.Page, err error) {
	return messageMeMapper.FindMessageMePage(ctx.Mtype, ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
}

/**
 * 系统消息
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMessageSystemPage(ctx context.MessagePageContext) (p *page.Page, err error) {
	return messageSystemMapper.FindMessageSystemPage(ctx.SystemId, ctx.SType, ctx.PageNo, ctx.PageSize)
}
