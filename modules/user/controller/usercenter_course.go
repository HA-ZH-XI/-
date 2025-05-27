package web

import "ksd-social-api/modules/user/model/context"

/**
 * 我的课程
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyCoursePage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyCoursesPage(userCenterContext)
	web.Ok(p)
}

/**
 * 课程清单
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyItemsCoursePage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyCoursesListPage(userCenterContext)
	web.Ok(p)
}

/**
 * 收藏文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindCourseUserFavPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindCourseUserFavPage(userCenterContext)
	web.Ok(p)
}

/**
 * 点赞文章
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindCourseUserLikePage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindCourseUserLikePage(userCenterContext)
	web.Ok(p)
}

/**
 * 足迹课程
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) FindMyCoursesBroswerPage() {
	userCenterContext := context.UserCenterContext{}
	web.BindJSON(&userCenterContext)
	userCenterContext.UserId = web.GetUserId()
	userCenterContext.SystemId = web.GetSystemId()
	p := userCenterService.FindMyCoursesBroswerPage(userCenterContext)
	web.Ok(p)
}

/**
 * 取消课程收藏
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) CancelFavCourse() {
	courseId, _ := web.GetUint64("id")
	web.Ok(userCenterService.CancelFavCourse(web.GetUserId(), courseId, web.GetSystemId()))
}

/**
 * 取消课程点赞
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) CanceLikeCourse() {
	courseId, _ := web.GetUint64("id")
	web.Ok(userCenterService.CanceLikeCourse(web.GetUserId(), courseId, web.GetSystemId()))
}

/**
 * 删除课程清单
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) CancelListCourse() {
	courseId, _ := web.GetUint64("id")
	web.Ok(userCenterService.CancelListCourse(web.GetUserId(), courseId, web.GetSystemId()))
}

/**
 * 删除课程浏览记录
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (web *UserCenterController) RemoveCourseHits() {
	courseId, _ := web.GetUint64("id")
	web.Ok(userCenterService.RemoveCourseHits(web.GetUserId(), courseId, web.GetSystemId()))
}
