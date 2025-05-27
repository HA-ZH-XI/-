package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	"ksd-social-api/modules/user/model"
)

type UserWalletCodeMapper struct{}

/**
 * 课程兑换码
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserWalletCodeMapper) SaveUserWalletCode(userWalletCode []model.UserWalletCode) bool {
	// 创建orm对象
	_, err := orm.NewOrm().InsertMulti(100, userWalletCode)
	return nil == err
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 查询用户是否兑换此课程
 */
func (mapper *UserWalletCodeMapper) CountUserWalletCode(code string, systemId int) int {
	db := orm.NewOrm()
	var total int
	err := db.Raw("select count(1) from xk_user_wallet_code where  code = ? and mark = 0 and system_id = ?", code, systemId).QueryRow(&total)
	if nil != err {
		return 0
	}
	return total
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 根据code删除
 */
func (mapper *UserWalletCodeMapper) DelUserWalletCode(code string, systemId int) bool {
	db := orm.NewOrm()
	_, err := db.Raw("delete from xk_user_wallet_code where code = ? and mark = 0 and system_id = ?", code, systemId).Exec()
	if nil != err {
		return false
	}
	return true
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 标识兑换码已被使用
 */
func (mapper *UserWalletCodeMapper) UpdateUserWalletCode(code string, systemId int) bool {
	db := orm.NewOrm()
	_, err := db.Raw("update xk_user_wallet_code set mark = 1 where code = ? and system_id = ?", code, systemId).Exec()
	if nil != err {
		return false
	}
	return true
}

/**
 * 根据ID获取用户VIP信息
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (mapper *UserWalletCodeMapper) GetUserWalletDetail(code string, systemId int) *model.UserWalletCode {
	db := orm.NewOrm()
	var userWalletCode model.UserWalletCode
	err := db.Raw("select id,cron,code,mark from xk_user_wallet_code  where code = ? and system_id = ?", code, systemId).QueryRow(&userWalletCode)
	if nil != err {
		return nil
	}
	return &userWalletCode
}
