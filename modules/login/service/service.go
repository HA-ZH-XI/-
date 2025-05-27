package service

import (
	"errors"
	"fmt"
	"ksd-social-api/commons/base/service"
	"ksd-social-api/modules/login/model/context"
	"ksd-social-api/modules/login/model/dto"
	"ksd-social-api/modules/login/model/vo"
	mapper2 "ksd-social-api/modules/msg/mapper"
	context3 "ksd-social-api/modules/msg/model/context"
	service2 "ksd-social-api/modules/msg/service"
	"ksd-social-api/modules/user/mapper"
	"ksd-social-api/modules/user/model"
	vo2 "ksd-social-api/modules/user/model/vo"
	"ksd-social-api/utils"
	"ksd-social-api/utils/rdb"
)

type LoginService struct {
	service.BaseService
}

var userMapper = mapper.UserMapper{}
var messagePointMapper = mapper2.MessagePointMapper{}
var messageMeService = service2.MessageMeService{}

/**
 * @author feige
 * @date 2023-09-29
 * @desc 登录接口 - 根据手机号码查询用户信息
 */
func (service *LoginService) LoginByAUPhoneAndPassword(ctx *context.LoginPasswordContext) (o *vo.LoginVo, err error) {
	// 1: 根据手机号码查询用户
	user := userMapper.GetUserByTelephone(ctx.Telephone, ctx.SystemId)
	if user == nil {
		// 2: 根据用户账号去数据库中查询用户是否存在，存在就返回用户信息，不存在就说明：输入的账号和密码有误
		user = userMapper.GetUserByAccount(ctx.Telephone, ctx.SystemId)
		if user == nil {
			// 2: 根据用户账号去数据库中查询用户是否存在，存在就返回用户信息，不存在就说明：输入的账号和密码有误
			user = userMapper.GetUserByUserName(ctx.Telephone, ctx.SystemId)
			if user == nil {
				return nil, errors.New("找不到用户信息!!!")
			}
		}
	}

	// 如果用户不为空，但是被拉黑了直接
	if user != nil && user.Forbidden == 1 {
		return nil, errors.New("该用户已被禁止了!!!")
	}

	// 3: 根据输入密码和用户数据库查询密码。进行比对，如果相同，就代表登录成功，如果不同输入的账号和密码有误
	// decodePassword := adr.RsaDEncryptString(ctx.Password)
	md5Password := utils.Md5Slat(ctx.Password, user.Slat)
	if user.Password != md5Password {
		return nil, errors.New("输入账号和密码有误!!!")
	}
	// 4: 开始把登录的用户信息，生成token进行返回
	loginVo := vo.LoginVo{}
	// 5: 一定静态uuid（注册得来）（可以做挤下线）
	loginVo.Uuid = user.Uuid
	loginVo.UserName = user.UserName
	loginVo.UserAvatar = user.Avatar
	loginVo.UserPhone = user.Telephone
	loginVo.UserAddress = user.Address
	loginVo.UserNickname = user.NickName
	// 把登录的用户信息写入到缓存中 --map
	cacheKey := fmt.Sprintf("LOGIN:USER:%s", loginVo.Uuid)
	// 这里把用用户需要的信息，放入缓存中
	userDto := dto.UserDto{}
	utils.CopyProperties(&userDto, user)
	// 开始把用户信息写入到缓存中
	rdb.RdbHSet(cacheKey, "user", utils.StructToJson(userDto))
	return &loginVo, nil
}

/**
 * @author feige
 * @date 2023-09-29
 * @desc 登录接口 - 根据手机和短信码进行登录
 */
func (service *LoginService) LoginByTelePhoneCode(ctx *context.LoginCodeContext) (o *vo.LoginVo, err error) {
	// 1: 根据手机号码查询用户
	user := userMapper.GetUserByTelephone(ctx.Telephone, ctx.SystemId)
	// 2: 根据用户账号去数据库中查询用户是否存在，存在就返回用户信息，不存在就说明：输入的账号和密码有误
	if user == nil {
		// 如果不存在，就注册一个用户
		dbUser := model.User{}
		dbUser.Slat = utils.GetUUID()
		dbUser.Telephone = ctx.Telephone
		dbUser.NickName = service.GetRandNickname()
		dbUser.Avatar = service.GetRandAvatar()
		dbUser.Account = ctx.Telephone
		dbUser.UserName = ctx.Telephone
		dbUser.BgImg = service.GetRandAvatar()
		dbUser.Address = "中国"
		dbUser.Password = ""
		dbUser.Sign = "Ta什么都没留下!"
		dbUser.BirthDay = "2024-04-01"
		dbUser.Male = 2
		dbUser.OpenId = ""
		dbUser.IsDeleted = 0
		dbUser.Active = 1
		dbUser.Forbidden = 0
		dbUser.BbsNum = 0
		dbUser.CourseNum = 0
		dbUser.FansNum = 0
		dbUser.GzsNum = 0
		dbUser.Uuid = fmt.Sprintf("%d%s", ctx.SystemId, utils.GetUUID())
		dbUser.VipFlag = 1
		dbUser.Cron = 10
		dbUser.Income = 0
		dbUser.SystemId = ctx.SystemId
		user = userMapper.SaveUser(dbUser)
		// 保存默认消息
		messagePointMapper.SaveMessagePonitDefault(user.ID, ctx.SystemId, user.Uuid)
	}

	// 如果用户不为空，但是被拉黑了直接
	if user != nil && user.Forbidden == 1 {
		return nil, errors.New("该用户已被禁止了!!!")
	}

	// 4: 开始把登录的用户信息，生成token进行返回
	loginVo := vo.LoginVo{}
	loginVo.Uuid = user.Uuid
	loginVo.UserName = user.UserName
	loginVo.UserAvatar = user.Avatar
	loginVo.UserPhone = user.Telephone
	loginVo.UserAddress = user.Address
	loginVo.UserNickname = user.NickName
	// 把登录的用户信息写入到缓存中 --map
	cacheKey := fmt.Sprintf("LOGIN:USER:%s", loginVo.Uuid)
	// 这里把用用户需要的信息，放入缓存中
	userDto := dto.UserDto{}
	utils.CopyProperties(&userDto, user)
	// 开始把用户信息写入到缓存中
	rdb.RdbHSet(cacheKey, "user", utils.StructToJson(userDto))
	return &loginVo, nil
}

