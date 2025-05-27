package web

import (
	"encoding/json"
	"fmt"
	"ksd-social-api/commons/base/controller"
	"ksd-social-api/modules/common/sms/contants"
	"ksd-social-api/modules/login/model/context"
	"ksd-social-api/modules/login/model/dto"
	"ksd-social-api/modules/login/model/vo"
	context3 "ksd-social-api/modules/msg/model/context"
	"ksd-social-api/utils"
	"ksd-social-api/utils/rdb"
	"net/http"
)

type WeixinLoginController struct {
	controller.BaseController
}

//state生成器
//redis

/**
 * 微信登录
 * @author feige
 * @date 2024-01-17
 * @version 1.0
 * @desc
 */
// 微信扫码登录成功以后的回调地址
func (web *WeixinLoginController) WexinLogin() {
	// code ---> openid
	// 获取微信用户扫一扫以后登录的code信息
	code := web.GetString("code")
	// 获取到验证的state信息
	state := web.GetString("state")
	// 获取到验证的state信息
	systemId, _ := web.GetInt("systemid")

	// 通过code获取到openid的过程。
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	targetUrl := fmt.Sprintf(url, "wx26e1928e2b3d8260", "529b17a2083988d37c4f9111e693b085", code)
	// 发起的登录请求获取openid
	resp, err := http.Get(targetUrl)
	// 关闭链接
	defer resp.Body.Close()

	// 如果请求没有异常
	if err == nil {
		wxLoginResp := context.WXLoginResp{}
		wxLoginResp.SystemId = systemId
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&wxLoginResp)
		if err == nil {
			userId := web.GetUserId()
			if userId > 0 { // 如果是登录状态的话，就直接绑定微信登录
				userService.UpdateUserOpenId(userId, web.GetSystemId(), wxLoginResp.OpenId)
				web.Ok("success")
			} else {
				cacheKey := fmt.Sprintf("%s:%d:%s:%s", contants.WEIXIN_LOGIN_KEY, wxLoginResp.SystemId, state, web.GetIpAddr())
				fmt.Println("WexinLogin=====cacheKey", cacheKey)
				cacheCode, _ := rdb.RdbGet(cacheKey)
				fmt.Println("WexinLogin=====cacheCode", cacheKey)
				if len(cacheCode) > 0 {
					web.TplName = "login.html"
					web.Render()
					return
				}
				// 判断当前的openid是不是已经注册了
				user := userService.GetUserInfoByOpenId(wxLoginResp.OpenId, wxLoginResp.SystemId)
				if user == nil {
					// 开始注册
					userModel, _ := loginService.RegByWeixin(&wxLoginResp)
					// 更新openid
					rdb.RdbSetEX(cacheKey, userModel.Uuid, contants.WEIXIN_LOGIN_LIMIT_SECONDS)
					// 为什么返回页面，如果返回json,在前端有格式，就很丑
					web.TplName = "login.html"
					web.Render()
				} else {
					rdb.RdbSetEX(cacheKey, user.Uuid, contants.WEIXIN_LOGIN_LIMIT_SECONDS)
					web.TplName = "login.html"
					web.Render()
				}
			}

		}
	}
}

/**
 * 微信登录回调 1706018042899 1706018042899
 * @author feige
 * @date 2024-01-21
 * @version 1.0
 * @desc
 */
func (web WeixinLoginController) WexinLoginCallback() {
	// 获取到验证的state信息
	state := web.GetString("state")
	// 获取到验证的state信息
	systemId, _ := web.GetInt("systemid")

	loginCacheKey := fmt.Sprintf("%s:%d:%s:%s", contants.WEIXIN_LOGIN_KEY, systemId, state, web.GetIpAddr())
	uuidCache, _ := rdb.RdbGet(loginCacheKey)
	if len(uuidCache) == 0 {
		web.FailCodeMsg(602, "fail")
		return
	}
	saveUser := userService.GetUserInfoByUuid(uuidCache, systemId)
	if saveUser != nil {
		// 4: 开始把登录的用户信息，生成token进行返回
		loginVo := vo.LoginVo{}
		loginVo.Uuid = saveUser.Uuid
		loginVo.UserName = saveUser.UserName
		loginVo.UserAvatar = saveUser.Avatar
		loginVo.UserPhone = saveUser.Telephone
		loginVo.UserAddress = saveUser.Address
		loginVo.UserNickname = saveUser.NickName

		// 把登录的用户信息写入到缓存中 --map
		cacheKey := fmt.Sprintf("LOGIN:USER:%s", loginVo.Uuid)
		// 这里把用用户需要的信息，放入缓存中
		userDto := dto.UserDto{}
		utils.CopyProperties(&userDto, saveUser)
		// 开始把用户信息写入到缓存中
		rdb.RdbHSet(cacheKey, "user", utils.StructToJson(userDto))

		go func() {
			rdb.RdbDel(loginCacheKey)
			// 保存默认消息---rabbitmq---生成这 --- 消费者
			messagePointService.SaveMessagePonitDefault(saveUser.ID, systemId, saveUser.Uuid)
			// 注册消息发送---rabbitmq 100ms
			messageMeContext := context3.MessageMeContext{}
			messageMeContext.UserId = saveUser.ID
			messageMeContext.Uuid = saveUser.Uuid
			messageMeContext.SystemId = systemId
			messageMeContext.UserName = saveUser.NickName
			messageMeService.SaveMessageMeReg(&messageMeContext)
		}()

		fmt.Println("WexinLoginCallback=====uuidCache==== web ok")
		web.Ok(loginVo)
	}
}
