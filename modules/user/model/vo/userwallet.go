package vo

import "time"

type UserWalletRecordsVo struct {
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
