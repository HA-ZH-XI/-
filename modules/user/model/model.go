package model

import (
	"time"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @desc 用户实体
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type User struct {
	// 主键
	ID uint64 `orm:"column(id);pk;auto;description(主键)"`
	// 密码的盐
	Slat string `orm:"unique;column(slat);size(60);description(密码的盐);default:('')"`
	// 用户名
	UserName string `orm:"unique;column(user_name);size(100);description(用户名);default:('')"`
	// 用户账号（数字账号）
	Account string `orm:"unique;column(account);size(100);description(用户账号);default:('')"`
	// 昵称
	NickName string `orm:"column(nick_name);size(100);description(昵称);default:('')"`
	// 手机
	Telephone string `orm:"unique;column(telephone);size(20);description(手机);default:('')"`
	// 地址
	Address string `orm:"column(address);size(20);description(地址);default:('')"`
	// 密码
	Password string `orm:"column(password);size(60);description(密码);default:('')"`
	// 签名
	Sign string `orm:"column(sign);size(200);description(签名);default:('')"`
	// 背景
	BgImg string `orm:"column(bg_img);size(200);description(背景图);default:('')"`
	// 头像
	Avatar string `orm:"column(avatar);size(400);description(头像);default:('')"`
	// 生日
	BirthDay string `orm:"column(birth_day);description(生日);"`
	// 性别 0女 1男 2 保密
	Male int `orm:"column(male);size(1);description(0女 1男 2 保密);default(1);"`
	// 微信ID
	OpenId string `orm:"column(open_id);size(60);description(微信登录ID);default('');"`
	// 删除状态 0 未删除 1 删除
	IsDeleted int `orm:"column(is_deleted);size(1);description(0 未删除 1 删除);default(0);"`
	// 激活 0 未激活 1 激活
	Active int `orm:"column(active);size(1);description(0 未激活 1 激活);default(1);"`
	// 是否禁止 0 未禁止 1 禁止
	Forbidden int `orm:"column(forbidden);size(1);description(0 未禁止 1 禁止);default(0);"`
	// 创建时间
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime);description(创建时间);"`
	// 更新时间
	UpdateTime time.Time `orm:"column(update_time);auto_now;type(datetime);description(更新时间);"`
	AuthorFlag int       // 作者 0否 1是
	BbsNum     int       //文章数
	CourseNum  int       // 学习课程数
	FansNum    int       //粉丝数
	GzsNum     int       //关注数
	Uuid       string    // 用户唯一标识
	VipFlag    int       // vip身份标识 1 游客 2 vip  3 svip
	VipTime    time.Time // vip的过期时间
	Cron       float32   // 钱包余额
	Income     float32   // 提现收入
	Province   string    // 省份
	City       string    // 城市
	Job        string    // 职业
	Realname   string    // 真实姓名
	Idcard     string    // 身份证
	Idcardimgf string    // 身份证正面
	Idcardimgc string    // 身份证反面
	Alipaycode string    // 支付宝
	Weixincode string    // 微信号
	Bankcode   string    // 银行卡
	Bankimg    string    // 银行封面
	Bankaddr   string    // 银行卡开户行
	SystemId   int       // 系统ID
}

func (u *User) TableName() string {
	return "xk_user"
}