/**
 * @author feige
 * @date 2023-09-29
 * @desc 登录接口 - 根据手机和短信码进行登录
 */
func (service *LoginService) RegByAccount(ctx *context.RegContext) (o *vo.LoginVo, err error) {
	dbUser := model.User{}
	// 如果不存在，就注册一个用户
	dbUser.Slat = utils.GetUUID()
	dbUser.NickName = service.GetRandNickname()
	dbUser.Avatar = service.GetRandAvatar()
	dbUser.Account = ctx.Account
	dbUser.UserName = ctx.Account
	dbUser.BgImg = service.GetRandAvatar()
	dbUser.Address = "中国"
	dbUser.Password = ""
	dbUser.Sign = "Ta什么都没留下!"
	dbUser.BirthDay = "2024-04-01"
	dbUser.Male = 2
	dbUser.OpenId = ""
	dbUser.IsDeleted = 0
	dbUser.Active = 1
	dbUser.Forbidden = 0
	dbUser.BbsNum = 0
	dbUser.CourseNum = 0
	dbUser.FansNum = 0
	dbUser.GzsNum = 0
	dbUser.Uuid = fmt.Sprintf("%d%s", ctx.SystemId, utils.GetUUID())
	dbUser.VipFlag = 1
	dbUser.Cron = 10
	dbUser.Income = 0
	dbUser.SystemId = ctx.SystemId
	// 密码加密
	dbUser.Password = utils.Md5Slat(ctx.Password, dbUser.Slat)
	// 注册方法
	saveUser := userMapper.SaveUser(dbUser)

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
		// 保存默认消息---rabbitmq---生成这 --- 消费者
		messagePointMapper.SaveMessagePonitDefault(saveUser.ID, ctx.SystemId, saveUser.Uuid)
		// 注册消息发送---rabbitmq 100ms
		messageMeContext := context3.MessageMeContext{}
		messageMeContext.UserId = saveUser.ID
		messageMeContext.Uuid = saveUser.Uuid
		messageMeContext.SystemId = ctx.SystemId
		messageMeContext.UserName = saveUser.NickName
		messageMeService.SaveMessageMeReg(&messageMeContext)
	}()

	return &loginVo, nil
}

/**
 * @author feige
 * @date 2023-09-29
 * @desc 登录接口 - 微信登录
 */
func (service *LoginService) RegByWeixin(ctx *context.WXLoginResp) (o *vo2.UserVo, err error) {
	dbUser := model.User{}
	// 如果不存在，就注册一个用户
	dbUser.Slat = utils.GetUUID()
	dbUser.NickName = service.GetRandNickname()
	dbUser.Avatar = service.GetRandAvatar()
	dbUser.Account = ctx.OpenId
	dbUser.UserName = ctx.OpenId
	dbUser.BgImg = service.GetRandAvatar()
	dbUser.Address = "中国"
	dbUser.Password = ""
	dbUser.Sign = "Ta什么都没留下!"
	dbUser.BirthDay = "2024-04-01"
	dbUser.Male = 2
	dbUser.IsDeleted = 0
	dbUser.Active = 1
	dbUser.Forbidden = 0
	dbUser.BbsNum = 0
	dbUser.CourseNum = 0
	dbUser.FansNum = 0
	dbUser.GzsNum = 0
	dbUser.Uuid = fmt.Sprintf("%d%s", ctx.SystemId, utils.GetUUID())
	dbUser.VipFlag = 1
	dbUser.Cron = 10
	dbUser.Income = 0
	dbUser.OpenId = ctx.OpenId
	dbUser.SystemId = ctx.SystemId
	// 密码加密
	dbUser.Password = ""
	// 注册方法
	saveUser := userMapper.SaveUser(dbUser)

	userVo := vo2.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, saveUser)

	return &userVo, nil
}
