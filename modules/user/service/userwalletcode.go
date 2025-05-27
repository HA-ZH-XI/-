package service

import (
	"ksd-social-api/commons/base/service"
	"ksd-social-api/modules/user/model"
	"strconv"
)

type UserWalletCodeService struct {
	service.BaseService
}

/**
 * 用户身份兑换码
 * @author feige
 * @date 2023-10-10
 * @version 1.0 */
func (service *UserWalletCodeService) SaveUserWalletCode() bool {
	userWalletCodes := []model.UserWalletCode{}
	for i := 1; i <= 100; i++ {
		userWalletCode := model.UserWalletCode{}
		userWalletCode.Code = service.GetSnowWorkerIdString(int64(i))
		userWalletCode.Cron = 20 * (i % 10)
		userWalletCode.Mark = 0
		userWalletCodes = append(userWalletCodes, userWalletCode)
	}
	userWalletCodeMapper.SaveUserWalletCode(userWalletCodes)
	return true
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 开始兑换
 */
func (service *UserWalletCodeService) DuihuanUserWallet(code string, userId uint64, systemId int) bool {
	courseCode := userWalletCodeMapper.CountUserWalletCode(code, systemId)
	// 如果还没兑换，就开始兑换，生成充值订单，同时删除订单
	if courseCode > 0 {
		// 删除兑换的订单
		userWalletCodeMapper.UpdateUserWalletCode(code, systemId)
		// 开始保存订单
		service.saveUserWalletCodeOrder(code, userId, systemId)
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
func (service *UserWalletCodeService) saveUserWalletCodeOrder(code string, userId uint64, systemId int) {
	// 1： 判断用户是否购买此课程，如果购买过了就就不要去保存订单
	isBuyFlag := userBuyVipMapper.CountUserBuyVipByCode(code, systemId)
	if isBuyFlag == 0 {
		// 这里就是解析下单的附属参数
		// 根据课程id查询课程信息
		userVo := userMapper.GetUserByID(userId, systemId)
		userWalletRecords := model.UserWalletRecords{}
		userWalletCode := userWalletCodeMapper.GetUserWalletDetail(code, systemId)
		// 处理返回结果
		userWalletRecords.Title = "学习币充值"
		userWalletRecords.Price = strconv.Itoa(userWalletCode.Cron)

		userWalletRecords.SystemId = systemId
		userWalletRecords.UserId = userId
		userWalletRecords.Username = userVo.UserName
		userWalletRecords.Avatar = userVo.Avatar
		userWalletRecords.Uuid = userVo.Uuid
		userWalletRecords.Phone = userVo.Telephone
		userWalletRecords.Address = userVo.Address
		userWalletRecords.Nickname = userVo.NickName

		userWalletRecords.PayMethod = 3
		userWalletRecords.PayMethodName = "兑换码"
		userWalletRecords.Description = "兑换码充值"
		userWalletRecords.Tradeno = service.GetSnowWorkerIdString(19)
		userWalletRecords.Orderno = service.GetSnowWorkerIdString(20)
		userWalletRecords.OrderJson = ""

		// 保存订单的接口
		userWalletMapper.SaveUserWalletRecords(userWalletRecords)
	}
}
