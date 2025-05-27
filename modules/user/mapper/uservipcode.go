package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	"ksd-social-api/modules/user/model"
)

type UserVipCodeMapper struct{}

/**
 * 课程兑换码
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserVipCodeMapper) SaveUserVipCode(userVipCode []model.UserVipCode) bool {
	// 创建orm对象
	_, err := orm.NewOrm().InsertMulti(100, userVipCode)
	return nil == err
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 查询code是否存在
 */
func (mapper *UserVipCodeMapper) CountUserVipCode(code string, systemId int) int {
	db := orm.NewOrm()
	var total int
	err := db.Raw("select count(1) from xk_user_vip_code where  code = ? and mark = 0 and system_id = ?", code, systemId).QueryRow(&total)
	if nil != err {
		return 0
	}
	return total
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 根据code和课程id删除兑换码
 */
func (mapper *UserVipCodeMapper) DelUserVipCode(code string, systemId int) bool {
	db := orm.NewOrm()
	_, err := db.Raw("delete from xk_user_vip_code where code = ? and mark = 0 and system_id = ?", code, systemId).Exec()
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
func (mapper *UserVipCodeMapper) UpdateUserVipCode(code string, systemId int) bool {
	db := orm.NewOrm()
	_, err := db.Raw("update xk_user_vip_code set mark = 1 where  code = ? and system_id = ?", code, systemId).Exec()
	if nil != err {
		return false
	}
	return true
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据code获取UserVip
 */
func (mapper *UserVipCodeMapper) GetUserVipByCode(code string, systemId int) *model.UserVipCode {
	// 创建orm对象
	mysql := orm.NewOrm()
	var userVipCode model.UserVipCode
	err := mysql.Raw("select * from xk_user_vip_code where code = ? and system_id = ?").SetArgs(code, systemId).QueryRow(&userVipCode)
	if err != nil {
		return nil
	}
	return &userVipCode
}
