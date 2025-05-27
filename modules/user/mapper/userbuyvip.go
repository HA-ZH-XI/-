package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	context2 "golang.org/x/net/context"
	"ksd-social-api/commons/page"
	"ksd-social-api/modules/user/model"
	"time"
)

/**
 * 用户购买VIP
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserBuyVipMapper struct {
}

var userVipMapper = UserVipMapper{}
var userMapper = UserMapper{}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  保存用户购买VIP信息
 */
func (mapper *UserBuyVipMapper) SaveUserBuyVip(userBuyVip model.UserBuyVip) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		// 保存用户课程订单
		_, err = txOrm.Insert(&userBuyVip)
		// 用户升级VIP
		userVip := userVipMapper.GetUserVipDetail(userBuyVip.VipId, userBuyVip.SystemId)
		if nil != userVip {
			// 会员卡类型: 1天卡 2 周卡 3月卡 4 季度 5 年卡 6 SVIP
			// 开始查询用户的充值信息
			user := userMapper.GetUserByID(userBuyVip.UserId, userBuyVip.SystemId)
			if nil != user {
				// 获取用户会员时间
				vipTime := user.VipTime
				if vipTime.IsZero() {
					// 如果是nil，就获取当前时间
					vipTime = time.Now()
				}

				// 如果是初次或者每次身份叠加就修改身份的类型
				if user.VipFlag == 0 || userVip.VipType > user.VipFlag {
					user.VipFlag = userVip.VipType
				}

				// 累加天数的VIP时间
				date := vipTime.AddDate(0, 0, userVip.VipDays)
				_, err = orm.NewOrm().Raw("UPDATE xk_user SET vip_flag = ?,vip_time = ? WHERE id = ? and system_id = ?").
					SetArgs(user.VipFlag, date, userBuyVip.UserId, userBuyVip.SystemId).
					Exec()
			}
		}

		return err
	})

	return err == nil
}

/**
 * 统计订单是否支付--购买VIP
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (mapper *UserBuyVipMapper) CountUserBuyVipByOutTradeNo(OutTradeNo string, systemId int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_vip WHERE orderno = ? and system_id = ?", OutTradeNo, systemId).QueryRow(&total)
	// 返回
	if nil != err {
		return 0
	}
	return total
}

/**
 *
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (mapper *UserBuyVipMapper) CountUserBuyVipByCode(code string, systemId int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_vip WHERE code = ? and mark = 0 and system_id = ?", code, systemId).QueryRow(&total)
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
 * @desc 我学习的VIP
 */
func (mapper *UserBuyVipMapper) FindMyUserBuyVipOrderPage(userId uint64, systemId int, pageNo int64, pageSize int64) (p *page.Page, err error) {
	// 准备容器对象，开始装载数据库数据
	userBuyVipList := []model.UserBuyVip{}
	//创建orm对象
	mysql := orm.NewOrm()
	cond := orm.NewCondition()
	// 开始执行sql查询
	qs := mysql.QueryTable("xk_user_buy_vip")
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
	_, err2 := qs.SetCond(cond).OrderBy("-create_time").Limit(page.PageSize, page.CurrentPage).All(&userBuyVipList,
		"id",
		"user_id",
		"vip_id",
		"create_time",
		"update_time",
		"nickname",
		"avatar",
		"title",
		"description",
		"code",
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
	)
	// 6：把查询的数据放入到分页的records字段，准备返回
	page.Records = userBuyVipList
	// 7: 如果没有找到直接返回
	if err2 != nil {
		return nil, err2
	}
	return page, nil
}
