package vo

import "time"

type UserVipVo struct {
	ID          int               `json:"id"`
	Title       string            `json:"title"`       //	VIP类型名称,比如:VIP年卡
	Tag         string            `json:"tag"`         //VIP类型标签,比如:超值,畅销
	Note        string            `json:"note"`        //VIP类型说明,比如:永久VIP特权
	Price       string            `json:"price"`       //价格
	Realprice   string            `json:"realprice"`   //折扣前价格
	VipDays     int               `json:"vipDays"`     //会员卡实际天数: 比如:永久365*2,年365,月30,周7
	VipType     int               `json:"vipType"`     //会员卡类型: 7半年,6永久,5年,4季,3月,2周,1天,0无
	BenefitsIds []*UserBenefitsVo `json:"benefitsIds"` //权益Ids--点亮权益(逗号隔开)
}

type UserBenefitsVo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`       //权益的名称
	Icon        string `json:"icon"`        //图片正面
	Description string `json:"description"` //权益描述
}

/**
 * @desc 用户购买文章
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserBuyVipVo struct {
	ID            uint64    `json:"id"`            // 主键
	UserId        uint64    `json:"userId"`        //用户id
	Nickname      string    `json:"nickname"`      //购买用户
	Avatar        string    `json:"avatar"`        //头像
	Title         string    `json:"title"`         //标题
	Description   string    `json:"description"`   //描述
	Price         string    `json:"price"`         //购买价格
	Phone         string    `json:"phone"`         //用户手机
	Username      string    `json:"username"`      //用户姓名
	Address       string    `json:"address"`       //用户地址
	Orderno       string    `json:"orderno"`       //订单编号
	Uuid          string    `json:"uuid"`          //用户UUID
	OrderJson     string    `json:"orderJson"`     //订单完整信息
	Tradeno       string    `json:"tradeno"`       //订单交易号
	PayMethod     int       `json:"payMethod"`     //1 微信 2 支付宝 3 兑换码
	PayMethodName string    `json:"payMethodName"` //支付方式的名字
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}
