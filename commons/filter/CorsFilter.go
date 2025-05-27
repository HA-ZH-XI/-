package filter

import (
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

// 常量定义
var success = []byte("SUPPORT OPTIONS")

func CorsFilter(ctx *context.Context) {
	//跨域设置
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//允许访问源
	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	ctx.Output.Header("Access-Control-Max-Age", "3600")
	//允许所有类型的请求头访问
	ctx.Output.Header("Access-Control-Allow-Headers", "*")
	//设置响应头Access-Control-Allow-Credentials: false  不允许携带认证信息(如cookies)
	ctx.Output.Header("Access-Control-Allow-Credentials", "false")
	//OPTIONS请求处理如果是OPTIONS请求：
	//设置响应状态码为200(OK)
	//设置响应体为之前定义的success常量("SUPPORT OPTIONS")
	//忽略错误返回值(使用_)
	if ctx.Input.Method() == http.MethodOptions {
		// options请求，返回200
		ctx.Output.SetStatus(http.StatusOK)
		_ = ctx.Output.Body(success)
	}
}

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/*这段代码实现了一个完整的CORS(跨域资源共享)过滤器，主要功能包括：
允许所有来源的跨域请求
支持多种HTTP方法
设置预检请求的缓存时间
允许所有请求头
明确不允许携带认证信息
专门处理OPTIONS预检请求，直接返回成功响应*/
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/*使用场景
这种过滤器通常用于：
开发API服务时支持跨域请求
前后端分离架构中，前端应用需要访问不同域的后端API
微服务架构中服务间的跨域调用*/
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/*Access-Control-Allow-Origin: *在生产环境中可能需要更严格的设置
Access-Control-Allow-Headers: *在某些浏览器中可能不被支持
如果需要支持认证信息(cookies等)，需要将Access-Control-Allow-Credentials设为true，并且不能使用通配符(*)的Access-Control-Allow-Origin*/
