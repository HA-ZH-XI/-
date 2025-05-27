package service

import (
	"ksd-social-api/commons/base/service"
	"ksd-social-api/commons/page"
	"ksd-social-api/modules/common/pay/context"
	"ksd-social-api/modules/user/model"
	context2 "ksd-social-api/modules/user/model/context"
	vo2 "ksd-social-api/modules/user/model/vo"
	"ksd-social-api/utils"
)

type UserWalletService struct {
	service.BaseService
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 充值订单
 */
func (mapper *UserWalletService) SaveUserWalletRecords(ctx context.UserWalletRecordsContext) bool {
	var userWalletRecords model.UserWalletRecords
	utils.CopyProperties(&userWalletRecords, ctx)
	return userWalletMapper.SaveUserWalletRecords(userWalletRecords)
}

/**
 * 统计订单是否支付--充值
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserWalletService) CountUserWalletRecordsByOutTradeNo(OutTradeNo string, systemId int) int {
	return userWalletMapper.CountUserWalletRecordsByOutTradeNo(OutTradeNo, systemId)
}

/**
 * @desc 查询文章列表信息
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserWalletService) FindUserWalletRecords(ctx context2.UserWalletContext) *page.Page {
	userWalletRecprdsPage, _ := userWalletMapper.FindUserWalletRecordsPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	userWalletRecordsList := userWalletRecprdsPage.Records.([]model.UserWalletRecords)
	if userWalletRecordsList != nil {
		userWalletRecordsVos := []vo2.UserWalletRecordsVo{}
		for _, topics := range userWalletRecordsList {
			userWalletRecordsVo := vo2.UserWalletRecordsVo{}
			utils.CopyProperties(&userWalletRecordsVo, topics)
			userWalletRecordsVos = append(userWalletRecordsVos, userWalletRecordsVo)
		}
		userWalletRecprdsPage.Records = userWalletRecordsVos
	}
	return userWalletRecprdsPage
}
