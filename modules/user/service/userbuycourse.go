package service

/**
 * 用户购买课程
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserBuyCourseService struct {
}

/**
 * 查询用户是否购买课程
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyCourseService) CountUserBuyCourseNo(userId uint64, courseId uint64, systemId int) int {
	return userBuyCourseMapper.CountUserBuyCourseNo(userId, courseId, systemId)
}
