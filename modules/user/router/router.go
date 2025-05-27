package router

import (
	beego "github.com/beego/beego/v2/server/web"
	web "ksd-social-api/modules/user/controller"
)

type UserRouter struct {
}

/**
 * @author feige
 * @date 2023-10-14
 * @version 1.0
 * @desc 无需登录
 */
func (router *UserRouter) InitNoLoginRouter() beego.LinkNamespace {
	// 用户中心--子命名空间
	namespace := beego.NSNamespace("/user")
	return namespace
}

/**
 * @author feige
 * @date 2023-10-14
 * @version 1.0
 * @desc 需要登录 --- 用户个人信息
 */
func (router *UserRouter) InitRouter() beego.LinkNamespace {
	namespace := beego.NSNamespace("/user",
		// 根据用户id查询用户明细
		beego.NSCtrlPost("/getUserInfo", (*web.UserController).GetUserInfo),
		beego.NSCtrlPost("/getUserBankInfo", (*web.UserController).GetUserBankInfo),
		// 修改个人密码
		beego.NSCtrlPost("/updatePwd", (*web.UserController).UpdateUserPwd),
		// 绑定用户手机
		beego.NSCtrlPost("/bind/userPhone", (*web.UserController).UserBindPhone),
		// 修改个人资料
		beego.NSCtrlPost("/update/userInfo", (*web.UserController).UpdateUserInfo),
		// 收益设置
		beego.NSCtrlPost("/update/password", (*web.UserCenterController).UpdateUserPassword),
		// 申请作者
		beego.NSCtrlPost("/update/author", (*web.UserCenterController).UpdateUserAuthor),
		// 修改密码
		beego.NSCtrlPost("/update/bank", (*web.UserCenterController).UpdateUserBank),

		// 关注和取消和在关注
		beego.NSCtrlPost("/focus/savecancel", (*web.UserFocusController).SaveCancelUserFocus),
		// 取消关注
		beego.NSCtrlPost("/focus/cancle", (*web.UserFocusController).CancelUserFocus),
		// 删除关注
		beego.NSCtrlPost("/focus/delgz", (*web.UserFocusController).DelUserFocus),
		// 删除粉丝
		beego.NSCtrlPost("/focus/delfans", (*web.UserFocusController).CancelUserFansFocus),
		// 用户粉丝列表
		beego.NSCtrlPost("/focus/fans", (*web.UserCenterController).FindUserFocusPage),
		// 用户关注列表
		beego.NSCtrlPost("/focus/gz", (*web.UserCenterController).FindUserFocusGzPage),

		// 用户身份兑换
		beego.NSCtrlPost("/code/vip/dh", (*web.UserVipCodeController).DuihuanUserVIP),
		// 用户身份兑换码生成
		beego.NSCtrlPost("/code/vip/create", (*web.UserVipCodeController).SaveUserVipCode),
		// 用户身份信息查找
		beego.NSCtrlPost("/code/vip/query", (*web.UserVipCodeController).QueryInfoByCode),

		// 用户充值
		beego.NSCtrlPost("/code/wallet/dh", (*web.UserWalletCodeController).DuihuanUserWallet),
		// 用户充值兑换码生成
		beego.NSCtrlPost("/code/wallet/create", (*web.UserWalletCodeController).SaveUserWalletCode),

		// 用户中心 --我的文章
		beego.NSCtrlPost("/center/bbs/list", (*web.UserCenterController).FindMyTopicsPage),
		// 足迹文章
		beego.NSCtrlPost("/center/bbs/broswer", (*web.UserCenterController).FindMyBroswerTopicsPage),
		// 收藏文章
		beego.NSCtrlPost("/center/bbs/fav", (*web.UserCenterController).FindBbsUserFavPage),
		// 点赞文章
		beego.NSCtrlPost("/center/bbs/like", (*web.UserCenterController).FindBbsUserLikePage),
		// 我购买文章
		beego.NSCtrlPost("/center/bbs/buy", (*web.UserCenterController).FindMyUserBuyBbsPage),
		// 删除浏览记录
		beego.NSCtrlPost("/center/bbs/remove/broswer", (*web.UserCenterController).RemoveBbsBroswer),
		// 取消点赞
		beego.NSCtrlPost("/center/bbs/cancel/like", (*web.UserCenterController).CancelBbsLike),
		// 取消收藏
		beego.NSCtrlPost("/center/bbs/cancel/fav", (*web.UserCenterController).CancelBbsFav),

		// 用户中心 --我的课程
		beego.NSCtrlPost("/center/course/list", (*web.UserCenterController).FindMyCoursePage),
		// 足迹课程
		beego.NSCtrlPost("/center/course/broswer", (*web.UserCenterController).FindMyCoursesBroswerPage),
		// 课程清单
		beego.NSCtrlPost("/center/course/items", (*web.UserCenterController).FindMyItemsCoursePage),
		// 收藏文章
		beego.NSCtrlPost("/center/course/fav", (*web.UserCenterController).FindCourseUserFavPage),
		// 点赞文章
		beego.NSCtrlPost("/center/course/like", (*web.UserCenterController).FindCourseUserLikePage),
		// 删除浏览记录
		beego.NSCtrlPost("/center/course/remove/broswer", (*web.UserCenterController).RemoveCourseHits),
		// 删除课程清单
		beego.NSCtrlPost("/center/course/remove/items", (*web.UserCenterController).CancelListCourse),
		// 取消点赞
		beego.NSCtrlPost("/center/course/cancel/like", (*web.UserCenterController).CanceLikeCourse),
		// 取消收藏
		beego.NSCtrlPost("/center/course/cancel/fav", (*web.UserCenterController).CancelFavCourse),

		// 上架文章
		beego.NSCtrlPost("/center/bbs/status/up", (*web.UserCenterController).UpdateTopicStatusUp),
		// 下架文章
		beego.NSCtrlPost("/center/bbs/status/down", (*web.UserCenterController).UpdateTopicStatusDown),
		// 删除文章
		beego.NSCtrlPost("/center/bbs/delete", (*web.UserCenterController).DelTopicById),
		// 彻底删除文章
		beego.NSCtrlPost("/center/bbs/remove", (*web.UserCenterController).RemoveTopicById),
		// 恢复删除
		beego.NSCtrlPost("/center/bbs/recover", (*web.UserCenterController).RecoverTopicById),
		// 评论控制
		beego.NSCtrlPost("/center/bbs/control", (*web.UserCenterController).UpdateCommentFlagTopicById),

		// 用户中心 -- 用户充值
		beego.NSCtrlPost("/center/wallet/records", (*web.UserCenterController).FindUserWalletRecords),
		// 用户基本资料修改
		beego.NSCtrlPost("/center/update/settings", (*web.UserCenterController).UpdateUserSetting),
		// 我的消息
		beego.NSCtrlPost("/center/message/me", (*web.UserCenterController).FindUserMeMessageRecords),
		// 一键已读
		beego.NSCtrlPost("/center/message/markall", (*web.UserCenterController).UpdateMessageMeMarkByUserId),
		// 已读
		beego.NSCtrlPost("/center/message/mark", (*web.UserCenterController).UpdateMessageMeMarkById),
		// 删除
		beego.NSCtrlPost("/center/message/del", (*web.UserCenterController).DelMessageMeById),
		// 系统消息
		beego.NSCtrlPost("/center/message/system", (*web.UserCenterController).FindUserSystemMessageRecords),
		// 统计消息-我的消息和系统消息
		beego.NSCtrlPost("/relation/countmsg", (*web.UserController).CountMessage),
		beego.NSCtrlPost("/relation/countmsgall", (*web.UserController).CountMessageAll),

		// 会员订单
		beego.NSCtrlPost("/order/course", (*web.UserCenterController).FindMyCoursesOrderPage),
		beego.NSCtrlPost("/order/bbs", (*web.UserCenterController).FindMyUserBuyBbsOrderPage),
		beego.NSCtrlPost("/order/vip", (*web.UserCenterController).FindMyUserBuyVipOrderPage),

		// 首页数据总览和统计分析
		beego.NSCtrlPost("/center/user/state", (*web.UserCenterController).CountUserRelationState),
		beego.NSCtrlPost("/center/user/statemodal", (*web.UserCenterController).CountUserModelState),

		// 查询vip列表
		beego.NSCtrlPost("/uservip/list", (*web.UserVipController).FindUserVipCardList),
	)
	return namespace
}
