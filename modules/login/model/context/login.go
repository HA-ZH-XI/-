package context

import "github.com/gookit/validate"

type LoginPasswordContext struct {
	Telephone  string `validate:"required" json:"telephone"` // 手机号码
	Password   string `validate:"required" json:"password"`  // 密码
	CaptchaId  string `json:"captchaId"`                     // 验证码的key
	VerifyCode string `json:"verifyCode"`                    // 具体输入的验证码
	ErrorCount int    `json:"errorCount"`                    //错误的次数
	SystemId   int    `json:"systemId"`
}

// Messages 您可以自定义验证器错误消息
func (f LoginPasswordContext) Messages() map[string]string {
	return validate.MS{
		"required":          "{field}不能为空",
		"Password.minLen":   "{field}不能少于%d位",
		"Password.maxLen":   "{field}最大不能超过%d位",
		"VerifyCode.length": "{field}长度必须是%d位",
	}
}

// Translates 你可以自定义字段翻译
func (f LoginPasswordContext) Translates() map[string]string {
	return validate.MS{
		"Telephone":  "用户手机",
		"Password":   "用户密码",
		"CaptchaId":  "验证码key",
		"VerifyCode": "验证码",
	}
}
