package service

import (
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/model/vo"
)

type UserCenterService struct {
}

/**
 * 查询用户中心首页总览的数据信息
 * @author feige
 * @date 2024-01-08
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) CountUserRelationState(userId uint64, systemId int) []*vo.UserStateCountAllVo {
	return userCenterMapper.CountUserRelationState(userId, systemId)
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  统计用户的每日发文章数量，学习课程，新增粉丝
 */
func (service *UserCenterService) CountUserModelState(ctx context.UserStateModelContext) []*vo.UserStateModelDataVo {
	return userCenterMapper.CountUserModelState(ctx.UserId, ctx.Day, ctx.SystemId)
}
