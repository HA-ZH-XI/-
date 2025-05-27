package vo

import (
	"time"
)

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserVo struct {
	// 主键
	ID uint64 `json:"id"`
	// 用户uuid
	Uuid string `json:"uuid"`
	// 用户名
	UserName string `json:"username"`
	// 密码
	Password string `json:"password"`
	// 账号
	Account string `json:"account"`
	// 昵称
	NickName string `json:"nickname"`
	// 手机
	Telephone string `json:"telephone"`
	// 地址
	Address string `json:"address"`
	// 签名
	Sign string `json:"sign"`
	// 背景
	BgImg string `json:"bgimg"`
	// 头像
	Avatar string `json:"avatar"`
	// 生日
	BirthDay string `json:"birthDay"`
	// 性别 0女 1男 2 保密
	Male int `json:"male"`
	// 1 普通 2天 3周 4 月  5季度  6年 7 SVIP
	VipFlag int `json:"vipFlag"`
	// VIP过期时间
	VipTime time.Time `json:"vipTime"`
	// VIP过期时间
	VipDays int `json:"vipDays"`
	// 注册天数
	RegDays int `json:"RegDays"`
	// 收入
	Income float32 `json:"income"`
	// 学习币
	Cron float32 `json:"cron"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 星座
	Constellation string `json:"constellation"`
	// 属相
	Twelve string `json:"twelve"`
	// 省份
	Province string `json:"province"`
	// 城市
	City string `json:"city"`
	// 职业
	Job        string `json:"job"`
	Realname   string `json:"realname"`   // 真实姓名
	Idcard     string `json:"idcard"`     // 身份证
	Idcardimgf string `json:"idcardimgf"` // 身份证正面
	Idcardimgc string `json:"idcardimgc"` // 身份证反面
	Alipaycode string `json:"alipaycode"` // 支付宝
	Weixincode string `json:"weixincode"` // 微信号
	Bankcode   string `json:"bankcode"`   // 银行卡
	Bankimg    string `json:"bankimg"`    // 银行封面
	Bankaddr   string `json:"bankaddr"`   // 银行卡开户行
	UserId     uint64 `json:"userId"`     // 用户ID
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserBankVo struct {
	// 主键
	ID uint64 `json:"id"`
	// 用户uuid
	Uuid string `json:"uuid"`
	// 用户名
	UserName string `json:"username"`
	// 昵称
	NickName string `json:"nickname"`
	// 手机
	Telephone string `json:"telephone"`
	// 职业
	Job        string `json:"job"`
	Realname   string `json:"realname"`   // 真实姓名
	Idcard     string `json:"idcard"`     // 身份证
	Idcardimgf string `json:"idcardimgf"` // 身份证正面
	Idcardimgc string `json:"idcardimgc"` // 身份证反面
	Alipaycode string `json:"alipaycode"` // 支付宝
	Weixincode string `json:"weixincode"` // 微信号
	Bankcode   string `json:"bankcode"`   // 银行卡
	Bankimg    string `json:"bankimg"`    // 银行封面
	Bankaddr   string `json:"bankaddr"`   // 银行卡开户行
	UserId     uint64 `json:"userId"`     // 用户ID
}

/**
 * @desc 用户购买课程
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserBuyCourseVo struct {
	ID            uint64    `json:"id"`            // 主键
	UserId        uint64    `json:"userId"`        //用户id
	CourseId      uint64    `json:"courseId"`      //用户学习的课程id
	Nickname      string    `json:"nickname"`      //购买用户
	Avatar        string    `json:"avatar"`        //头像
	Coursetitle   string    `json:"coursetitle"`   //购买课程
	Coursecover   string    `json:"coursecover"`   //课程封面
	Description   string    `json:"description"`   //描述
	Code          string    `json:"code"`          //兑换码
	Price         string    `json:"price"`         //购买价格
	Phone         string    `json:"phone"`         //用户手机
	Username      string    `json:"username"`      //用户姓名
	Address       string    `json:"address"`       //用户地址
	Orderno       string    `json:"orderno"`       //订单编号
	Uuid          string    `json:"uuid"`          //用户UUID
	OrderJson     string    `json:"orderJson"`     //订单完整信息
	Tradeno       string    `json:"tradeno"`       //订单交易号
	PayMethod     int       `json:"payMethod"`     //1 微信 2 支付宝 3 兑换码
	PayMethodName string    `json:"payMethodName"` // 支付方式
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}

/**
 * @desc 用户购买文章
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserBuyBbsVo struct {
	ID            uint64    `json:"id"`            // 主键
	UserId        uint64    `json:"userId"`        //用户id
	BbsId         uint64    `json:"bbsId"`         //文章ID
	Nickname      string    `json:"nickname"`      //购买用户
	Avatar        string    `json:"avatar"`        //头像
	Title         string    `json:"title"`         //标题
	Description   string    `json:"description"`   //描述
	Cover         string    `json:"cover"`         //封面
	Code          string    `json:"code"`          //兑换码
	Price         string    `json:"price"`         //购买价格
	Phone         string    `json:"phone"`         //用户手机
	Username      string    `json:"username"`      //用户姓名
	Address       string    `json:"address"`       //用户地址
	Orderno       string    `json:"orderno"`       //订单编号
	Uuid          string    `json:"uuid"`          //用户UUID
	OrderJson     string    `json:"orderJson"`     //订单完整信息
	Tradeno       string    `json:"tradeno"`       //订单交易号
	PayMethod     int       `json:"payMethod"`     //1 微信 2 支付宝 3 兑换码
	PayMethodName string    `json:"payMethodName"` // 支付方式
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserFansVo struct {
	// 主键
	Id uint64 `json:"id"`
	// 用户uuid
	Uuid string `json:"uuid"`
	// 用户名
	UserName string `json:"username"`
	// 昵称
	NickName string `json:"nickname"`
	// 地址
	Address string `json:"address"`
	// 签名
	Sign string `json:"sign"`
	// 头像
	Avatar string `json:"avatar"`
	// 生日
	BirthDay string `json:"birthDay"`
	// 性别 0女 1男 2 保密
	Male int `json:"male"`
	// 1 普通 2天 3周 4 月  5季度  6年 7 SVIP
	VipFlag int `json:"vipFlag"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 是否关注
	IsFocus int `json:"isFocus"`
	// 主键
	Opid uint64 `json:"opid"`
}
