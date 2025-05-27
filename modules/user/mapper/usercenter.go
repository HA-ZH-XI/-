package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	"ksd-social-api/modules/user/model/vo"
)

type UserCenterMapper struct {
}

/*
*

	 *
	 * @author feige
	 * @date 2024-01-08
	 * @version 1.0
	 * @desc
		{
			"bbscount":14855,
			"coursecount":2,
			"fancount":1,
			"gzcount":1,
		} == js object == go map/struct实例
*/
//func (mapper *UserCenterMapper) CountUserRelationState(userId uint64) *vo.UserStateCountVo {
//	var bbscount, coursecount, fanscount, gzcount int
//	// -- 用户发表的文章
//	orm.NewOrm().Raw("SELECT count(1) FROM xk_bbs_topics WHERE user_id = ? AND `status` = 1 AND is_deleted = 0").SetArgs(userId).QueryRow(&bbscount)
//	// -- 用户学习的课程
//	orm.NewOrm().Raw("SELECT count(1) FROM xk_user_buy_course WHERE user_id = ?").SetArgs(userId).QueryRow(&coursecount)
//	// -- 查询用户粉丝数量
//	orm.NewOrm().Raw("SELECT count(1) FROM xk_focus_user WHERE focus_id = ?").SetArgs(userId).QueryRow(&fanscount)
//	// -- 查询用户关注用户
//	orm.NewOrm().Raw("SELECT count(1) FROM xk_focus_user WHERE user_id = ?").SetArgs(userId).QueryRow(&gzcount)
//	// -- 查询用户关注用户
//	orm.NewOrm().Raw("SELECT count(1) FROM xk_focus_user WHERE user_id = ?").SetArgs(userId).QueryRow(&gzcount)
//
//	vo := vo.UserStateCountVo{}
//	vo.Bbscount = bbscount
//	vo.Coursecount = coursecount
//	vo.Fanscount = fanscount
//	vo.Gzcount = gzcount
//
//	return &vo
//}
func (mapper *UserCenterMapper) CountUserRelationStateChildren(userId uint64, systemId int) *vo.UserStateCountVo {
	var bbscount, coursecount, fanscount, gzcount int
	orm.NewOrm().Raw(`
		SELECT 
		(SELECT count(1) as snum FROM xk_bbs_topics WHERE user_id = ? and system_id = ? AND status = 1 AND is_deleted = 0) as bbscount,
		(SELECT count(1) as snum FROM xk_user_buy_course WHERE user_id = ? and system_id = ? ) as coursecount,
		(SELECT count(1) as snum FROM xk_focus_user WHERE focus_id = ? and system_id = ? ) as fanscount,
		(SELECT count(1) as snum FROM xk_focus_user WHERE user_id = ? and system_id = ? ) as gzcount
		FROM DUAL
	`).SetArgs(userId, systemId, userId, systemId, userId, systemId, userId, systemId).QueryRow(&bbscount, &coursecount, &fanscount, &gzcount)

	vo := vo.UserStateCountVo{}
	vo.Bbscount = bbscount
	vo.Coursecount = coursecount
	vo.Fanscount = fanscount
	vo.Gzcount = gzcount
	return &vo
}

