package mapper

import (
	"github.com/beego/beego/v2/client/orm"
	context2 "golang.org/x/net/context"
	"ksd-social-api/commons/page"
	cmodel "ksd-social-api/modules/course/model"
)

/**
 * 用户购买课程
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
type UserBuyCourseMapper struct {
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  保存用户购买课程信息
 */
func (service *UserBuyCourseMapper) SaveUserBuyCourse(userUserBuyCourse cmodel.UserBuyCourse) bool {
	// 创建orm对象
	err := orm.NewOrm().DoTx(func(ctx context2.Context, txOrm orm.TxOrmer) (err error) {
		// 保存用户课程订单
		_, err = txOrm.Insert(&userUserBuyCourse)
		// 同步购买和订阅的用户数量
		_, err = txOrm.Raw("update xk_course set buy_num = buy_num + 1  where id = ? and buy_num >= 0 and system_id = ?", userUserBuyCourse.CourseId, userUserBuyCourse.SystemId).Exec()
		return err
	})
	return err == nil
}

/**
 * 查询用户是否购买课程
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyCourseMapper) CountUserBuyCourse(userId uint64, courseId uint64, systemId int, payMethod int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_course WHERE user_id  = ? AND course_id = ? and pay_method = ? and system_id = ?", userId, courseId, payMethod, systemId).QueryRow(&total)
	// 返回
	if nil != err {
		return 0
	}
	return total
}

/**
 * 查询用户是否购买课程
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyCourseMapper) CountUserBuyCourseNo(userId uint64, courseId uint64, systemId int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_course WHERE user_id  = ? AND course_id = ? and system_id = ?", userId, courseId, systemId).QueryRow(&total)
	// 返回
	if nil != err {
		return 0
	}
	return total
}

/**
 * 统计课程购买人数
 * @author feige
 * @date 2023-11-23
 * @version 1.0
 * @desc
 */
func (service *UserBuyCourseMapper) CountUserBuyCourseId(courseId uint64, systemId int) int {
	// 创建orm对象
	mysql := orm.NewOrm()
	var total int
	err := mysql.Raw("SELECT count(1) FROM xk_user_buy_course WHERE course_id = ? and system_id = ?", courseId, systemId).QueryRow(&total)
	// 返回
	if nil != err {
		return 0
	}
	return total
}

/**
 * @author feige
 * @date 2023-10-12
 * @version 1.0
 * @desc 我学习的课程
 */
func (mapper *UserBuyCourseMapper) FindMyCoursesOrderPage(userId uint64, systemId int, pageNo int64, pageSize int64) (p *page.Page, err error) {
	// 准备容器对象，开始装载数据库数据
	userBuyCourseList := []cmodel.UserBuyCourse{}
	//创建orm对象
	mysql := orm.NewOrm()
	var total int64 = 0
	err1 := mysql.Raw(`
		select 
			count(1)
		from xk_course t1,xk_user_buy_course t2 
		where t1.id = t2.course_id and t1.status = 1 and t1.is_deleted = 0 and t2.user_id = ? and t2.system_id = ?
	`).SetArgs(userId, systemId).QueryRow(&total)
	if err1 != nil {
		return nil, err
	}
	// 重新换算分页和规则
	page := p.Page(pageNo, pageSize, total)
	// 分页查询
	_, err2 := mysql.Raw(`
		select 
			t2.id,
			t2.user_id,
			t2.course_id,
			t2.create_time,
			t2.update_time,
			t2.nickname,
			t2.avatar,
			t2.coursetitle,
			t2.coursecover,
			t2.description,
			t2.code,
			t2.price,
			t2.phone,
			t2.username,
			t2.address,
			t2.orderno,
			t2.uuid,
			t2.order_json,
			t2.tradeno,
			t2.pay_method,
			t2.pay_method_name
		from xk_course t1,xk_user_buy_course t2 
		where t1.id = t2.course_id and t1.status = 1 and t1.is_deleted = 0 and t2.user_id = ? and t2.system_id = ?
		order by t2.create_time desc
		limit ?,?
	`).SetArgs(userId, systemId, page.CurrentPage, page.PageSize).QueryRows(&userBuyCourseList)

	// 执行count查询
	if err1 != nil {
		return nil, err
	}
	// 6：把查询的数据放入到分页的records字段，准备返回
	page.Records = userBuyCourseList
	// 7: 如果没有找到直接返回
	if err2 != nil {
		return nil, err2
	}
	return page, nil
}
