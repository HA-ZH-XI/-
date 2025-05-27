package web

import (
	"fmt"
	"github.com/gookit/validate"
	"github.com/mojocn/base64Captcha"
	"ksd-social-api/modules/common/code/contants"
	contants2 "ksd-social-api/modules/common/sms/contants"
	"ksd-social-api/modules/login/model/context"
	"ksd-social-api/utils/rdb"
)

/**
* @author feige
* @date 2023-09-26
* @desc 登录逻辑 -- 手机+短信验证码
 */
func (web *LoginController) LoginByPhoneCode() {
	// 1: 获取登录用户信息（账号、手机+密码+数字验证码）----context
	loginContext := context.LoginCodeContext{}
	// 2: 获取参数信息把参数信息绑定到结构体中
	web.BindJSON(&loginContext)
	// 3: 对登录用户信息进行验证
	validation := validate.Struct(loginContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}

	//4： 验证码的校验
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 参数1：是验证码key, 参数2 是验证码明文 参数3：true代表只要验签过参数1的key就立即失效。前端一般要重新生成一个id和验证码
	verify := captcha.Verify(loginContext.CaptchaId, loginContext.Vcode, false)
	if !verify {
		// 输入的验证码有误
		web.FailCodeMsg(602, "你输入的验证码有误！")
		return
	}

	// 根据用户输入的手机短信码和服务端发送的短信码进行比较
	cacheKey := fmt.Sprintf(contants2.SMS_PHONE_KEY+"%d:%s", web.GetSystemId(), web.GetIpAddr())
	smsCode, _ := rdb.RdbGet(cacheKey)
	if len(smsCode) == 0 || len(loginContext.PhoneCode) == 0 {
		web.FailCodeMsg(601, "请点击短信发送按钮，获取短信码!")
		return
	}

	// 如果用户输入的短信码和服务端存储短信不一致。
	if smsCode != loginContext.PhoneCode {
		web.FailCodeMsg(601, "您输入短信码有误!")
		return
	}

	// 4: 开始执行业务登录逻辑
	loginContext.SystemId = web.GetSystemId()
	vo, validErr := loginService.LoginByTelePhoneCode(&loginContext)
	if validErr != nil {
		// 返回错误
		web.FailCodeMsg(contants.USER_ERROR_CODE, fmt.Sprintf("%s", validErr.Error()))
		return
	}
	// 然后返回
	web.Ok(vo)
}
