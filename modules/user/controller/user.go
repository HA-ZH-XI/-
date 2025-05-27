package web

import (
	"fmt"
	"github.com/gookit/validate"
	"ksd-social-api/commons/base/controller"
	contants2 "ksd-social-api/modules/common/sms/contants"
	"ksd-social-api/modules/user/contants"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/service"
	"ksd-social-api/utils/rdb"
)

type UserController struct {
	controller.BaseController
}

var userService = service.UserService{}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  查询用户明细
 */
func (web *UserController) GetUserInfo() {
	// 1： 获取当前登录的用户id，从token中获取的
	userId := web.GetUserId()
	// 2: 根据用户id查询用户西信息
	userVo := userService.GetUserInfo(userId, web.GetSystemId())
	// 3: 返回，因为查询没有错误信息，所以只有正确返回
	web.Ok(userVo)
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  查询用户明细
 */
func (web *UserController) GetUserBankInfo() {
	// 1： 获取当前登录的用户id，从token中获取的
	userId := web.GetUserId()
	// 2: 根据用户id查询用户西信息
	userVo := userService.GetUserByIDBank(userId, web.GetSystemId())
	// 3: 返回，因为查询没有错误信息，所以只有正确返回
	web.Ok(userVo)
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  修改用户密码
 */
func (web *UserController) UpdateUserPwd() {
	// 1: 开始绑定参数
	pwdContext := context.SysUserPwdContext{}
	web.BindJSON(&pwdContext)
	// 3: 对登录用户信息进行验证
	validation := validate.Struct(pwdContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}
	// 4: 判断两次输入的密码是否一致
	if pwdContext.ConfirmPassword != pwdContext.Password {
		web.FailCodeMsg(contants.USER_STATUS, contants.USER_UPATE_PWD_NO_SAME)
		return
	}

	pwdContext.SystemId = web.GetSystemId()
	// 5: 如果一致就开始修改用户的密码
	userId := web.GetUserId()
	// 6: 执行修改密码
	flag, err := userService.UpdateUserPassword(userId, web.GetSystemId(), pwdContext.Password)
	if flag {
		web.Ok("修改密码成功")
	} else {
		web.FailData(fmt.Sprintf("修改密码失败，错误原因是：%s", err.Error()))
	}
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 用户绑定手机
 */
func (web *UserController) UserBindPhone() {
	// 1: 定义接受参数上下文
	userBindPhoneContext := context.UserBindPhoneContext{}
	// 2: 参数绑定进去
	web.BindJSON(&userBindPhoneContext)
	// 3: 进行参数校验
	validation := validate.Struct(userBindPhoneContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}
	// 4: 根据用户传递进来的短信码和服务端存储的短信码然后进行比对，如果一致就修改，
	cacheKey := fmt.Sprintf(contants2.SMS_PHONE_KEY+"%s", userBindPhoneContext.Telephone)
	cachePhoneCode, _ := rdb.RdbGet(cacheKey)
	if len(cachePhoneCode) == 0 {
		web.FailCodeMsg(contants2.SMS_PHONE_STATUS, contants2.SMS_PHONE_NO_SEND)
		return
	}

	// 如果用户输入的短信码和服务端存储发送的短信码不一致， 如果不一致就返回错误信息：你输入的短信验证码有误！
	if userBindPhoneContext.PhoneCode != cachePhoneCode {
		web.FailCodeMsg(contants2.SMS_PHONE_STATUS, contants2.SMS_PHONE_NO_SAME)
		return
	}

	userBindPhoneContext.SystemId = web.GetSystemId()
	// 如果没有任何问题就开始绑定
	flag, err := userService.BindUserPhone(web.GetUserId(), web.GetSystemId(), userBindPhoneContext.Telephone)
	if flag {
		web.Ok("success")
	} else {
		web.FailCodeMsg(contants2.SMS_PHONE_STATUS, err.Error())
	}
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 修改用户个人资料
 */
func (web *UserController) UpdateUserInfo() {
	// 1: 定义接受参数上下文
	userInfoContext := context.UserInfoContext{}
	// 2: 参数绑定进去 ----如果接受字符串以外的参数 问题确定旧这个json不支持字符串日志---time.time
	web.BindJSON(&userInfoContext)
	// 3: 进行参数校验
	validation := validate.Struct(userInfoContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}

	userInfoContext.SystemId = web.GetSystemId()
	// 修改个人信息
	updateFlag, err := userService.UpdateUserInfo(web.GetUserId(), web.GetSystemId(), userInfoContext)
	if updateFlag {
		web.Ok("success")
	} else {
		web.FailCodeMsg(contants.USER_STATUS, err.Error())
	}
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc 统计我的消息和系统消息个数
 */
func (web *UserController) CountMessage() {
	web.Ok(userService.CountMessage(web.GetUserId(), web.GetSystemId()))
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc 统计我的消息和系统消息个数
 */
func (web *UserController) CountMessageAll() {
	web.Ok(userService.CountMessageAll(web.GetUserId(), web.GetSystemId()))
}
