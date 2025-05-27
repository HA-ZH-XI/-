package web

import (
	"github.com/gookit/validate"
	"github.com/mojocn/base64Captcha"
	"ksd-social-api/modules/login/model/context"
)

/**
* @author feige
* @date 2023-09-26
* @desc 登录逻辑 -- 账号和密码的方式
 */
func (web *LoginController) Reg() {
	// 1: 获取登录用户信息（账号、手机+密码+数字验证码）----context
	regContext := context.RegContext{}
	// 2: 获取参数信息把参数信息绑定到结构体中
	web.BindJSON(&regContext)
	// 3: 对登录用户信息进行验证
	validation := validate.Struct(regContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}

	//4： 验证码的校验
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 参数1：是验证码key, 参数2 是验证码明文 参数3：true代表只要验签过参数1的key就立即失效。前端一般要重新生成一个id和验证码
	verify := captcha.Verify(regContext.CaptchaId, regContext.Code, true)
	if !verify {
		// 输入的验证码有误
		web.FailCodeMsg(602, "你输入的验证码有误！")
		return
	}

	// 比较两次密码是否相同
	if regContext.Password != regContext.ConfirmPassword {
		web.FailCodeMsg(601, "两次密码不一致，请确认!")
		return
	}

	// 账号校验是否注册
	userInfoByAccount := userService.GetUserInfoByAccount(regContext.Account, web.GetSystemId())
	if userInfoByAccount != nil {
		web.FailCodeMsg(602, "该账号已被注册使用!")
		return
	}

	// 注册--立即登录
	regContext.SystemId = web.GetSystemId()
	loginVo, _ := loginService.RegByAccount(&regContext)
	web.Ok(loginVo)
}
