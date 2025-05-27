package router

import (
	beego "github.com/beego/beego/v2/server/web"
	web "ksd-social-api/modules/login/controller"
)

type LoginRouter struct{}

/**
 * @author feige
 * @date 2023-10-14
 * @version 1.0
 * @desc 无需登录
 */
func (router *LoginRouter) InitNoLoginRouter() beego.LinkNamespace {
	// 用户中心--子命名空间
	namespace := beego.NSNamespace("/login",
		// 密码模式登录
		beego.NSCtrlPost("toLogin", (*web.LoginController).LoginAUPhoneByPassword),
		// 短信模式登录
		beego.NSCtrlPost("phonecode", (*web.LoginController).LoginByPhoneCode),
		// 注册
		beego.NSCtrlPost("reg", (*web.LoginController).Reg),
		// 忘记密码
		beego.NSCtrlPost("setting/forget", (*web.LoginController).ForgetPassword),
		beego.NSCtrlGet("/weixin/login", (*web.WeixinLoginController).WexinLogin),
		beego.NSCtrlPost("/weixin/callback", (*web.WeixinLoginController).WexinLoginCallback),
	)
	return namespace
}

/**
 * @author feige
 * @date 2023-10-14
 * @version 1.0
 * @desc 需要登录
 */
func (router *LoginRouter) InitRouter() beego.LinkNamespace {
	// 用户中心--子命名空间
	namespace := beego.NSNamespace("/login",
		// 退出---拦截
		beego.NSCtrlPost("logout", (*web.LoginController).ToLogout),
		beego.NSCtrlPost("setting/phone", (*web.LoginController).SettingPhone),
	)
	return namespace
}
