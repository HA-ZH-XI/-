package context

import "github.com/gookit/validate"

/**
 * @author feige
 * @date 2023-10-09
 * @version 1.0
 * @desc  用于手机验证码登录
 */
type LoginCodeContext struct {
	Telephone string `validate:"required|cnMobile" json:"telephone"` // 手机号码
	PhoneCode string `validate:"required|len:6" json:"phoneCode"`    // 手机短信验证码
	CaptchaId string `json:"captchaId"`                              // 验证码的key
	Vcode     string `json:"vcode"`                                  // 具体输入的验证码
	SystemId  int    `json:"systemId"`
}

// Messages 您可以自定义验证器错误消息
func (f LoginCodeContext) Messages() map[string]string {
	return validate.MS{
		"required":           "{field}不能为空",
		"PhoneCode.len":      "{field}必须是%d位",
		"Telephone.cnMobile": "{field}输入不合法",
		"Vcode.length":       "{field}长度必须是%d位",
	}
}

// Translates 你可以自定义字段翻译
func (f LoginCodeContext) Translates() map[string]string {
	return validate.MS{
		"Telephone": "用户手机",
		"PhoneCode": "手机短信验证码",
		"CaptchaId": "验证码key",
		"Vcode":     "验证码",
	}
}
