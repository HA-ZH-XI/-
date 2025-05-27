package dto

import (
	"time"
)

type UserDto struct {
	// 主键
	ID uint64 `json:"id"`
	// 用户名
	UserName string `json:"username"`
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
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 用户状态
	Forbidden int `json:"forbidden"`
}
