package web

import (
	"fmt"
	"github.com/gookit/validate"
	"github.com/mojocn/base64Captcha"
	"ksd-social-api/commons/base/controller"
	"ksd-social-api/modules/common/code/contants"
	"ksd-social-api/modules/login/model/context"
	"ksd-social-api/utils"
	"ksd-social-api/utils/rdb"
	"strconv"
)

type LoginController struct {
	controller.BaseController
}

/**
* @author feige
* @date 2023-09-26
* @desc 登录逻辑 -- 用户名or账号or手机 + 密码的方式
 */
func (web *LoginController) LoginAUPhoneByPassword() {
	// 1: 获取登录用户信息（账号、手机+密码+数字验证码）----context
	loginContext := context.LoginPasswordContext{}
	// 2: 获取参数信息把参数信息绑定到结构体中
	web.BindJSON(&loginContext)
	// 3: 对登录用户信息进行验证
	validation := validate.Struct(loginContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}

	// 计数的key,考虑到唯一性：可以在增加ip维度，最好的做法是在登录之前给每个使用网站或者app的用户生成一个uuid
	// key的改造：一定是使用ip
	userIp := web.GetIpAddr()
	countKey := fmt.Sprintf("LOGIN:CODE:ERRCOUNT:" + userIp)
	get, _ := rdb.RdbGet(countKey)
	errorCount, _ := strconv.ParseInt(get, 10, 64)
	// 这里要给每个key生成一个过期时间
	// 判断登录失败次数如果超过3次就出现验证码验证
	if errorCount >= 3 || loginContext.ErrorCount >= 3 {
		// 如果用户没有输入验证码或者codekey那么直接返回
		if len(loginContext.CaptchaId) == 0 ||
			len(loginContext.VerifyCode) == 0 {
			web.FailCodeMsg(contants.VAL_CODE_INPUT, contants.VAL_CODE_INPUT_MSG)
			return
		}

		//4： 验证码的校验
		captcha := base64Captcha.NewCaptcha(driver, store)
		// 参数1：是验证码key, 参数2 是验证码明文 参数3：true代表只要验签过参数1的key就立即失效。前端一般要重新生成一个id和验证码
		verify := captcha.Verify(loginContext.CaptchaId, loginContext.VerifyCode, true)
		if !verify {
			// 输入的验证码有误
			web.FailCodeMsg(contants.VAL_CODE_FAIL, contants.VAL_CODE_FAIL_MSG)
			return
		}
	}

	// 4: 开始执行业务登录逻辑
	loginContext.SystemId = web.GetSystemId()
	vo, validErr := loginService.LoginByAUPhoneAndPassword(&loginContext)
	if validErr != nil {
		// 账号和密码计数
		rdb.RdbIncr(countKey)
		// 把key设置一个过期时间，防止用户恶意输入请求，写入到缓存中，可以使用过期时间自动取清理那些错误的缓存的账号和信息
		if errorCount == 0 {
			rdb.RdbExKey(countKey, 120)
		}
		// 返回错误
		web.FailCodeMsgData(contants.USER_ERROR_CODE, validErr.Error(), errorCount+1)
		return
	}
	// 如果登录成功，就直接把缓存中的计数器给删除掉。
	rdb.RdbDel(countKey)
	web.Ok(vo)
}

/**
* @author feige
* @date 2023-09-26
* @desc 退出登录
 */
func (web *LoginController) ToLogout() {
	// 1: 获取需要退出的用户id
	uuid := web.GetUuid()
	// 2: 获取缓存中的key
	loginCacheKey := fmt.Sprintf("LOGIN:USER:%s", uuid)
	// 3: 然后直接删除数据
	rdb.RdbDel(loginCacheKey)
	// 4：退出成功
	web.Ok("success")
}

/**
* @author feige
* @date 2023-09-26
* @desc 创建一个token（给所有的接口进行安全校验，交易加秘返回，而且你这个token最好赋予意义）
 */
func (web *LoginController) CreateSecurityToken() {
	// 2: 获取需要退出的用户id----这里要加密
	uuid := "KSD_" + utils.GetUUID()
	// 5：退出成功
	web.Ok(uuid)
}
