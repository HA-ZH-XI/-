package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	context2 "golang.org/x/net/context"
	"ksd-social-api/commons/page"
	"ksd-social-api/modules/user/model"
	"strconv"
)

type UserWalletMapper struct {
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  保存用户钱包
 */
func (mapper *UserWalletMapper) SaveUserWalletRecords(userWalletRecords model.UserWalletRecords) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		// 保存用户课程订单
		_, err = txOrm.Insert(&userWalletRecords)
		// 同步购买和订阅的用户数量
		price, _ := strconv.ParseFloat(userWalletRecords.Price, 64)
		_, err = txOrm.Raw("update xk_user set cron = cron + ?  where id = ? and cron >= 0 and system_id = ?", price, userWalletRecords.UserId, userWalletRecords.SystemId).Exec()
		return err
	})
	return err == nil
}

/**
 * 统计订单是否支付--充值
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserWalletMapper) CountUserWalletRecordsByOutTradeNo(OutTradeNo string, systemId int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_wallet_records WHERE  orderno = ? and system_id = ?", OutTradeNo, systemId).QueryRow(&total)
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
 * @desc 查看充值
 */
func (mapper *UserWalletMapper) FindUserWalletRecordsPage(userId uint64, systemId int, pageNo int64, pageSize int64) (p *page.Page, err error) {
	// 准备容器对象，开始装载数据库数据
	walletRecords := []model.UserWalletRecords{}
	//创建orm对象
	mysql := orm.NewOrm()
	cond := orm.NewCondition()
	// 开始执行sql查询
	qs := mysql.QueryTable("xk_user_wallet_records")
	// 设定两个条件
	cond = cond.And("user_id", userId).And("system_id", systemId)
	// 执行count查询
	total, err1 := qs.SetCond(cond).Count()
	if err1 != nil {
		return nil, err
	}
	// 重新换算分页和规则
	page := p.Page(pageNo, pageSize, total)
	// 分页查询
	_, err2 := qs.SetCond(cond).OrderBy("-create_time").Limit(page.PageSize, page.CurrentPage).All(&walletRecords,
		"id",
		"user_id",
		"create_time",
		"update_time",
		"nickname",
		"avatar",
		"title",
		"price",
		"phone",
		"username",
		"address",
		"orderno",
		"uuid",
		"order_json",
		"tradeno",
		"pay_method",
		"pay_method_name",
		"description",
	)
	// 6：把查询的数据放入到分页的records字段，准备返回
	page.Records = walletRecords
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
 * @desc 查看收入
 */
func (mapper *UserWalletMapper) FindUserWalletIncomePage(userId uint64, systemId int, pageNo int64, pageSize int64) (p *page.Page, err error) {
	// 准备容器对象，开始装载数据库数据
	walletIncomes := []model.UserWalletIncome{}
	//创建orm对象
	mysql := orm.NewOrm()
	cond := orm.NewCondition()
	// 开始执行sql查询
	qs := mysql.QueryTable("xk_user_wallet_income")
	// 设定两个条件
	cond = cond.And("user_id", userId).And("system_id", systemId)
	// 执行count查询
	total, err1 := qs.SetCond(cond).Count()
	if err1 != nil {
		return nil, err
	}
	// 重新换算分页和规则
	page := p.Page(pageNo, pageSize, total)
	// 分页查询
	_, err2 := qs.SetCond(cond).OrderBy("-create_time").Limit(page.PageSize, page.CurrentPage).All(&walletIncomes,
		"id",
		"user_id",
		"create_time",
		"update_time",
		"nickname",
		"avatar",
		"title",
		"cover",
		"price",
		"phone",
		"username",
		"address",
		"orderno",
		"uuid",
		"order_json",
		"tradeno",
		"pay_method",
		"pay_method_name",
		"description",
	)
	// 6：把查询的数据放入到分页的records字段，准备返回
	page.Records = walletIncomes
	// 7: 如果没有找到直接返回
	if err2 != nil {
		return nil, err2
	}
	return page, nil
}
