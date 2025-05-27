package service

import (
	"ksd-social-api/modules/user/model"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/utils"
)

/**
 * 修改密码
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) UpdateUserPassword(ctx context.UserPasswordContext) bool {
	user := userMapper.GetUserByID(ctx.UserId, ctx.SystemId)
	md5Password := utils.Md5Slat(ctx.Password, user.Slat)
	flag, _ := userMapper.UpdateUserPassword(ctx.UserId, ctx.SystemId, md5Password)
	return flag
}

/**
 * 修改成为作者
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) UpdateUserAuthor(userId uint64, systemId int) bool {
	flag, _ := userMapper.UpdateUserAuthor(userId, systemId)
	return flag
}

/**
 * 用户收益设置
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) UpdateUserBank(ctx context.UserBankContext) bool {
	flag, _ := userMapper.UpdateUserBank(ctx)
	return flag
}

/**
 * 账户设置
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) UpdateUserInfo(ctx context.UserInfoContext) bool {
	var user model.User
	utils.CopyProperties(&user, ctx)
	return userMapper.UpdateUserCenterInfo(&user)
}