/**
 * @desc 用户权益
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserBenefits struct {
	ID          int       `orm:"column(id);pk;auto;description(主键)"` // 主键
	Title       string    //权益的名称
	Icon        string    //图片正面
	Sorted      int       //排序
	Status      int       //0 未发布 1 已发布
	Description string    //权益描述
	CreateTime  time.Time `orm:"auto_now_add;type(datetime);description(创建时间);"` // 创建时间
	UpdateTime  time.Time `orm:"auto_now;type(datetime);description(更新时间);"`     // 更新时间
	SystemId    int       // 系统ID
}

func (u *UserBenefits) TableName() string {
	return "xk_user_benefits"
}

/**
 * @desc 用户权益
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserVip struct {
	ID          int       `orm:"column(id);pk;auto;description(主键)"` // 主键
	Title       string    //	VIP类型名称,比如:VIP年卡
	Tag         string    //VIP类型标签,比如:超值,畅销
	Note        string    //VIP类型说明,比如:永久VIP特权
	Price       string    //价格
	Realprice   string    //折扣前价格
	VipDays     int       //会员卡实际天数: 比如:永久365*2,年365,月30,周7
	VipType     int       //会员卡类型: 7半年,6永久,5年,4季,3月,2周,1天,0无
	Status      int       //是否启用 启用后可购买, 禁用后不能购买,已购买可生效
	CreateTime  time.Time `orm:"auto_now_add;type(datetime);description(创建时间);"` // 创建时间
	UpdateTime  time.Time `orm:"auto_now;type(datetime);description(更新时间);"`     // 更新时间
	BenefitsIds string    //权益Ids--点亮权益(逗号隔开)
	Sorted      int       //排序字段
	SystemId    int       // 系统ID
}

func (u *UserVip) TableName() string {
	return "xk_user_vip"
}

/**
 * @desc 用户购买文章
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserBuyBbs struct {
	ID            uint64 `orm:"column(id);pk;auto;description(主键)"` // 主键
	UserId        uint64 //用户id
	BbsId         uint64 //文章ID
	Nickname      string //购买用户
	Avatar        string //头像
	Title         string //标题
	Description   string //描述
	Cover         string //封面
	Code          string //兑换码
	Price         string //购买价格
	Phone         string //用户手机
	Username      string //用户姓名
	Address       string //用户地址
	Orderno       string //订单编号
	Uuid          string //用户UUID
	OrderJson     string //订单完整信息
	Tradeno       string //订单交易号
	PayMethod     int    //1 微信 2 支付宝 3 兑换码
	PayMethodName string
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);description(创建时间);"` // 创建时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime);description(更新时间);"`     // 更新时间
	SystemId      int       // 系统ID
}

func (u *UserBuyBbs) TableName() string {
	return "xk_user_buy_bbs"
}

/**
 * @desc 用户购买文章
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserBuyVip struct {
	ID            uint64 `orm:"column(id);pk;auto;description(主键)"` // 主键
	UserId        uint64 //用户id
	VipId         int    //用户VIP
	Nickname      string //购买用户
	Avatar        string //头像
	Title         string //标题
	Description   string //描述
	Code          string //兑换码
	Price         string //购买价格
	Phone         string //用户手机
	Username      string //用户姓名
	Address       string //用户地址
	Orderno       string //订单编号
	Uuid          string //用户UUID
	OrderJson     string //订单完整信息
	Tradeno       string //订单交易号
	PayMethod     int    //1 微信 2 支付宝 3 兑换码
	PayMethodName string
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);description(创建时间);"` // 创建时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime);description(更新时间);"`     // 更新时间
	SystemId      int       // 系统ID
}

func (u *UserBuyVip) TableName() string {
	return "xk_user_buy_vip"
}

/**
 * @desc 关注和粉丝
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserFocus struct {
	ID            uint64    `orm:"column(id);pk;auto;description(主键)"` // 主键
	UserId        uint64    //用户id
	Uuid          string    //用户id
	FocusUuid     string    //用户id
	FocusId       uint64    //被关注人用户id
	IsFocus       int       //0:取消关注 1:关注  (取消关注不删除记录数据,用以记录主播奖励,单个用户反复关注不多给奖励)
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);description(创建时间);"` // 创建时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime);description(更新时间);"`     // 更新时间
	Nickname      string    // 用户昵称
	Avatar        string    // 用户头像
	FocusNickname string    // 被关注用户昵称
	FocusAvatar   string    // 被关注用户头像
	SystemId      int       // 系统ID
}

func (u *UserFocus) TableName() string {
	return "xk_focus_user"
}

/**
 * @desc 用户充值
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserWalletRecords struct {
	ID            uint64    `orm:"column(id);pk;auto;description(主键)"` // 主键
	UserId        uint64    //用户id
	Nickname      string    //购买用户
	Avatar        string    //头像
	Title         string    //标题
	Description   string    //描述
	Price         string    //购买价格
	Phone         string    //用户手机
	Username      string    //用户姓名
	Address       string    //用户地址
	Orderno       string    //订单编号
	Uuid          string    //用户UUID
	OrderJson     string    //订单完整信息
	Tradeno       string    //订单交易号
	PayMethod     int       //1 微信 2 支付宝 3 兑换码
	PayMethodName string    //支付方式的名字
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);description(创建时间);"` // 创建时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime);description(更新时间);"`     // 更新时间
	SystemId      int       // 系统ID
}

func (u *UserWalletRecords) TableName() string {
	return "xk_user_wallet_records"
}

/**
 * @desc 用户收入
 * @author feige
 * @date 2023-11-15
 * @version 1.0
 */
type UserWalletIncome struct {
	ID            uint64    `orm:"column(id);pk;auto;description(主键)"` // 主键
	UserId        uint64    //用户id
	Nickname      string    //购买用户
	Avatar        string    //头像
	Title         string    //标题
	Description   string    //描述
	Price         string    //购买价格
	Phone         string    //用户手机
	Username      string    //用户姓名
	Address       string    //用户地址
	Orderno       string    //订单编号
	Uuid          string    //用户UUID
	OrderJson     string    //订单完整信息
	Tradeno       string    //订单交易号
	PayMethod     int       //1 微信 2 支付宝 3 兑换码
	PayMethodName string    //支付方式的名字
	CreateTime    time.Time `orm:"auto_now_add;type(datetime);description(创建时间);"` // 创建时间
	UpdateTime    time.Time `orm:"auto_now;type(datetime);description(更新时间);"`     // 更新时间
	SystemId      int       // 系统ID
}

func (u *UserWalletIncome) TableName() string {
	return "xk_user_wallet_income"
}

/**
 * @author feige
 * @date 2023-10-08
 * @version 1.0
 * @desc 兑换码
 */
type UserVipCode struct {
	ID       int    `orm:"column(id);pk;auto;description(主键)"` // 业务主键
	VipId    int    //vip身份
	Code     string //兑换码
	Mark     int    //是否兑换
	SystemId int    // 系统ID
}

func (u *UserVipCode) TableName() string {
	return "xk_user_vip_code"
}

/**
 * @author feige
 * @date 2023-10-08
 * @version 1.0
 * @desc 兑换码
 */
type UserWalletCode struct {
	ID       int    `orm:"column(id);pk;auto;description(主键)"` // 业务主键
	Cron     int    //兑换金额
	Code     string //兑换码
	Mark     int    //是否兑换
	SystemId int    // 系统ID
}

func (u *UserWalletCode) TableName() string {
	return "xk_user_wallet_code"
}
