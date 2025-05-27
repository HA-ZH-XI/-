package service

import (
	"ksd-social-api/commons/page"
	"ksd-social-api/modules/user/model"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/model/vo"
	"ksd-social-api/utils"
)

/**
 * 我的钱包
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) GetUserWalletByUserID(userId uint64, systemId int) *model.User {
	return userMapper.GetUserByID(userId, systemId)
}

/**
 * 我的钱包 -- 充值列表
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindUserWalletRecordsPage(ctx context.UserCenterContext) (p *page.Page) {
	userWalletRecprdsPage, _ := userWalletMapper.FindUserWalletRecordsPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	userWalletRecordsList := userWalletRecprdsPage.Records.([]model.UserWalletRecords)
	if userWalletRecordsList != nil {
		userWalletRecordsVos := []vo.UserWalletRecordsVo{}
		for _, topics := range userWalletRecordsList {
			userWalletRecordsVo := vo.UserWalletRecordsVo{}
			utils.CopyProperties(&userWalletRecordsVo, topics)
			userWalletRecordsVos = append(userWalletRecordsVos, userWalletRecordsVo)
		}
		userWalletRecprdsPage.Records = userWalletRecordsVos
	}
	return userWalletRecprdsPage
}

/**
 * 我的钱包 -- 收入列表
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindUserWalletIncomePage(ctx context.UserCenterContext) (p *page.Page, err error) {
	return userWalletMapper.FindUserWalletIncomePage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
}
