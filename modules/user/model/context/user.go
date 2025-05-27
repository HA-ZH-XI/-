package context

import (
	"github.com/gookit/validate"
	"ksd-social-api/commons/lib"
)

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc  管理员修改用户数据的载体
 */
type UserPwdContext struct {
	SystemId        int    `json:"systemId"`
	Uuid            string `json:"uuid"`                                                    //用户UUID
	UserId          uint   `validate:"required|gt:0" json:"userId"`                         // 修改那个用户的id
	Password        string `validate:"required|minLen:6|maxLen:16" json:"password" `        // 密码
	ConfirmPassword string `validate:"required|minLen:6|maxLen:16" json:"confirmPassword" ` // 确认密码
}

// Messages 您可以自定义验证器错误消息
func (f UserPwdContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
		"gt":       "{field}必须大于0",
		"minLen":   "{field}大于等于6位",
		"maxLen":   "{field}小于等于16位",
	}
}

// Translates 你可以自定义字段翻译
func (f UserPwdContext) Translates() map[string]string {
	return validate.MS{
		"UserId":          "用户IdXXXX",
		"Password":        "密码XXXX",
		"ConfirmPassword": "确认密码XXXX",
	}
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc  修改自己密码的使用
 */
type SysUserPwdContext struct {
	Password        string `validate:"required|minLen:6|maxLen:16" json:"password" `        // 密码
	ConfirmPassword string `validate:"required|minLen:6|maxLen:16" json:"confirmPassword" ` // 确认密码
	SystemId        int    `json:"systemId"`
}

// Messages 您可以自定义验证器错误消息
func (f SysUserPwdContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
		"minLen":   "{field}最小6位",
		"maxLen":   "{field}最大16位",
	}
}

// Translates 你可以自定义字段翻译
func (f SysUserPwdContext) Translates() map[string]string {
	return validate.MS{
		"Password":        "新密码",
		"ConfirmPassword": "确认密码",
	}
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 绑定用户手机
 */
type UserBindPhoneContext struct {
	Telephone string `validate:"required|cnMobile" json:"telephone"` // 手机号码
	PhoneCode string `validate:"required|len:6" json:"phoneCode"`    // 手机短信验证码
	SystemId  int    `json:"systemId"`
}

// Messages 您可以自定义验证器错误消息
func (f UserBindPhoneContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
		"len":      "{field}必须是6位",
	}
}

// Translates 你可以自定义字段翻译
func (f UserBindPhoneContext) Translates() map[string]string {
	return validate.MS{
		"Telephone": "绑定手机号码",
		"PhoneCode": "手机短信码",
	}
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 用户基本信息上下文
 */
type UserInfoContext struct {
	// 用户ID
	ID uint64 `json:"id"`
	// 昵称
	UserName string `validate:"required" json:"username"`
	// 头像
	Avatar string `validate:"required" json:"avatar"`
	// 地址
	Address string `json:"address"`
	// 手机号码
	Telephone string `json:"telephone"`
	// 签名
	Sign string `json:"sign"`
	// 生日
	BirthDay string `json:"birthDay"`
	// 性别 0女 1男 2 保密
	Male int ` json:"male"`
	// 省份
	Province string `json:"province"`
	// 城市
	City string `json:"city"`
	// 职业
	Job string `json:"job"`
	// 系统
	SystemId int `json:"systemId"`
}

// Messages 您可以自定义验证器错误消息
func (f UserInfoContext) Messages() map[string]string {
	return validate.MS{
		"required": "{field}不能为空",
	}
}

// Translates 你可以自定义字段翻译
func (f UserInfoContext) Translates() map[string]string {
	return validate.MS{
		"NickName": "昵称",
		"Avatar":   "头像",
	}
}

/**
 * @desc 关注列表分页
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserFocusPageContext struct {
	PageNo   int64  `json:"pageNo"`   // 第几页
	PageSize int64  `json:"pageSize"` // 每页显示多少条
	UserId   uint64 `json:"userId"`   // 用户ID
	SystemId int    `json:"systemId"`
}

/**
 * @desc 关注和取消关注
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserFocusContext struct {
	UserId        uint64     `json:"userId"`        // 用户ID
	Uuid          string     `json:"uuid"`          // 用户UUID
	FocusId       lib.BigInt `json:"focusId"`       // 关注的用户ID
	FocusUuid     string     `json:"focusUuid"`     // 关注的用户uuid
	Nickname      string     `json:"nickname"`      // 用户昵称
	Avatar        string     `json:"avatar"`        // 用户头像
	FocusNickname string     `json:"focusNickname"` // 用户昵称
	FocusAvatar   string     `json:"focusAvatar"`   // 用户头像
	SystemId      int        `json:"systemId"`
}

/**
 * @desc 关注和取消关注
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserFocusDelContext struct {
	Id       lib.BigInt `json:"id"`      // 删除ID
	UserId   uint64     `json:"userId"`  // 用户ID
	FocusId  lib.BigInt `json:"focusId"` // 关注的用户ID
	SystemId int        `json:"systemId"`
}

/**
 * @desc 用户统计模块相关
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserStateModelContext struct {
	Day      int    `json:"day"`    // 天数
	UserId   uint64 `json:"userId"` // 用户ID
	SystemId int    `json:"systemId"`
}

/**
 * @desc 关注和取消关注
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserFocusDelFansContext struct {
	Id       lib.BigInt `json:"id"`      // 删除ID
	UserId   uint64     `json:"userId"`  // 用户ID
	FocusId  uint64     `json:"focusId"` // 关注的用户ID
	SystemId int        `json:"systemId"`
}
