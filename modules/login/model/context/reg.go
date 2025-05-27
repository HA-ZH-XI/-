package context

import "github.com/gookit/validate"

type RegContext struct {
	Account         string `validate:"required" json:"account"`                            // 注册手机
	Password        string `validate:"required|minLen:6|maxLen:16" json:"password"`        // 密码
	ConfirmPassword string `validate:"required|minLen:6|maxLen:16" json:"confirmPassword"` // 短信验证码
	SystemId        int    `json:"systemId"`
	Code            string `validate:"required" json:"code"`      // 验证码
	CaptchaId       string `validate:"required" json:"captchaId"` // 验证码的key
}

// Messages 您可以自定义验证器错误消息
func (f RegContext) Messages() map[string]string {
	return validate.MS{
		"required":               "{field}不能为空",
		"Password.minLen":        "{field}不能少于%d位",
		"Password.maxLen":        "{field}最大不能超过%d位",
		"ConfirmPassword.minLen": "{field}不能少于%d位",
		"ConfirmPassword.maxLen": "{field}最大不能超过%d位",
	}
}

// Translates 你可以自定义字段翻译
func (f RegContext) Translates() map[string]string {
	return validate.MS{
		"Account":         "用户账号",
		"Code":            "验证码",
		"CaptchaId":       "验证码ID",
		"Password":        "用户密码",
		"ConfirmPassword": "用户确认密码",
	}
}

// 信息登录的响应结果接受结构体
type WXLoginResp struct {
	OpenId       string `json:"openid"`
	UnionId      string `json:"unionid"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	SystemId     int    // 系统ID
}
