package web

import (
	"github.com/gookit/validate"
	"ksd-social-api/modules/user/model/context"
)

/**
 * @desc 修改用户基本信息
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) UpdateUserSetting() {
	userInfoContext := context.UserInfoContext{}
	web.BindJSON(&userInfoContext)
	userInfoContext.ID = web.GetUserId()
	userInfoContext.SystemId = web.GetSystemId()
	uvo := userService.GetUserInfoByTelephone(userInfoContext.Telephone, web.GetSystemId())
	if uvo.ID > 0 {
		if uvo.ID != web.GetUserId() {
			web.FailCodeMsg(901, "该手机号码已被使用...")
			return
		}
	}

	uvo1 := userService.GetUserInfoByUsername(userInfoContext.UserName, web.GetSystemId())
	if uvo1.ID > 0 {
		if uvo1.ID != web.GetUserId() {
			web.FailCodeMsg(901, "该昵称已被使用...")
			return
		}
	}

	boolFlag := userCenterService.UpdateUserInfo(userInfoContext)
	web.Ok(boolFlag)
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 修改密码
 */
func (web *UserCenterController) UpdateUserPassword() {
	userPasswordContext := context.UserPasswordContext{}
	web.BindJSON(&userPasswordContext)
	userPasswordContext.UserId = web.GetUserId()
	validation := validate.Struct(userPasswordContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}

	if userPasswordContext.Password != userPasswordContext.Newpassword {
		web.FailCodeMsg(901, "两次输入密码不一致...")
		return
	}

	userPasswordContext.SystemId = web.GetSystemId()
	userCenterService.UpdateUserPassword(userPasswordContext)
	web.Ok("success")
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 收益设置
 */
func (web *UserCenterController) UpdateUserBank() {
	userBankContext := context.UserBankContext{}
	web.BindJSON(&userBankContext)
	userBankContext.UserId = web.GetUserId()
	validation := validate.Struct(userBankContext)
	if !validation.Validate() {
		web.FailWithValidatorData(validation)
		return
	}
	userBankContext.SystemId = web.GetSystemId()
	userCenterService.UpdateUserBank(userBankContext)
	web.Ok("success")
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 申请作者
 */
func (web *UserCenterController) UpdateUserAuthor() {
	userCenterService.UpdateUserAuthor(web.GetUserId(), web.GetSystemId())
	web.Ok("success")
}
