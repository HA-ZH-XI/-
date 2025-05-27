package service

import (
	"ksd-social-api/commons/base/service"
	"ksd-social-api/modules/user/model/vo"
	"ksd-social-api/utils"
)

/**
 * 用户
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserVipService struct {
	service.BaseService
}

/**
 * 根据ID查询用户VIP信息
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (service *UserVipService) GetUserVipById(id int, systemId int) *vo.UserVipVo {
	userVip := userVipMapper.GetUserVipDetail(id, systemId)
	if nil != userVip {
		var userVipVo vo.UserVipVo
		utils.CopyProperties(&userVipVo, userVip)
		return &userVipVo
	}
	return nil
}

/**
 * 根据ID查询用户VIP信息
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (service *UserVipService) GetUserVipByCode(code string, systemId int) *vo.UserVipVo {
	userVip := userVipMapper.GetUserVipDetailByCode(code, systemId)
	if nil != userVip {
		var userVipVo vo.UserVipVo
		utils.CopyProperties(&userVipVo, userVip)
		return &userVipVo
	}
	return nil
}

/**
 * 查看用户vip列表
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (service *UserVipService) FindUserVipCardList(systemId int) []*vo.UserVipVo {
	userVipList := userVipMapper.FindUserVipList(systemId)
	var vos []*vo.UserVipVo
	for _, userVip := range userVipList {
		var userVipVo vo.UserVipVo
		utils.CopyProperties(&userVipVo, userVip)
		benefits := userVipMapper.FindUserBenefits(userVip.BenefitsIds, systemId)
		var benefitsVos []*vo.UserBenefitsVo
		for _, benefit := range benefits {
			var userBenefitVo vo.UserBenefitsVo
			utils.CopyProperties(&userBenefitVo, benefit)
			benefitsVos = append(benefitsVos, &userBenefitVo)
		}
		userVipVo.BenefitsIds = benefitsVos
		vos = append(vos, &userVipVo)
	}
	return vos
}

/**
 * 用户VIP身份过期设置
 * @author feige
 * @date 2023-12-17
 * @version 1.0
 * @desc
 */
func (service *UserVipService) UserVipSettingPeriod(userId uint64, systemId int) bool {
	flag, _ := userVipMapper.UpdateVIPPeriod(userId, systemId)
	return flag
}
