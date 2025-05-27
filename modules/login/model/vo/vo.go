package vo

type LoginVo struct {
	Uuid         string `json:"uuid"`       // 生成uuid
	UserName     string `json:"username"`   // 用户名字
	UserNickname string `json:"nickname"`   // 用户昵称
	UserAvatar   string `json:"avatar"`     // 头像
	UserPhone    string `json:"userPhone"`  // 手机
	UserAddress  string `json:"address"`    // 收货地址
	ErrorCount   int    `json:"errorCount"` // 错误处理
}
