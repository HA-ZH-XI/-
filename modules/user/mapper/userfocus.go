package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	context2 "golang.org/x/net/context"
	"ksd-social-api/commons/page"
	"ksd-social-api/modules/user/model"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/model/vo"
)

/**
 * 用户关注
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserFocusMapper struct {
}

/**
 * 关注用户
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusMapper) SaveUserFocus(userFocus model.UserFocus) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		_, err = txOrm.Insert(&userFocus)
		_, err = txOrm.Raw("update  xk_user set gzs_num = gzs_num + 1,update_time = now()  where id = ? and gzs_num >= 0 and system_id = ?", userFocus.UserId, userFocus.SystemId).Exec()
		_, err = txOrm.Raw("update  xk_user set fans_num = fans_num + 1,update_time = now()  where id = ? and fans_num >= 0 and system_id = ?", userFocus.FocusId, userFocus.SystemId).Exec()
		return err
	})
	return nil == err
}

/**
 * 查询是否关注
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserFocusMapper) CountUserFocus(userId uint64, focusId uint64, systemId int) int64 {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int64
	err := mysql.Raw("SELECT is_focus FROM xk_focus_user WHERE user_id = ? AND focus_id = ? and system_id = ?", userId, focusId, systemId).QueryRow(&total)
	if nil != err {
		return -1
	}
	return total
}

/**
 * 取消关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusMapper) CancelUserFocus(userId uint64, focusId uint64, systemId int) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		_, err = txOrm.Raw("UPDATE xk_focus_user SET is_focus = 0,update_time = now() WHERE  user_id = ? AND focus_id = ? and system_id = ?", userId, focusId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set gzs_num = gzs_num - 1  where id = ? and gzs_num > 0 and system_id = ?", userId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set fans_num = fans_num - 1  where id = ? and fans_num > 0 and system_id = ?", focusId, systemId).Exec()
		return err
	})
	return nil == err
}

/**
 * 取消关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusMapper) CancelUserFanGzFocus(id uint64, userId uint64, focusId uint64, systemId int) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		_, err = txOrm.Raw("UPDATE xk_focus_user SET is_focus = 0,update_time = now() WHERE id= ? and user_id = ? AND focus_id = 1 and system_id = ?", id, userId, focusId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set gzs_num = gzs_num - 1  where id = ? and gzs_num > 0 and system_id = ?", userId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set fans_num = fans_num - 1  where id = ? and fans_num > 0 and system_id = ?", focusId, systemId).Exec()
		return err
	})
	return nil == err
}

/**
 * 移除粉丝
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusMapper) DelUserFanGzFocus(id uint64, userId uint64, focusId uint64, systemId int) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		_, err = txOrm.Raw("UPDATE xk_focus_user SET is_focus = 0,update_time = now() WHERE id= ? and user_id = ? AND focus_id = 1 and system_id = ?", id, userId, focusId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set gzs_num = gzs_num - 1  where id = ? and gzs_num > 0 and system_id = ?", userId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set fans_num = fans_num - 1  where id = ? and fans_num > 0 and system_id = ?", focusId, systemId).Exec()
		return err
	})
	return nil == err
}

/**
 * 恢复关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusMapper) RecoverUserFocus(userId uint64, focusId uint64, systemId int) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		_, err = txOrm.Raw("UPDATE xk_focus_user SET is_focus = 1,update_time = now() WHERE user_id = ? AND focus_id = ? and system_id = ?", userId, focusId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set gzs_num = gzs_num + 1,update_time = now()  where id = ? and gzs_num >= 0 and system_id = ?", userId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set fans_num = fans_num + 1,update_time = now()  where id = ? and fans_num >= 0 and system_id = ?", focusId, systemId).Exec()
		return err
	})
	return nil == err
}

/**
 * 删除关注
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 */
func (service *UserFocusMapper) DelUserFocus(id uint64, userId uint64, focusId uint64, systemId int) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		_, err = txOrm.Raw("delete from xk_focus_user where id = ? and user_id = ? AND focus_id = ? and system_id = ?", id, userId, focusId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set gzs_num = gzs_num - 1,update_time = now()  where id = ? and gzs_num > 0 and system_id = ?", userId, systemId).Exec()
		_, err = txOrm.Raw("update xk_user set fans_num = fans_num - 1,update_time = now()  where id = ? and fans_num > 0 and system_id = ?", focusId, systemId).Exec()
		return err
	})
	return nil == err
}

