package filter

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	R "ksd-social-api/commons/result"
	"ksd-social-api/modules/login/model/dto"
	"ksd-social-api/modules/user/service"
	"ksd-social-api/utils"
	"ksd-social-api/utils/rdb"
	"strconv"
)

var userVipService = service.UserVipService{}
var userService = service.UserService{}

func LoginFilter(ctx *context.Context) {
	isUrlExist := ValidateURL(ctx.Request.URL.Path)
	if !isUrlExist {
		Token := ctx.Request.Header["Authorization"]
		if len(Token) == 0 {
			ctx.Output.JSON(R.FailCodeMsg(60001, "请携带token进行访问"), true, false)
			return
		}

		// 获取头部用户的uid
		userLoginUuid := Token[0]
		if len(userLoginUuid) == 0 {
			ctx.Output.JSON(R.FailCodeMsg(60001, "请携带token进行访问"), true, false)
			return
		}

		// 判断当前的token是不是在redis的黑名单中，如果是就直接返回 ---退出 ---7
		cacheKey := fmt.Sprintf("USER:BLACKlIST:%s", userLoginUuid)
		cacheUuid, _ := rdb.RdbGet(cacheKey)
		if len(cacheUuid) > 0 && cacheUuid == userLoginUuid {
			// 如果过期了就直接返回，token失效
			ctx.Output.JSON(R.FailCodeMsg(60001, "请重新登录!"), true, false)
			return
		}

		// 如果用户还在登录中就开始使用UUID到缓存中获取用户信息
		cacheLoginKey := fmt.Sprintf("LOGIN:USER:%s", userLoginUuid)
		cacheUser, err := rdb.RdbHGet(cacheLoginKey, "user")
		if err == nil {
			/*反序列化*/
			userDto := dto.UserDto{}
			utils.JsonToStruct(cacheUser.(string), &userDto)

			SystemIdArr := ctx.Request.Header["Systemid"]
			systemId, _ := strconv.Atoi(SystemIdArr[0])

			// 如果用户不为空，但是被拉黑了直接
			user := userService.IsForbiddenUser(userDto.ID, systemId)
			if user == nil {
				// 如果过期了就直接返回，token失效
				ctx.Output.JSON(R.FailCodeMsg(60001, "账号已被禁止!"), true, false)
				return
			}
			// 如果没有。就获取用户信息，下放到路由处理方法中，这样给后续需要用户信息的地方可以直接获取
			ctx.Input.SetData("uuid", user.Uuid)
			ctx.Input.SetData("userId", user.ID)
			ctx.Input.SetData("username", user.UserName)
			ctx.Input.SetData("avatar", user.Avatar)
			ctx.Input.SetData("phone", user.Telephone)
			ctx.Input.SetData("nickname", user.NickName)
			ctx.Input.SetData("address", user.Address)
			ctx.Input.SetData("systemId", SystemIdArr[0])
			// 实时监听用户是否过期
			userVipService.UserVipSettingPeriod(user.ID, systemId)
		}
	}
}

//主要作用是为需要登录的接口提供统一的安全验证和用户信息管理
//逻辑
/* URL验证
URL验证码不合法进入逻辑*/

/*Token验证 是否有token
获取头部用户的uid
黑名单检查
用户信息获取 （redis获取通过UUID）*/

/*从Redis获取用户信息，Key格式为"LOGIN:USER:{token}"
将JSON格式的用户数据反序列化为UserDto对象
从请求头获取Systemid并转换为整数
检查用户是否被禁用 （token是否失效）
将用户信息设置到请求上下文中，供后续处理使用
检查并更新用户的VIP状态 （实时监听用户是否过期）*/

/*. 典型应用场景
需要登录的 API 接口：
如用户个人资料修改、订单查询等。

敏感操作拦截：
如支付前检查用户状态是否正常。

权限实时更新：
VIP 会员过期后立即失去特权。*/
