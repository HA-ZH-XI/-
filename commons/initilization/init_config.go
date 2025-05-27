package initilization

import (
	beego "github.com/beego/beego/v2/server/web"
	"ksd-social-api/global"
)

func InitMyConfig() {
	//// 定义一个新的解析自定义配置的对象 cfg，自定义的文件名是：app-xxx.conf
	//cfg, err := config.NewConfig("ini", "conf/app-"+global.Env+".conf")
	//if err != nil {
	//	logs.Error(err)
	//}
	//// 这里就把自定义对象使用一个全局对象进行接管，方便后续进行获取和操作
	//global.Config = cfg

	beego.LoadAppConfig("ini", "conf/app-"+global.Env+".conf")

}
