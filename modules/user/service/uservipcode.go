package service

import (
	"ksd-social-api/commons/base/service"
	"ksd-social-api/modules/user/model"
)

type UserVipCodeService struct {
	service.BaseService
}

/**
 * 查询兑换卡信息
 * @author feige
 * @date 2023-12-26
 * @version 1.0
 * @desc
 */
func (service *UserVipCodeService) QueryInfoByCode(code string, systemId int) *model.UserVipCode {
	return userVipCodeMapper.GetUserVipByCode(code, systemId)
}

/**
 * 用户身份兑换码
 * @author feige
 * @date 2023-10-10
 * @version 1.0 */
func (service *UserVipCodeService) SaveUserVipCode() bool {
	userVipCodes := []model.UserVipCode{}
	for i := 0; i < 100; i++ {
		userVipCode := model.UserVipCode{}
		userVipCode.Code = service.GetSnowWorkerIdString(int64(i))
		userVipCode.VipId = i%6 + 1
		userVipCode.Mark = 0
		userVipCodes = append(userVipCodes, userVipCode)
	}
	userVipCodeMapper.SaveUserVipCode(userVipCodes)
	return true
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 查询会员升级和code是否存在
 */
func (service *UserVipCodeService) DuihuanUserVIP(code string, userId uint64, systemId int) bool {
	courseCode := userVipCodeMapper.CountUserVipCode(code, systemId)
	// 如果还没兑换，就开始兑换，生成会员升级订单，同时删除订单
	if courseCode > 0 {
		// 删除兑换的订单
		userVipCodeMapper.UpdateUserVipCode(code, systemId)
		// 根据code查询
		userVipCode := userVipCodeMapper.GetUserVipByCode(code, systemId)
		// 开始保存订单
		service.saveUserVipCodeOrder(code, userVipCode.VipId, userId, systemId)
		// 兑换成功
		return true
	}
	return false
}

/**
 * 保存用户身份和兑换码
 * @author feige
 * @date 2023-12-06
 * @version 1.0
 * @desc
 */
func (service *UserVipCodeService) saveUserVipCodeOrder(code string, vipId int, userId uint64, systemId int) {
	// 1： 判断用户是否购买此会员升级，如果购买过了就就不要去保存订单
	isBuyFlag := userBuyVipMapper.CountUserBuyVipByCode(code, systemId)
	if isBuyFlag == 0 {
		// 这里就是解析下单的附属参数
		// 根据会员升级id查询会员升级信息
		userVipVo := userVipMapper.GetUserVipDetail(vipId, systemId)
		userVo := userMapper.GetUserByID(userId, systemId)

		userBuyVip := model.UserBuyVip{}
		// 处理返回结果
		userBuyVip.VipId = vipId
		userBuyVip.Title = userVipVo.Title
		userBuyVip.Description = userVipVo.Note

		userBuyVip.UserId = userId
		userBuyVip.Username = userVo.UserName
		userBuyVip.Avatar = userVo.Avatar
		userBuyVip.Uuid = userVo.Uuid
		userBuyVip.Phone = userVo.Telephone
		userBuyVip.Address = userVo.Address
		userBuyVip.Nickname = userVo.NickName
		userBuyVip.SystemId = systemId

		userBuyVip.Code = code
		userBuyVip.PayMethod = 3
		userBuyVip.PayMethodName = "兑换码"
		userBuyVip.Tradeno = service.GetSnowWorkerIdString(9)
		userBuyVip.Orderno = service.GetSnowWorkerIdString(10)
		userBuyVip.OrderJson = ""
		userBuyVip.Price = userVipVo.Realprice
		// 保存订单的接口
		userBuyVipMapper.SaveUserBuyVip(userBuyVip)
	}
}
