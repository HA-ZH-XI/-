package web

import "ksd-social-api/modules/user/model/context"

/**
 * @desc 查询所有的充值记录
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) FindUserWalletRecords() {
	pageContext := context.UserCenterContext{}
	web.BindJSON(&pageContext)
	pageContext.UserId = web.GetUserId()
	pageContext.SystemId = web.GetSystemId()
	userWalletPages := userCenterService.FindUserWalletRecordsPage(pageContext)
	web.Ok(userWalletPages)
}
