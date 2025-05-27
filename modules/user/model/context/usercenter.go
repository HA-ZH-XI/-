package context

import "github.com/gookit/validate"

/**
 * @desc 用户中心
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserCenterContext struct {
	PageNo    int64  `json:"pageNo"`    // 第几页
	PageSize  int64  `json:"pageSize"`  // 每页显示多少条
	Status    int64  `json:"status"`    // 1 发布中  0 草稿
	IsDeleted int64  `json:"isDeleted"` // 0 未删除  1 已删除
	UserId    uint64 `json:"userId"`    // 用户ID
	SystemId  int    `json:"systemId"`
}

/**
 * @desc 用户中心 - 消息
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type MessageContext struct {
	PageNo   int64  `json:"pageNo"`   // 第几页
	PageSize int64  `json:"pageSize"` // 每页显示多少条
	Mtype    int    `json:"mtype"`    // 消息类型
	UserId   uint64 `json:"userId"`   // 用户ID
	SystemId int    `json:"systemId"`
}

/**
 * @desc 用户中心 - 系统消息
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type MessagePageContext struct {
	PageNo   int64 `json:"pageNo"`   // 第几页
	SType    int   `json:"stype"`    // 消息类型  1平台消息 2课程
	PageSize int64 `json:"pageSize"` // 每页显示多少条
	SystemId int   `json:"systemId"`
}

/**
 * @desc 用户密码修改
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserPasswordContext struct {
	Password    string `validate:"required" json:"password"`    // 密码
	Newpassword string `validate:"required" json:"newpassword"` // 确认密码
	UserId      uint64 `json:"userId"`                          // 用户ID
	SystemId    int    `json:"systemId"`
}

// Messages 您可以自定义验证器错误消息
func (f UserPasswordContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
	}
}

// Translates 你可以自定义字段翻译
func (f UserPasswordContext) Translates() map[string]string {
	return validate.MS{
		"Password":    "密码",
		"Newpassword": "确认密码",
	}
}

/**
 * @desc 用户密码修改
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserBankContext struct {
	Realname   string `validate:"required" json:"realname"`   // 真实姓名
	Telephone  string `validate:"required" json:"telephone"`  // 手机
	Idcard     string `validate:"required" json:"idcard"`     // 身份证
	Idcardimgf string `validate:"required" json:"idcardimgf"` // 身份证正面
	Idcardimgc string `validate:"required" json:"idcardimgc"` // 身份证反面
	Alipaycode string `validate:"required" json:"alipaycode"` // 支付宝
	Weixincode string `validate:"required" json:"weixincode"` // 微信号
	Bankcode   string `validate:"required" json:"bankcode"`   // 银行卡
	Bankimg    string `validate:"required" json:"bankimg"`    // 银行封面
	Bankaddr   string `validate:"required" json:"bankaddr"`   // 银行卡开户行
	UserId     uint64 `json:"userId"`                         // 用户ID
	SystemId   int    `json:"systemId"`
}

// Messages 您可以自定义验证器错误消息
func (f UserBankContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
	}
}

// Translates 你可以自定义字段翻译
func (f UserBankContext) Translates() map[string]string {
	return validate.MS{
		"Realname":   "昵称",
		"Telephone":  "手机",
		"Idcard":     "身份证",
		"IdcardImgf": "身份证正面",
		"IdcardImgc": "身份证反面",
		"Alipaycode": "支付宝",
		"Weixincode": "微信号",
		"BankImg":    "银行卡",
		"Bankcode":   "银行封面",
		"BankAddr":   "银行卡开户行",
	}
}
