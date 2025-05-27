package web

import "ksd-social-api/modules/user/model/context"

/**
 * 我的购买-文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyUserBuyBbsPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyUserBuyBbsPage(userCenterContext)
	web.Ok(p)
}

/**
 * 我的文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyTopicsPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyTopicsPage(userCenterContext)
	web.Ok(p)
}

/**
 * 足迹文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyBroswerTopicsPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyBroswerTopicsPage(userCenterContext)
	web.Ok(p)
}

/**
 * 收藏文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindBbsUserFavPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindBbsUserFavPage(userCenterContext)
	web.Ok(p)
}

/**
 * 喜欢文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindBbsUserLikePage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindBbsUserLikePage(userCenterContext)
	web.Ok(p)
}

/**
 * @desc 根据id删除 文章
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) DelTopicById() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.DelTopicById(bbsId, web.GetUserId(), web.GetSystemId()))
}

/**
 * @desc 修改文章状态-上架
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) UpdateTopicStatusUp() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.UpdateTopicStatusUp(bbsId, web.GetUserId(), web.GetSystemId()))
}

/**
 * @desc 修改文章状态--下架
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) UpdateTopicStatusDown() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.UpdateTopicStatusDown(bbsId, web.GetUserId(), web.GetSystemId()))
}

/**
 * @desc 彻底删除
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) RemoveTopicById() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.RemoveTopicById(bbsId, web.GetUserId(), web.GetSystemId()))
}

/**
 * @desc恢复删除
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) RecoverTopicById() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.RecoverTopicById(bbsId, web.GetUserId(), web.GetSystemId()))
}

/**
 * @控制是否允许评论 1允许 0关闭
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) UpdateCommentFlagTopicById() {
	commentFlag, _ := web.GetInt64("commentFlag")
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.UpdateCommentFlagTopicById(commentFlag, bbsId, web.GetUserId(), web.GetSystemId()))
}

/**
 * 文章取消收藏
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) CancelBbsFav() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.CancelBbsFav(web.GetUserId(), bbsId, web.GetSystemId()))
}

/**
 * 文章取消点赞
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) CancelBbsLike() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.CancelBbsLike(web.GetUserId(), bbsId, web.GetSystemId()))
}

/**
 * 删除浏览记录
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (web *UserCenterController) RemoveBbsBroswer() {
	bbsId, _ := web.GetUint64("id")
	web.Ok(userCenterService.RemoveBbsBroswer(web.GetUserId(), bbsId, web.GetSystemId()))
}
