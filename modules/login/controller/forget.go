package web

import (
	"fmt"
	"github.com/gookit/validate"
	contants2 "ksd-social-api/modules/common/sms/contants"
	context2 "ksd-social-api/modules/common/sms/context"
	"ksd-social-api/utils/rdb"
)

/**
 * 忘记密码
 * @author feige
 * @date 2024-01-17
 * @version 1.0
 * @desc
 */
func (web *LoginController) ForgetPassword() {
	// 1: 准备一个参数容器
	phoneContext := context2.ForgetPasswordContext{}
	// 2: 绑定参数
	web.BindJSON(&phoneContext)
	// 3: 对登录用户信息进行验证
	validation := validate.Struct(phoneContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}

	// 5: 使用redis缓存进行数据共享
	cacheKey := fmt.Sprintf(contants2.SMS_PHONE_KEY+"%s:%d:%s", phoneContext.Telephone, web.GetSystemId(), web.GetIpAddr())
	cacheCode, _ := rdb.RdbGet(cacheKey)
	if len(cacheCode) == 0 {
		web.FailCodeMsg(606, "请点击获取短信码按钮！")
		return
	}

	if cacheCode != phoneContext.Code {
		web.FailCodeMsg(604, "输入的手机短信码有误！")
		return
	}

	// 判断手机是否是否注册过
	userVo := userService.GetUserInfoByTelephone(phoneContext.Telephone, web.GetSystemId())
	if userVo != nil {
		rdb.RdbDel(cacheKey)
		userService.UpdateUserPassword(userVo.UserId, web.GetSystemId(), phoneContext.Password)
		web.Ok(true)
	} else {
		web.FailCodeMsg(605, "查无此手机账号！")
	}
}
