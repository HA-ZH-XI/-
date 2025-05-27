package service

import (
	"ksd-social-api/commons/page"
	cmodel "ksd-social-api/modules/course/model"
	cvo "ksd-social-api/modules/course/model/vo"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/utils"
)

/**
 * 我的课程
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyCoursesPage(ctx context.UserCenterContext) *page.Page {
	coursePage, _ := courseMapper.FindMyCoursesPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	courseList := coursePage.Records.([]cmodel.Course)
	if courseList != nil {
		courseDetailVos := []cvo.CourseDetailVo{}
		for _, course := range courseList {
			courseVo := cvo.CourseDetailVo{}
			utils.CopyProperties(&courseVo, course)
			courseDetailVos = append(courseDetailVos, courseVo)
		}
		coursePage.Records = courseDetailVos
	}
	return coursePage
}

/**
 * 我的清单课程
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyCoursesListPage(ctx context.UserCenterContext) *page.Page {
	coursePage, _ := courseMapper.FindMyCoursesListPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	courseList := coursePage.Records.([]cmodel.Course)
	if courseList != nil {
		courseDetailVos := []cvo.CourseDetailVo{}
		for _, course := range courseList {
			courseVo := cvo.CourseDetailVo{}
			utils.CopyProperties(&courseVo, course)
			courseDetailVos = append(courseDetailVos, courseVo)
		}
		coursePage.Records = courseDetailVos
	}
	return coursePage
}

/**
 * 我的浏览记录课程
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindMyCoursesBroswerPage(ctx context.UserCenterContext) *page.Page {
	coursePage, _ := courseMapper.FindMyCoursesBroswerPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	courseList := coursePage.Records.([]cmodel.Course)
	if courseList != nil {
		courseDetailVos := []cvo.CourseDetailVo{}
		for _, course := range courseList {
			courseVo := cvo.CourseDetailVo{}
			utils.CopyProperties(&courseVo, course)
			courseDetailVos = append(courseDetailVos, courseVo)
		}
		coursePage.Records = courseDetailVos
	}
	return coursePage
}

/**
 * 收藏的课程
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindCourseUserFavPage(ctx context.UserCenterContext) *page.Page {
	coursePage, _ := courseUserFavMapper.FindCourseUserFavPage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	courseList := coursePage.Records.([]cmodel.Course)
	if courseList != nil {
		courseDetailVos := []cvo.CourseDetailVo{}
		for _, course := range courseList {
			courseVo := cvo.CourseDetailVo{}
			utils.CopyProperties(&courseVo, course)
			courseDetailVos = append(courseDetailVos, courseVo)
		}
		coursePage.Records = courseDetailVos
	}
	return coursePage
}

/**
 * 喜欢的课程
 * @author feige
 * @date 2023-12-14
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) FindCourseUserLikePage(ctx context.UserCenterContext) *page.Page {
	coursePage, _ := courseUserLikeMapper.FindCourseUserLikePage(ctx.UserId, ctx.SystemId, ctx.PageNo, ctx.PageSize)
	courseList := coursePage.Records.([]cmodel.Course)
	if courseList != nil {
		courseDetailVos := []cvo.CourseDetailVo{}
		for _, course := range courseList {
			courseVo := cvo.CourseDetailVo{}
			utils.CopyProperties(&courseVo, course)
			courseDetailVos = append(courseDetailVos, courseVo)
		}
		coursePage.Records = courseDetailVos
	}
	return coursePage
}

/**
 * 取消课程收藏
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) CancelFavCourse(userId uint64, courseId uint64, systemId int) bool {
	return courseUserFavMapper.CancelCourseUserFav(userId, courseId, systemId)
}

/**
 * 取消课程点赞
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) CanceLikeCourse(userId uint64, courseId uint64, systemId int) bool {
	return courseUserLikeMapper.CancelCourseUserLike(userId, courseId, systemId)
}

/**
 * 删除课程清单
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) CancelListCourse(userId uint64, courseId uint64, systemId int) bool {
	return courseUserListMapper.CancelCourseUserList(userId, courseId, systemId)
}

/**
 * 删除课程浏览记录
 * @author feige
 * @date 2024-01-13
 * @version 1.0
 * @desc
 */
func (service *UserCenterService) RemoveCourseHits(userId uint64, courseId uint64, systemId int) bool {
	return courseMapper.RemoveCourseHits(userId, courseId, systemId)
}