func (mapper *UserCenterMapper) CountUserRelationState(userId uint64, systemId int) []*vo.UserStateCountAllVo {
	var userStateCountAllVos []*vo.UserStateCountAllVo
	orm.NewOrm().Raw(`
		SELECT count(1) as snum,'发表文章' as sname,"bbscount" as sfield FROM xk_bbs_topics WHERE user_id = ? and system_id =? AND status = 1 AND is_deleted = 0
		UNION ALL
		SELECT count(1) as snum,'学习课程' as sname,"coursecount" as sfield FROM xk_user_buy_course WHERE user_id = ? and system_id =? 
		UNION ALL
		SELECT count(1) as snum,'粉丝数' as sname,"fanscount" as sfield FROM xk_focus_user WHERE focus_id = ? and system_id =? 
		UNION ALL
		SELECT count(1)  as snum,'我的关注' as sname,"gzcount" as sfield FROM xk_focus_user WHERE user_id = ? and system_id =? 
	`).SetArgs(userId, systemId, userId, systemId, userId, systemId, userId, systemId).QueryRows(&userStateCountAllVos)

	return userStateCountAllVos
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  统计用户的每日发文章数量，学习课程，新增粉丝
 */
func (mapper *UserCenterMapper) CountUserModelState(userId uint64, systemId int, day int) []*vo.UserStateModelDataVo {
	// 多少天学习文章
	var vos []*vo.UserStateModelDataVo

	var bbsUserStateModelVo vo.UserStateModelDataVo
	var bbsUserStateModelChildrens []*vo.UserStateModelDataChildren
	orm.NewOrm().Raw(`
			SELECT
				DATE(update_time) AS datestr,
				COUNT(1) AS snum
			FROM
				xk_bbs_user_hits
			WHERE
				user_id = ? and system_id = ?
			AND DATE(update_time) >= DATE_SUB(CURRENT_DATE, INTERVAL ? DAY)
			GROUP BY datestr
		`).SetArgs(userId, systemId, day).QueryRows(&bbsUserStateModelChildrens)

	bbsUserStateModelVo.Name = "学习文章"
	bbsUserStateModelVo.Model = "topic"
	bbsUserStateModelVo.SystemId = systemId
	bbsUserStateModelVo.Value = bbsUserStateModelChildrens
	vos = append(vos, &bbsUserStateModelVo)

	// 多少天学习课程
	var courseUserStateModelVo vo.UserStateModelDataVo
	var courseUserStateModelChildrens []*vo.UserStateModelDataChildren
	orm.NewOrm().Raw(`
		SELECT
			DATE(update_time) AS datestr,
			COUNT(1) AS snum
		FROM
			xk_course_user_hits
		WHERE
			user_id = ? AND system_id = ?
		AND DATE(update_time) >= DATE_SUB(CURRENT_DATE, INTERVAL ? DAY)
		GROUP BY datestr
	`).SetArgs(userId, systemId, day).QueryRows(&courseUserStateModelChildrens)
	courseUserStateModelVo.Name = "学习课程"
	courseUserStateModelVo.Model = "course"
	courseUserStateModelVo.SystemId = systemId
	courseUserStateModelVo.Value = courseUserStateModelChildrens
	vos = append(vos, &courseUserStateModelVo)

	// 多少天新增粉丝
	var fansUserStateModelVo vo.UserStateModelDataVo
	var fansUserStateModelChildrens []*vo.UserStateModelDataChildren
	orm.NewOrm().Raw(`
			SELECT
				DATE(create_time) AS datestr,
				COUNT(1) AS snum
			FROM
				xk_focus_user
			WHERE
				focus_id = ? AND system_id = ?
			AND DATE(create_time) >= DATE_SUB(CURRENT_DATE, INTERVAL ? DAY)
			GROUP BY datestr
	`).SetArgs(userId, systemId, day).QueryRows(&fansUserStateModelChildrens)
	fansUserStateModelVo.Name = "新增粉丝"
	fansUserStateModelVo.Model = "fans"
	fansUserStateModelVo.SystemId = systemId
	fansUserStateModelVo.Value = fansUserStateModelChildrens
	vos = append(vos, &fansUserStateModelVo)

	// 多少天的文章曝光
	var newUserStateModelVo vo.UserStateModelDataVo
	var newUserStateModelChildrens []*vo.UserStateModelDataChildren
	// 这里补充一个sql
	orm.NewOrm().Raw(`
			SELECT
				DATE(update_time) AS datestr,
				COUNT(1) AS snum
			FROM
				xk_bbs_topics
			WHERE
				user_id = ? AND system_id = ?
			AND DATE(update_time) >= DATE_SUB(CURRENT_DATE, INTERVAL ? DAY)
			GROUP BY datestr
		`).SetArgs(userId, systemId, day).QueryRows(&newUserStateModelChildrens)
	newUserStateModelVo.Name = "文章曝光"
	newUserStateModelVo.Model = "news"
	newUserStateModelVo.SystemId = systemId
	newUserStateModelVo.Value = newUserStateModelChildrens
	vos = append(vos, &newUserStateModelVo)

	return vos
}
