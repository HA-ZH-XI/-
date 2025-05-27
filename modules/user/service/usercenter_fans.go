package service

import (
	"ksd-social-api/commons/page"
	"ksd-social-api/modules/user/model/context"
)

/**
* 粉丝列表
* @author feige
* @date 2023-11-03
* @version 1.0
* @desc
 */
func (mapper *UserFocusService) FindUserFocusFansPage(ctx context.UserFocusPageContext) *page.Page {
	userPage, _ := userFocusMapper.FindUserFocusFansPage(ctx)
	return userPage
}

/**
* 关注列表
* @author feige
* @date 2023-11-03
* @version 1.0
* @desc
 */
func (mapper *UserFocusService) FindUserFocusGzPage(ctx context.UserFocusPageContext) *page.Page {
	userPage, _ := userFocusMapper.FindUserFocusGzPage(ctx)
	return userPage
}