/**
* 粉丝列表
* @author feige
* @date 2023-11-03
* @version 1.0
* @desc
 */
func (mapper *UserFocusMapper) FindUserFocusFansPage(ctx context.UserFocusPageContext) (p *page.Page, err error) {
	// 1：准备一个文章数据容器
	userList := []vo.UserFansVo{}
	// 获取数据库链接
	mysql := orm.NewOrm()
	// 根据查询条件求总数 --------------------------------------------------------结束 Count() ---All()
	var total int64 = 0
	err1 := mysql.Raw(`
		select
		    count(1)
		from  xk_focus_user t1,xk_user t2 
		where t1.user_id = t2.id and t2.active =1 and t2.forbidden = 0 and t2.is_deleted=0 and t1.is_focus = 1 and t1.focus_id = ? and t1.system_id = ?
	`).SetArgs(ctx.UserId, ctx.SystemId).QueryRow(&total)
	if err1 != nil {
		return nil, err
	}
	// 开始根据总数和用户传递进来的分页信息，开始计算sql所需要的分页信息
	page := p.Page(ctx.PageNo, ctx.PageSize, total)
	_, err1 = mysql.Raw(`
		select
		    t1.id as opid,
		    t2.id,
			t2.user_name,
			t2.nick_name,
			t2.address,
			t2.sign,
			t2.uuid,
			t2.birth_day,
			t2.avatar,
			t2.male,
			t2.vip_flag,
			t1.create_time,
			t1.update_time,
			t2.author_flag,
			(select is_focus from xk_focus_user where focus_id = t1.user_id and user_id = ?) as is_focus
		from  xk_focus_user t1,xk_user t2 
		where t1.user_id = t2.id and t2.active =1 and t2.forbidden = 0 and t2.is_deleted=0 and t1.is_focus = 1 and t1.focus_id = ? and t1.system_id = ?
		order by t2.update_time desc
		limit ?,?
	`).SetArgs(ctx.UserId, ctx.UserId, ctx.SystemId, page.CurrentPage, page.PageSize).QueryRows(&userList)

	if err1 != nil {
		return nil, err
	}
	page.Records = userList
	return page, nil
}

/**
* 我的关注
* @author feige
* @date 2023-11-03
* @version 1.0
* @desc
 */
func (mapper *UserFocusMapper) FindUserFocusGzPage(ctx context.UserFocusPageContext) (p *page.Page, err error) {
	// 1：准备一个文章数据容器
	userList := []vo.UserFansVo{}
	// 获取数据库链接
	mysql := orm.NewOrm()
	// 根据查询条件求总数 --------------------------------------------------------结束 Count() ---All()
	var total int64 = 0
	err1 := mysql.Raw(`
		select
		    count(1)
		from  xk_focus_user t1,xk_user t2 
		where t1.focus_id = t2.id and t2.active =1 and t2.forbidden = 0 and t2.is_deleted=0 and t1.is_focus = 1 and t1.user_id = ? and t1.system_id = ?
	`).SetArgs(ctx.UserId, ctx.SystemId).QueryRow(&total)
	if err1 != nil {
		return nil, err
	}
	// 开始根据总数和用户传递进来的分页信息，开始计算sql所需要的分页信息
	page := p.Page(ctx.PageNo, ctx.PageSize, total)
	_, err1 = mysql.Raw(`
		select
		    t1.id opid,
		    t2.id,
			t2.user_name,
			t2.nick_name,
			t2.address,
			t2.sign,
			t2.uuid,
			t2.birth_day,
			t2.avatar,
			t2.male,
			t2.vip_flag,
			t1.create_time,
			t1.update_time,
			t2.author_flag,
			t1.is_focus
		from  xk_focus_user t1,xk_user t2 
		where t1.focus_id = t2.id and t2.active =1 and t2.forbidden = 0 and t2.is_deleted=0 and t1.is_focus = 1 and t1.user_id = ? and t1.system_id = ?
		order by t2.update_time desc
		limit ?,?
	`).SetArgs(ctx.UserId, ctx.SystemId, page.CurrentPage, page.PageSize).QueryRows(&userList)

	if err1 != nil {
		return nil, err
	}
	page.Records = userList
	return page, nil
}
