package main

import (
	"fmt"
	"github.com/beego/beego/v2/task"
	initilization2 "ksd-social-api/commons/initilization"
	"ksd-social-api/commons/taskjob"
	"ksd-social-api/global"
	initilization "ksd-social-api/initilization"
	_ "ksd-social-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//1：记录环境信息
	global.Env = beego.BConfig.RunMode
	//2：设定环境的日志级别，默认是prod环境下的error级别，意思就是：只有程序报错了才会进行日志写入到文件和控制台中。
	loggerLevel := "error"
	// 根据环境来设置日志级别
	if global.Env == "dev" {
		// 如果是开发环境，我们要方便调试我们代码。所以建议使用debug或者info都可以
		loggerLevel = "debug"
	}
	fmt.Println("当前启动的环境是：" + global.Env)
	//3：解析自定义的配置文件
	initilization2.InitMyConfig()
	//4：初始化zap日志
	initilization.InitLogger(loggerLevel)
	//5：初始化redis
	initilization2.InitRedis()
	//6: 初始化数据库
	initilization.InitMYSQL()
	// oss文件存储
	initilization.InitOssClient()
	//7: 打印路由信息
	//在开发环境可以使用打印把定义和初始化成功的路由全部打印出来，方便你查询和确认你定义的路由是否生效
	if global.Env == "dev" {
		tree := beego.PrintTree()
		methods := tree["Data"].(beego.M)
		for k, v := range methods {
			fmt.Printf("%s => %v\n", k, v)
		}
	}

	//8: 启动定时任务
	taskjob.InitTask()
	task.StartTask()

	// 静态目录设置
	beego.SetStaticPath("/static", "resources")
	// 跨域问题
	beego.InsertFilter("/api/*", beego.BeforeRouter, filter.CorsFilter)
	// 白名单验证
	beego.InsertFilter("/api/*", beego.BeforeRouter, filter.AppidFilter)
	// 登录和不登录的安全校验token
	beego.InsertFilter("/api/*", beego.BeforeRouter, filter.TokenFilter)
	// 全局拦截器--登录拦截
	beego.InsertFilter("/api/v1/*", beego.BeforeRouter, filter.LoginFilter)
	//9: beego框架的运行和启动
	beego.Run()
}
