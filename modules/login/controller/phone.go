package web

import (
	"fmt"
	"github.com/gookit/validate"
	contants2 "ksd-social-api/modules/common/sms/contants"
	context2 "ksd-social-api/modules/common/sms/context"
	"ksd-social-api/utils/rdb"
)

/**
 * 设置手机
 * @author feige
 * @date 2024-01-17
 * @version 1.0
 * @desc
 */
func (web *LoginController) SettingPhone() {
	// 1: 准备一个参数容器
	phoneContext := context2.SettingSMSPhoneContext{}
	// 2: 绑定参数
	web.BindJSON(&phoneContext)
	// 3: 对登录用户信息进行验证
	validation := validate.Struct(phoneContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}

	// 如果已经绑定手机直接返回
	//uvo := userService.GetUserById(web.GetUserId())
	//if len(uvo.Telephone) > 0 {
	//	web.FailCodeMsg(607, "此账号已绑定手机！")
	//	return
	//}

	// 5: 使用redis缓存进行数据共享
	cacheKey := fmt.Sprintf(contants2.SMS_PHONE_KEY+"%d:%s:%d", web.GetUserId(), phoneContext.Telephone, web.GetSystemId())
	cacheCode, _ := rdb.RdbGet(cacheKey)
	if len(cacheCode) == 0 {
		web.FailCodeMsg(606, "请点击获取短信码按钮！")
		return
	}

	// 把服务端缓存的短信和用户输入短信码比对，如果不一致就提示：“输入的手机短信码有误”
	if cacheCode != phoneContext.Code {
		web.FailCodeMsg(604, "输入的手机短信码有误！")
		return
	}

	// 判断手机是否是否注册过
	userVo := userService.GetUserInfoByTelephone(phoneContext.Telephone, web.GetSystemId())
	if userVo != nil {
		web.FailCodeMsg(603, "你的手机已被使用！")
		return
	}

	// 开始更新用户的手机号码
	flag, err := userService.UpdateTelephone(web.GetUserId(), web.GetSystemId(), phoneContext.Telephone)
	if err != nil {
		web.FailCodeMsg(605, "手机设置失败！")
	}
	web.Ok(flag)
}
