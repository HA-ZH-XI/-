package service

import (
	"ksd-social-api/commons/page"
	bmodel "ksd-social-api/modules/social/model"
	bvo "ksd-social-api/modules/social/model/vo"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/utils"
)

/**
 * 我的文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyTopicsPage(ctx context.UserCenterContext) *page.Page {
	bbsTopicPage, _ := bbsTopicMapper.FindMyTopicPage(ctx.UserId, ctx.SystemId, ctx.IsDeleted, ctx.Status, ctx.PageNo, ctx.PageSize)
	bbsTopicList := bbsTopicPage.Records.([]bmodel.BbsTopics)
	if bbsTopicList != nil {
		bbsTopicsVos := []bvo.BbsTopicsVo{}
		for _, topics := range bbsTopicList {
			topicVo := bvo.BbsTopicsVo{}
			utils.CopyProperties(&topicVo, topics)
			bbsTopicsVos = append(bbsTopicsVos, topicVo)
		}
		bbsTopicPage.Records = bbsTopicsVos
	}
	return bbsTopicPage
}

/**
 * 我的购买文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyUserBuyBbsPage(ctx context.UserCenterContext) *page.Page {
	bbsTopicPage, _ := userBuyBbsMapper.FindMyUserBuyBbsPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	bbsTopicList := bbsTopicPage.Records.([]bmodel.BbsTopics)
	if bbsTopicList != nil {
		bbsTopicsVos := []bvo.BbsTopicsVo{}
		for _, topics := range bbsTopicList {
			topicVo := bvo.BbsTopicsVo{}
			utils.CopyProperties(&topicVo, topics)
			bbsTopicsVos = append(bbsTopicsVos, topicVo)
		}
		bbsTopicPage.Records = bbsTopicsVos
	}
	return bbsTopicPage
}

/**
 * 我的文章 - 浏览文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyBroswerTopicsPage(ctx context.UserCenterContext) *page.Page {
	bbsTopicPage, _ := bbsTopicMapper.FindMyBroswerTopicsPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	bbsTopicList := bbsTopicPage.Records.([]bmodel.BbsTopics)
	if bbsTopicList != nil {
		bbsTopicsVos := []bvo.BbsTopicsVo{}
		for _, topics := range bbsTopicList {
			topicVo := bvo.BbsTopicsVo{}
			utils.CopyProperties(&topicVo, topics)
			bbsTopicsVos = append(bbsTopicsVos, topicVo)
		}
		bbsTopicPage.Records = bbsTopicsVos
	}
	return bbsTopicPage
}

/**
 * 收藏的文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindBbsUserFavPage(ctx context.UserCenterContext) *page.Page {
	bbsTopicPage, _ := bbsUserFavMapper.FindBbsUserFavPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	bbsTopicList := bbsTopicPage.Records.([]bmodel.BbsTopics)
	if bbsTopicList != nil {
		bbsTopicsVos := []bvo.BbsTopicsVo{}
		for _, topics := range bbsTopicList {
			topicVo := bvo.BbsTopicsVo{}
			utils.CopyProperties(&topicVo, topics)
			bbsTopicsVos = append(bbsTopicsVos, topicVo)
		}
		bbsTopicPage.Records = bbsTopicsVos
	}
	return bbsTopicPage
}

/**
 * 喜欢的文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindBbsUserLikePage(ctx context.UserCenterContext) *page.Page {
	bbsTopicPage, _ := bbsUserLikeMapper.FindBbsUserLikePage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	bbsTopicList := bbsTopicPage.Records.([]bmodel.BbsTopics)
	if bbsTopicList != nil {
		bbsTopicsVos := []bvo.BbsTopicsVo{}
		for _, topics := range bbsTopicList {
			topicVo := bvo.BbsTopicsVo{}
			utils.CopyProperties(&topicVo, topics)
			bbsTopicsVos = append(bbsTopicsVos, topicVo)
		}
		bbsTopicPage.Records = bbsTopicsVos
	}
	return bbsTopicPage
}

/**
 * @desc 根据id删除
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) DelTopicById(id uint64, userId uint64, systemId int) bool {
	return bbsTopicMapper.DelTopicById(id, userId, systemId)
}

/**
 * @desc 修改文章状态--上架
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) UpdateTopicStatusUp(id uint64, userId uint64, systemId int) bool {
	return bbsTopicMapper.UpdateTopicStatus(1, id, userId, systemId)
}

/**
 * @desc 修改文章状态--下架
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) UpdateTopicStatusDown(id uint64, userId uint64, systemId int) bool {
	return bbsTopicMapper.UpdateTopicStatus(0, id, userId, systemId)
}

/**
 * @desc 彻底删除
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) RemoveTopicById(id uint64, userId uint64, systemId int) bool {
	return bbsTopicMapper.RemoveTopicById(id, userId, systemId)
}

/**
 * @desc恢复删除
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) RecoverTopicById(id uint64, userId uint64, systemId int) bool {
	return bbsTopicMapper.RecoverTopicById(id, userId, systemId)
}

/**
 * @控制是否允许评论 1允许 0关闭
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) UpdateCommentFlagTopicById(commentFlag int64, id uint64, userId uint64, systemId int) bool {
	return bbsTopicMapper.UpdateCommentFlagTopicById(commentFlag, id, userId, systemId)
}

/**
 * 文章取消收藏
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) CancelBbsFav(userId uint64, bbsId uint64, systemId int) bool {
	return bbsUserFavMapper.CancelBbsUserFav(userId, bbsId, systemId)
}

/**
 * 文章取消点赞
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) CancelBbsLike(userId uint64, bbsId uint64, systemId int) bool {
	return bbsUserLikeMapper.CancelBbsUserLike(userId, bbsId, systemId)
}

/**
 * 删除浏览记录
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
func (service *UserCenterService) RemoveBbsBroswer(userId uint64, bbsId uint64, systemId int) bool {
	return bbsTopicMapper.RemoveBbsBroswer(userId, bbsId, systemId)
}
