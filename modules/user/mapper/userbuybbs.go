package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	context2 "golang.org/x/net/context"
	"ksd-social-api/commons/page"
	bmodel "ksd-social-api/modules/social/model"
	"ksd-social-api/modules/user/model"
)

/**
 * 用户购买文章
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserBuyBbsMapper struct {
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  保存用户购买文章信息
 */
func (service *UserBuyBbsMapper) SaveUserBuyBbs(userUserBuyBbs model.UserBuyBbs) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		// 保存用户文章订单
		_, err = txOrm.Insert(&userUserBuyBbs)
		// 同步购买和订阅的用户数量
		_, err = txOrm.Raw("update xk_bbs_topics set buy_num = buy_num + 1  where id = ? and buy_num >= 0 and system_id = ?", userUserBuyBbs.BbsId, userUserBuyBbs.SystemId).Exec()
		return err
	})
	return err == nil
}

/**
 * 查询用户是否购买文章
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyBbsMapper) CountUserBuyBbs(userId uint64, bbsId uint64, systemId int, payMethod int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_bbs WHERE userid  = ? AND bbs_id = ? and pay_method = ? and system_id = ?", userId, bbsId, payMethod, systemId).QueryRow(&total)
	// 返回
	if nil != err {
		return 0
	}
	return total
}

/**
 * 查询用户是否购买文章
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyBbsMapper) CountUserBuyBbsNo(userId uint64, bbsId uint64, systemId int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_bbs WHERE userid  = ? AND bbs_id = ? and system_id = ?", userId, bbsId, systemId).QueryRow(&total)
	// 返回
	if nil != err {
		return 0
	}
	return total
}

/**
 * 统计文章购买人数
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyBbsMapper) CountUserBuyBbsId(bbsId uint64, systemId int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_bbs WHERE bbs_id = ? and system_id = ?", bbsId, systemId).QueryRow(&total)
	// 返回
	if nil != err {
		return 0
	}
	return total
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 我学习的文章-订单
 */
func (mapper *UserBuyBbsMapper) FindMyUserBuyBbsPage(userId uint64, systemId int, pageNo int64, pageSize int64) (p *page.Page, err error) {
	// 准备容器对象，开始装载数据库数据
	topicsList := []bmodel.BbsTopics{}
	//创建orm对象
	mysql := orm.NewOrm()
	var total int64 = 0
	mysql.Raw(`
		select 
			count(1)
		from
			xk_bbs_topics t1,xk_user_buy_bbs t2 
		where 
		    t1.id = t2.bbs_id and t1.status = 1 and t1.is_deleted = 0 and t2.user_id = ? and t2.system_id = ?
	`).SetArgs(userId, systemId).QueryRow(&total)
	// 重新换算分页和规则
	page := p.Page(pageNo, pageSize, total)
	// 开始执行sql查询
	_, err2 := mysql.Raw(`
		select 
			t1.id,
			t1.title,
			t1.tags,
			t1.cover,
			t1.user_id,
			t1.category_id,
			t1.viewcount,
			t1.create_time,
			t1.update_time,
			t1.static_link,
			t1.top_flag,
			t1.category_name,
			t1.vip_flag,
			t1.avatar,
			t1.nickname,
			t1.uuid,
			t1.comment_flag,
			t1.status,
			t1.comment_num,
			t1.fav_num,
			t1.like_num,
			t1.category_pname,
			t1.category_pid,
			t1.description
		from
			xk_bbs_topics t1,xk_user_buy_bbs t2 
		where 
		    t1.id = t2.bbs_id and t1.status = 1 and t1.is_deleted = 0 and t2.user_id = ? and t2.system_id = ?
		order by t2.create_time desc
		limit ?,?
	`).SetArgs(userId, systemId, page.CurrentPage, page.PageSize).QueryRows(&topicsList)
	// 6：把查询的数据放入到分页的records字段，准备返回
	page.Records = topicsList
	// 7: 如果没有找到直接返回
	if err2 != nil {
		return nil, err2
	}
	return page, nil
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 我学习的文章
 */
func (mapper *UserBuyBbsMapper) FindMyUserBuyBbsOrderPage(userId uint64, systemId int, pageNo int64, pageSize int64) (p *page.Page, err error) {
	// 准备容器对象，开始装载数据库数据
	userBuyBbsList := []model.UserBuyBbs{}
	//创建orm对象
	mysql := orm.NewOrm()
	var total int64 = 0
	mysql.Raw(`
		select 
			count(1)
		from
			xk_bbs_topics t1,xk_user_buy_bbs t2 
		where 
		    t1.id = t2.bbs_id and t1.status = 1 and t1.is_deleted = 0 and t2.user_id = ? and t2.system_id = ?
	`).SetArgs(userId, systemId).QueryRow(&total)
	// 重新换算分页和规则
	page := p.Page(pageNo, pageSize, total)
	// 开始执行sql查询
	_, err2 := mysql.Raw(`
		select 
			t2.id,
			t2.user_id,
			t2.bbs_id,
			t2.create_time,
			t2.update_time,
			t2.nickname,
			t2.avatar,
			t2.title,
			t2.description,
			t2.cover,
			t2.code,
			t2.price,
			t2.phone,
			t2.username,
			t2.address,
			t2.orderno,
			t2.uuid,
			t2.order_json,
			t2.tradeno,
			t2.pay_method,
			t2.pay_method_name
		from
			xk_bbs_topics t1,xk_user_buy_bbs t2 
		where 
		    t1.id = t2.bbs_id and t1.status = 1 and t1.is_deleted = 0 and t2.user_id = ? and t2.system_id = ?
		order by t2.create_time desc
		limit ?,?
	`).SetArgs(userId, systemId, page.CurrentPage, page.PageSize).QueryRows(&userBuyBbsList)
	// 6：把查询的数据放入到分页的records字段，准备返回
	page.Records = userBuyBbsList
	// 7: 如果没有找到直接返回
	if err2 != nil {
		return nil, err2
	}
	return page, nil
}
