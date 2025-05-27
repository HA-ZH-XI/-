package service

import (
	"ksd-social-api/modules/user/model"
	"ksd-social-api/modules/user/model/context"
)

/**
 * 用户关注
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserFocusService struct {
}

/**
 * 关注用户和取消关注和再关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusService) SaveCancelUserFocus(ctx context.UserFocusContext) int64 {
	isFocus := userFocusMapper.CountUserFocus(ctx.UserId, ctx.FocusId.Uint64(), ctx.SystemId)
	// 关注
	if isFocus == -1 {
		var userFocus model.UserFocus
		userFocus.UserId = ctx.UserId
		userFocus.Uuid = ctx.Uuid
		userFocus.Nickname = ctx.Nickname
		userFocus.Avatar = ctx.Avatar
		userFocus.FocusId = ctx.FocusId.Uint64()
		userFocus.FocusUuid = ctx.FocusUuid
		userFocus.FocusNickname = ctx.FocusNickname
		userFocus.FocusAvatar = ctx.FocusAvatar
		userFocus.IsFocus = 1
		userFocusMapper.SaveUserFocus(userFocus)
		return 1
	}

	// 再关注
	if isFocus == 0 {
		userFocusMapper.RecoverUserFocus(ctx.UserId, ctx.FocusId.Uint64(), ctx.SystemId)
		return 1
	}

	// 取消关注
	if isFocus == 1 {
		userFocusMapper.CancelUserFocus(ctx.UserId, ctx.FocusId.Uint64(), ctx.SystemId)
		return 0
	}

	return 0
}

/**
 * 查询是否关注
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserFocusService) GetUserFocus(userId uint64, focusId uint64, systemId int) int64 {
	return userFocusMapper.CountUserFocus(userId, focusId, systemId)
}

/**
 * 删除关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusService) DelUserFocus(ctx context.UserFocusDelContext) bool {
	return userFocusMapper.DelUserFocus(ctx.Id.Uint64(), ctx.UserId, ctx.FocusId.Uint64(), ctx.SystemId)
}

/**
 * 取消关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusService) CancelUserFocus(ctx context.UserFocusDelContext) bool {
	return userFocusMapper.CancelUserFanGzFocus(ctx.Id.Uint64(), ctx.UserId, ctx.FocusId.Uint64(), ctx.SystemId)
}

/**
 * 移除粉丝
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusService) DelUserFansFocus(ctx context.UserFocusDelFansContext) bool {
	return userFocusMapper.DelUserFanGzFocus(ctx.Id.Uint64(), ctx.FocusId, ctx.UserId, ctx.SystemId)
}
