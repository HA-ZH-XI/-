package service

import (
	"ksd-social-api/modules/common/pay/context"
	"ksd-social-api/modules/user/model"
	"ksd-social-api/utils"
)

/**
 * 用户购买会员身份
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserBuyVipService struct {
}

/**
 * 查询用户是否购买会员身份
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyVipService) SaveUserBuyVipService(ctx context.UserBuyVipContext) bool {
	var userBuyVip model.UserBuyVip
	utils.CopyProperties(&userBuyVip, ctx)
	return userBuyVipMapper.SaveUserBuyVip(userBuyVip)
}

/**
 * 查询用户是否购买会员身份
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyVipService) CountUserBuyVipByOutTradeNo(OutTradeNo string, systemId int) int {
	return userBuyVipMapper.CountUserBuyVipByOutTradeNo(OutTradeNo, systemId)
}
