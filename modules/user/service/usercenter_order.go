package service

import (
	"ksd-social-api/commons/page"
	cmodel "ksd-social-api/modules/course/model"
	"ksd-social-api/modules/user/model"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/model/vo"
	"ksd-social-api/utils"
)

/**
 * vip订单
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyUserBuyVipOrderPage(ctx context.UserCenterContext) *page.Page {
	userBuyVipPage, _ := userBuyVipMapper.FindMyUserBuyVipOrderPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	userBuyVipList := userBuyVipPage.Records.([]model.UserBuyVip)
	if userBuyVipList != nil {
		userBuyVipVos := []vo.UserBuyVipVo{}
		for _, userBuyVip := range userBuyVipList {
			userBuyVipVo := vo.UserBuyVipVo{}
			utils.CopyProperties(&userBuyVipVo, userBuyVip)
			userBuyVipVos = append(userBuyVipVos, userBuyVipVo)
		}
		userBuyVipPage.Records = userBuyVipVos
	}
	return userBuyVipPage
}

/**
 * 文章订单
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyUserBuyBbsOrderPage(ctx context.UserCenterContext) *page.Page {
	userBuyBbsPage, _ := userBuyBbsMapper.FindMyUserBuyBbsOrderPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	userBuyBbsList := userBuyBbsPage.Records.([]model.UserBuyBbs)
	if userBuyBbsList != nil {
		userBuyBbsVos := []vo.UserBuyBbsVo{}
		for _, userBuyBbs := range userBuyBbsList {
			userBuyBbsVo := vo.UserBuyBbsVo{}
			utils.CopyProperties(&userBuyBbsVo, userBuyBbs)
			userBuyBbsVos = append(userBuyBbsVos, userBuyBbsVo)
		}
		userBuyBbsPage.Records = userBuyBbsVos
	}
	return userBuyBbsPage
}

/**
 * 课程订单
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyCoursesOrderPage(ctx context.UserCenterContext) *page.Page {
	userBuyCoursePage, _ := userBuyCourseMapper.FindMyCoursesOrderPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	userBuyCourseList := userBuyCoursePage.Records.([]cmodel.UserBuyCourse)
	if userBuyCourseList != nil {
		userBuyCourseVos := []vo.UserBuyCourseVo{}
		for _, userBuyCourse := range userBuyCourseList {
			userBuyCourseVo := vo.UserBuyCourseVo{}
			utils.CopyProperties(&userBuyCourseVo, userBuyCourse)
			userBuyCourseVos = append(userBuyCourseVos, userBuyCourseVo)
		}
		userBuyCoursePage.Records = userBuyCourseVos
	}
	return userBuyCoursePage
}
