package service

import (
	"errors"
	"fmt"
	"ksd-social-api/commons/tools"
	"ksd-social-api/modules/user/model"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/model/vo"
	userUtils "ksd-social-api/modules/user/utils"
	"ksd-social-api/utils"
	"time"
)

type UserService struct{}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户id查询用户信息
 */
func (service *UserService) GetUserInfoByTelephone(telephone string, systemId int) (uvo *vo.UserVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserByPhone(telephone, systemId)
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 4: 返回
	return &userVo
}

/**
 * 根据openid查询
 * @author feige
 * @date 2023-12-28
 * @version 1.0
 * @desc
 */
func (service *UserService) GetUserInfoByOpenId(openid string, systemId int) (uvo *vo.UserVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserInfoByOpenId(openid, systemId)
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 4: 返回
	return &userVo
}

/**
 * 根据openid查询
 * @author feige
 * @date 2023-12-28
 * @version 1.0
 * @desc
 */
func (service *UserService) GetUserInfoByUuid(uuid string, systemId int) (uvo *vo.UserVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserInfoByUuid(uuid, systemId)
	fmt.Println("GetUserInfoByUuid=====user====" + tools.StructToJson(user))
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 4: 返回
	return &userVo
}

/**
 * 根据昵称查询
 * @author feige
 * @date 2023-12-28
 * @version 1.0
 * @desc
 */
func (service *UserService) GetUserInfoByUsername(username string, systemId int) (uvo *vo.UserVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserByUserName(username, systemId)
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 4: 返回
	return &userVo
}

/**
 * 根据账号查询
 * @author feige
 * @date 2023-12-28
 * @version 1.0
 * @desc
 */
func (service *UserService) GetUserInfoByAccount(account string, systemId int) (uvo *vo.UserVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserByAccount(account, systemId)
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 4: 返回
	return &userVo
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户id查询用户信息
 */
func (service *UserService) GetUserById(id uint64, systemId int) (uvo *vo.UserVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserByID(id, systemId)
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 4: 返回
	return &userVo
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户id查询用户信息
 */
func (service *UserService) GetUserByIDBank(id uint64, systemId int) (uvo *vo.UserBankVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserByIDBank(id, systemId)
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserBankVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 4: 返回
	return &userVo
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户id查询用户信息
 */
func (service *UserService) GetUserInfo(id uint64, systemId int) (uvo *vo.UserVo) {
	// 1: 根据用户id执行查询
	user := userMapper.GetUserByID(id, systemId)
	if user == nil {
		return nil
	}
	// 2: 把user的数据转换到uservo中返回
	userVo := vo.UserVo{}
	// 3: 把用户user复制到uservo中
	utils.CopyProperties(&userVo, user)
	// 生日如果不为nil才进行处理
	if len(user.BirthDay) > 0 {
		btime, _ := time.Parse("2006-01-02", user.BirthDay)
		// 获取星座
		userVo.Constellation = userUtils.GetZodiac(int(btime.Month()), btime.Day())
		// 获取用户生肖
		userVo.Twelve = userUtils.GetChineseZodiac(btime.Year())
	}

	// vip天数
	if &user.VipTime != nil {
		userVo.VipDays = int(user.VipTime.Sub(time.Now()).Hours()) / 24
	}
	// 注册天数
	userVo.RegDays = int(time.Now().Sub(user.CreateTime).Hours()) / 24
	// 4: 返回
	return &userVo
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  修改密码
 */
func (service *UserService) UpdateUserPassword(userId uint64, systemId int, password string) (bool, error) {
	// 使用新的uuid更新
	slat := utils.GetUUID()
	// 2: 获取用户唯一uuid和新密码进行加密
	return userMapper.UpdatePwd(userId, systemId, slat, utils.Md5Slat(password, slat))
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  修改openid
 */
func (service *UserService) UpdateUserOpenId(userId uint64, systemId int, openId string) (bool, error) {
	// 2: 获取用户唯一uuid和新密码进行加密
	return userMapper.UpdateUserOpenId(userId, systemId, openId)
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  修改手机
 */
func (service *UserService) UpdateTelephone(userId uint64, systemId int, telephone string) (bool, error) {
	// 2: 获取用户唯一uuid和新密码进行加密
	return userMapper.UpdateTelephone(userId, systemId, telephone)
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 绑定手机号码
 */
func (service *UserService) BindUserPhone(userid uint64, systemId int, telephone string) (bool, error) {
	user := userMapper.GetUserByID(userid, systemId)
	if user == nil {
		return false, errors.New("查无此用户!")
	}
	return userMapper.BindUserPhone(userid, systemId, telephone)
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc  修改个人信息
 */
func (service *UserService) UpdateUserInfo(userId uint64, systemId int, ctx context.UserInfoContext) (bool, error) {
	// 1: 把上下文的数据放入到数据模型中。切记不要把ctx直接丢到mapper去做。
	user := userMapper.GetUserByID(userId, systemId)
	if user != nil {
		// 直接把相同的数据进行拷贝
		utils.CopyProperties(user, ctx)
		// 修改用户基础信息
		return userMapper.UpdateUserInfo(user)
	}
	return false, errors.New("查无此用户")
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc  判断用户是否禁止
 */
func (service *UserService) IsForbiddenUser(userId uint64, systemId int) *model.User {
	// 1: 把上下文的数据放入到数据模型中。切记不要把ctx直接丢到mapper去做。
	user := userMapper.GetUserByID(userId, systemId)
	if user != nil && user.Forbidden == 1 {
		return nil
	}
	return user
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc  判断用户是否过期了
 */
func (service *UserService) IsUserVip(userId uint64, systemId int) bool {
	// 1: 把上下文的数据放入到数据模型中。切记不要把ctx直接丢到mapper去做。
	user := userMapper.GetUserByID(userId, systemId)
	if user != nil && user.VipFlag == 1 {
		return false
	}

	if user != nil && user.VipFlag > 1 && time.Now().Before(user.VipTime) {
		return false
	}

	return true
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc 统计我的消息和系统消息个数
 */
func (service *UserService) CountMessage(userId uint64, systemId int) []*vo.UserMessageVo {
	return userMapper.CountMessage(userId, systemId)
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc 统计我的消息和系统消息个数
 */
func (service *UserService) CountMessageAll(userId uint64, systemId int) []*vo.UserMessageVo {
	return userMapper.CountMessageAll(userId, systemId)
}
