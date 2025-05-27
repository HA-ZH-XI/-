package filter

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"ksd-social-api/modules/login/model/dto"
	"ksd-social-api/utils"
	"ksd-social-api/utils/rdb"
	"strings"
)

// 给那些不需要登录的接口在登录的状态获取用户信息---TokenFilter--没有/v1
// 而那些需要登录接口在登录状态获取用户----LoginFilter 有/v1
func TokenFilter(ctx *context.Context) {
	isUrlExist := ValidateURL(ctx.Request.URL.Path)
	if !isUrlExist {
		// 如果是登录接口，/api/v1/xxx就直接进入loginfilter去校验。
		contains := strings.Contains(ctx.Request.RequestURI, "v1")
		// 如果接口中包括的是v1就直接反行，去执行loginfilter
		if !contains {
			// 如果在登录的情况，获取登录的token
			Token := ctx.Request.Header["Authorization"]
			// 如果你确实在前端登录了len(Token) 。并且len(Token[0])也不是""
			if len(Token) > 0 && len(Token[0]) > 0 {
				userLoginUuid := Token[0]
				// 如果用户还在登录中就开始使用UUID到缓存中获取用户信息
				cacheLoginKey := fmt.Sprintf("LOGIN:USER:%s", userLoginUuid)
				cacheUser, _ := rdb.RdbHGet(cacheLoginKey, "user")
				if cacheUser != nil {
					//反序列
					userDto := dto.UserDto{}
					utils.JsonToStruct(cacheUser.(string), &userDto)
					// 如果没有。就获取用户信息，下放到路由处理方法中，这样给后续需要用户信息的地方可以直接获取
					ctx.Input.SetData("uuid", userLoginUuid)
					ctx.Input.SetData("userId", userDto.ID)
					ctx.Input.SetData("username", userDto.UserName)
					ctx.Input.SetData("avatar", userDto.Avatar)
					ctx.Input.SetData("phone", userDto.Telephone)
					ctx.Input.SetData("address", userDto.Address)
					ctx.Input.SetData("nickname", userDto.NickName)
				}
			}
		}
	}
}

//主要功能
//处理用户令牌(Token)的过滤器
//在用户已登录状态下获取用户信息，并将其存储在请求上下文中供后续处理使用。

/*逻辑
1.URL验证
2.URL不存在(即不需要特殊处理)时，才进入后续逻辑*/

/*后续逻辑
1.区分V1接口 v1接口走loginFilter  非v1接口走tokenTilter 以便于获得个性化信息

2. Token检查 从请求头获取Authorization字段 检查Token是否存在且不为空
3.从Redis获取用户信息  使用Token作为键从Redis哈希中获取用户信息
4.解析用户信息 将Redis中存储的JSON格式用户数据反序列化为UserDto对象
5.设置请求上下文数据  将用户信息存入请求上下文*/
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/*为什么要这样设计*/ //v1的作用
//v1路径扮演了一个重要的路由区分角色，它决定了请求是由TokenFilter处理还是由LoginFilter处理。

/*非v1路径
不需要强制登录
但如果用户提供了有效token，仍然可以获取用户信息
适用于公开API或部分功能对未登录用户可用的场景*/

/*v1路径：
必须登录才能访问
需要更严格的认证检查
适用于需要用户身份的核心功能*/

/*典型应用场景
/api/public/... → 公开API，由TokenFilter处理
/api/v1/user/profile → 需要登录的API，由LoginFilter处理
/api/items → 公开可读，但登录用户可以看到个性化内容*/

/*V1的3种主要作用
接口版本标识：
v1通常表示API的第一个版本(Version 1)
这是一种常见的API版本控制方式

权限控制分流：
包含v1的路径(/api/v1/xxx)会被跳过TokenFilter处理
这些请求将由LoginFilter处理，进行更严格的登录校验

安全层级区分：
非v1路径：可选认证（有token就获取用户信息，没有也可以访问）
v1路径：强制要求认证（必须登录才能访问）*/

/*这个过滤器的设计目的
根据注释，这个过滤器的设计目的是：
为不需要登录的接口提供在已登录状态下获取用户信息的能力

与LoginFilter分工明确：
TokenFilter处理非/v1路径的请求
LoginFilter处理/v1路径的请求(需要登录的接口)*/

/*在TokenFilter中，从Redis获取用户信息的过程是认证系统的核心部分*/
//Redis数据结构设计
/*键格式：LOGIN:USER:{uuid}
使用统一前缀LOGIN:USER:方便管理
{uuid}部分是用户的唯一标识符(从Token获取)*/
//数据获取过程
/*操作类型：使用HGet命令
从Hash中获取特定字段的值
这里获取的是"user"字段*/
//数据反序列化 json格式转换成string 将JSON字符串反序列化为UserDto结构体

//类型转换： 将从Redis获取的interface{}断言为string
//JSON解析 将JSON字符串反序列化为UserDto结构体

//用户信息存储设计分析 为什么使用Hash而不是String？

//使用Redis存储Token（会话令牌/访问令牌）是一种常见且高效的实践
//高性能与低延迟  读写速度极快（微秒级），适合高频访问的Token验证场景。 轻松应对大量用户同时认证的需求。
//会话状态管理 Token虽然由客户端携带，但需在服务端保存其有效性状态。Redis作为集中存储，方便多服务实例共享会话数据
//快速失效控制 用户登出时，直接删除Redis中的Token即可立即使其失效。 可灵活设置过期时间（如EXPIRE命令），实现自动清理。
//灵活的数据结构 Hash存储用户信息 可存储Token关联的权限、设备信息等元数据，并高效查询。
//增强安全性 即时吊销：发现可疑Token时，直接从Redis删除即可阻断攻击。 避免JWT的短板：若使用JWT Token，虽然无需存储，但无法主动失效，结合Redis可解决此问题（存储黑名单或有效Token）。

/*用户登录流程：
登录成功生成Token，存入Redis（Key=Token，Value=用户信息+过期时间）。
后续请求携带Token，网关/服务从Redis快速验证。
登出时删除Redis中的Token。*/

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//cacheLoginKey  通过拼接前缀和用户令牌，生成一个全局唯一的键，确保每个用户的登录数据在Redis中独立存储，避免冲突。
//通过键快速定位到用户的完整会话信息（如用户数据、过期时间、活跃时间等）。

/*实现高效认证流程
验证逻辑：
从请求头提取 Token → 生成 cacheLoginKey。
用该键查询Redis：
若存在且有效 → 用户已登录，允许访问。
若不存在或过期 → 用户未登录，拒绝访问。

性能优势：
Redis内存读写（微秒级）远快于数据库查询，支撑高并发认证。*/

/*cacheLoginKey 是 会话管理的核心枢纽，它：
通过结构化命名实现快速数据存取；
隔离不同用户和数据类型；
支撑高并发认证和安全控制；
为扩展功能（如多端登录、会话分析）提供基础。
其设计体现了典型的生产级会话管理方案，平衡了性能、安全性和可维护性。*/
