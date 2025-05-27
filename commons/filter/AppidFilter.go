package filter

import (
	"github.com/beego/beego/v2/server/web/context"
	R "ksd-social-api/commons/result"
)

func AppidFilter(ctx *context.Context) {
	/*1.调用 ValidateURL 检查请求的 URL 是否合法
	如果合法，直接放行；否则进入 AppID 验证逻辑。*/
	isUrlExist := ValidateURL(ctx.Request.URL.Path)
	if !isUrlExist {
		// 1: 获取头部信息的appid
		AppidArr := ctx.Request.Header["Appid"]
		SystemIdArr := ctx.Request.Header["Systemid"]
		////请求头验证
		if len(AppidArr) == 0 || len(AppidArr[0]) == 0 {
			ctx.Output.JSON(R.FailCodeMsg(60001, "不受信任的网站，非法请求"), true, false)
			return
		}

		// 2: 开始判断appid是不是在白名单中
		appid := AppidArr[0]
		systemId := SystemIdArr[0]
		//白名单验证
		exsit := ValidateAppid(appid)
		if !exsit {
			ctx.Output.JSON(R.FailCodeMsg(60001, "appid错误，非法请求"), true, false)
			return
		}
		//将 Systemid 存储到请求上下文中，后续中间件或控制器可通过 ctx.Input.GetData("systemId") 获取
		ctx.Input.SetData("systemId", systemId)
	}
}

//AppidFilter方法整体逻辑
// 1.URL 合法性检查 合法放行 不合法进入逻辑验证
/*2.逻辑验证包含
  a.请求头验证 请求头为空直接返回
  b.appid白名单验证  id不存在返回非法请求
  c.存放systemid在上下文中
*/

//==========================================================================
/*后续操作（结合上下文代码）：
调用 ValidateAppid(appid) 检查 Appid 是否在白名单中。
将 systemId 存储到请求上下文（ctx.Input.SetData），供后续业务逻辑使用。*/

//==========================================================================
/*获取请求头信息
AppidArr := ctx.Request.Header["Appid"]
从请求头中获取 Appid（通常用于标识调用方身份，如第三方服务或内部系统）。
SystemIdArr := ctx.Request.Header["Systemid"]
从请求头中获取 Systemid（可能用于标识子系统或业务模块）。

提取具体值
appid := AppidArr[0]
取 Appid 的第一个值（HTTP 头允许多个值，但此处默认只需第一个）。
systemId := SystemIdArr[0]
同理取 Systemid 的第一个值。*/

/*典型用例：
支付回调验证：微信/支付宝回调时携带 Appid 确认请求来源合法。
内部系统鉴权：微服务间调用通过 Appid + Systemid 区分权限。

//==========================================================================
/*核心目的：
通过 HTTP 头传递关键标识（Appid 用于鉴权，Systemid 用于业务路由）。

关键注意点：
必须处理头部缺失或空值的边界情况。
建议使用语义化变量名和日志监控。*/

//==========================================================================
//路径白名单 vs AppID 白名单的区别

/*路径白名单 （URL Whitelist）
作用：
通过校验请求的 URL 路径 是否在预定义的合法列表中，决定是否放行请求。

典型场景：
第三方回调验证（如微信支付、支付宝回调）。
公开API放行（如文档接口、健康检查）。*/
/*路径白名单特点：
简单直接：仅匹配路径字符串。
粗粒度控制：无法区分同一路径的不同调用方。
//==========================================================================
*/
/*AppID 白名单（AppID Whitelist）
作用：
通过校验请求头或参数中的 AppID 是否在合法列表中，验证调用方身份。

典型场景：
内部微服务间鉴权。
第三方API访问授权。*/

/*AppID 白名单特点：
身份导向：标识调用方而非请求内容。
细粒度控制：可为不同AppID分配不同权限*/
