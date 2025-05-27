package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	"ksd-social-api/modules/user/model"
)

/**
 * 用户
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserVipMapper struct {
}

/**
 * 查看用户vip列表
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (mapper *UserVipMapper) FindUserVipList(systemId int) []*model.UserVip {
	db := orm.NewOrm()
	var models []*model.UserVip
	_, err := db.Raw("select id,title,tag,note,price,realprice,vip_days,benefits_ids from xk_user_vip  where status = 1 and system_id = ? order by sorted asc").SetArgs(systemId).QueryRows(&models)
	if nil != err {
		return nil
	}
	return models
}

/**
 * 根据ID获取用户VIP信息
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (mapper *UserVipMapper) GetUserVipDetail(id int, systemId int) *model.UserVip {
	db := orm.NewOrm()
	var userVip model.UserVip
	err := db.Raw("select id,title,tag,note,price,realprice,vip_type,vip_days,benefits_ids from xk_user_vip  where id = ? and system_id = ?", id, systemId).QueryRow(&userVip)
	if nil != err {
		return nil
	}
	return &userVip
}

/**
 * 根据code查询用户身份信息
 * @author feige
 * @date 2023-12-26
 * @version 1.0
 * @desc
 */
func (mapper *UserVipMapper) GetUserVipDetailByCode(code string, systemId int) *model.UserVip {
	db := orm.NewOrm()
	var userVip model.UserVip
	err := db.Raw("select t1.id,t1.title,t1.tag,t1.note,t1.price,t1.realprice,t1.vip_type,t1.vip_days,t1.benefits_ids from xk_user_vip t1,xk_user_vip_code t2  where t1.id = t2.vip_id and and t2.system_id = ? and  t2.code = ?", systemId, code).QueryRow(&userVip)
	if nil != err {
		return nil
	}
	return &userVip
}

/**
 * 查看用户vip列表
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (mapper *UserVipMapper) FindUserBenefits(ids string, systemId int) []*model.UserBenefits {
	db := orm.NewOrm()
	var models []*model.UserBenefits
	_, err := db.Raw("SELECT id,title,icon,description FROM xk_user_benefits WHERE status = 1 and system_id = ? AND FIND_IN_SET(id,?)", systemId, ids).QueryRows(&models)
	if nil != err {
		return nil
	}
	return models
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc vip用户过期
 */
func (mapper *UserVipMapper) UpdateVIPPeriod(userId uint64, systemId int) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET vip_time = NULL , vip_flag = 1 , update_time = now()   WHERE id = ? and vip_time <= now() and system_id = ?").SetArgs(userId, systemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}
