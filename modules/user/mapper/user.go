package mapper

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"ksd-social-api/modules/user/model"
	"ksd-social-api/modules/user/model/context"
	"ksd-social-api/modules/user/model/vo"
	"time"
)

type UserMapper struct {
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  保存用户
 */
func (service *UserMapper) SaveUser(user model.User) (dbuser *model.User) {
	// 创建orm对象
	mysql := orm.NewOrm()
	// 保存用户
	_, err := mysql.Insert(&user)
	if err == nil {
		return &user
	}
	return nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据openid查询
 */
func (mapper *UserMapper) GetUserInfoByOpenId(openid string, systemId int) *model.User {
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select * from xk_user where open_id = ? and system_id = ?").SetArgs(openid, systemId).QueryRow(&user)
	if err != nil {
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户id查询用户明细
 */
func (mapper *UserMapper) GetUserByID(id uint64, systemId int) *model.User {
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select * from xk_user where id = ? and system_id = ?").SetArgs(id, systemId).QueryRow(&user)
	if err != nil {
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户id查询用户明细--收益设置
 */
func (mapper *UserMapper) GetUserByIDBank(id uint64, systemId int) *model.User {
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select id,uuid,user_name,nick_name,telephone,realname,idcard,idcardimgf,idcardimgc,alipaycode,weixincode,bankcode,bankimg,bankaddr from xk_user where id = ?  and system_id = ?").SetArgs(id, systemId).QueryRow(&user)
	if err != nil {
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户手机查询用户明细
 */
func (mapper *UserMapper) GetUserByPhone(telephone string, systemId int) *model.User {
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select * from xk_user where telephone = ? and system_id = ?").SetArgs(telephone, systemId).QueryRow(&user)
	if err != nil {
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据uuid查询用户明细
 */
func (mapper *UserMapper) GetUserInfoByUuid(uuid string, systemId int) *model.User {
	fmt.Println(fmt.Printf("参数是：%s:%d", uuid, systemId))
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select * from xk_user where uuid = ? and system_id = ?").SetArgs(uuid, systemId).QueryRow(&user)
	if err != nil {
		fmt.Println("GetUserInfoByUuid error" + err.Error())
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户名查询
 */
func (mapper *UserMapper) GetUserByUserName(username string, systemId int) *model.User {
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select * from xk_user where user_name = ? and system_id = ?").SetArgs(username, systemId).QueryRow(&user)
	if err != nil {
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据用户账号查询
 */
func (mapper *UserMapper) GetUserByAccount(account string, systemId int) *model.User {
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select * from xk_user where account = ? and system_id = ?").SetArgs(account, systemId).QueryRow(&user)
	if err != nil {
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  根据手机号码查询用户
 */
func (mapper *UserMapper) GetUserByTelephone(telephone string, systemId int) *model.User {
	// 创建orm对象
	mysql := orm.NewOrm()
	var user model.User
	err := mysql.Raw("select * from xk_user where telephone = ?  and system_id = ?").SetArgs(telephone, systemId).QueryRow(&user)
	if err != nil {
		return nil
	}
	return &user
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改手机
 */
func (mapper *UserMapper) UpdateTelephone(userId uint64, systemId int, telephone string) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET telephone = ? , update_time = now() WHERE id = ?  and system_id = ?").SetArgs(telephone, userId, systemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改密码
 */
func (mapper *UserMapper) UpdatePwd(userId uint64, systemId int, slat string, newPwd string) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET password = ? , slat = ? , update_time = now() WHERE id = ? and system_id = ?").SetArgs(newPwd, slat, userId, systemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 更新用户Openid
 */
func (mapper *UserMapper) UpdateUserOpenId(userId uint64, systemId int, openId string) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET open_id = ? , update_time = now() WHERE id = ? and system_id = ?").SetArgs(openId, userId, systemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  更新用户信息
 */
func UpdateInfo(user *model.User) int64 {
	// 创建orm对象
	mysql := orm.NewOrm()
	ruser := model.User{ID: user.ID, SystemId: user.SystemId}
	if mysql.Read(&ruser) == nil {
		ruser.Telephone = user.Telephone
		ruser.Address = user.Address
		ruser.Sign = user.Sign
		ruser.Male = user.Male
		if num, err := mysql.Update(&ruser, "telephone", "address", "sign", "male"); err == nil {
			return num
		}
	}
	return 0
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 绑定手机号码
 */
func (mapper *UserMapper) BindUserPhone(userId uint64, systemId int, telephone string) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET telephone = ? , update_time = now()  WHERE id = ? and system_id = ?").SetArgs(telephone, userId, systemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改密码
 */
func (mapper *UserMapper) UpdateUserPassword(userId uint64, systemId int, password string) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET password = ? , update_time = now()  WHERE id = ? and system_id = ?").SetArgs(password, userId, systemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改成作者
 */
func (mapper *UserMapper) UpdateUserAuthor(userId uint64, systemId int) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET author_flag = 1 , update_time = now()  WHERE id = ? and system_id = ?").SetArgs(userId, systemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改成作者
 */
func (mapper *UserMapper) UpdateUserBank(ctx context.UserBankContext) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET user_name =? , realname = ? ,telephone = ? , idcard = ? ,idcardimgf = ? ,idcardimgc = ? ,alipaycode = ? ,weixincode = ? ,bankcode = ? ,bankimg = ? ,bankaddr = ? ,update_time = now()  WHERE id = ? and system_id = ?").
		SetArgs(ctx.Realname, ctx.Realname, ctx.Telephone, ctx.Idcard, ctx.Idcardimgf, ctx.Idcardimgc, ctx.Alipaycode, ctx.Weixincode, ctx.Bankcode, ctx.Bankimg, ctx.Bankaddr, ctx.SystemId).Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改个人基础信息
 */
func (mapper *UserMapper) UpdateUserInfo(user *model.User) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET nick_name = ?,address = ?,sign = ?,bg_img = ?,birth_day = ?,avatar = ?,male = ? , update_time = now() WHERE id = ? and system_id = ?").
		SetArgs(user.NickName, user.Address, user.Sign, user.BgImg, user.BirthDay, user.Avatar, user.Male, user.ID, user.SystemId).
		Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected >= 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改身份
 */
func (mapper *UserMapper) UpdateUserVip(userId uint64, systemId int, vipTime time.Time, vipFlag int) (flag bool, err error) {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET vip_flag = ?,vip_time = ? , update_time = now() WHERE id = ? and system_id = ?").
		SetArgs(vipFlag, vipTime, userId, systemId).
		Exec()
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected >= 0, nil
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc 修改个人基础信息
 */
func (mapper *UserMapper) UpdateUserCenterInfo(user *model.User) bool {
	// 1: 获取链接
	db := orm.NewOrm()
	// 2: 开始执行修改密码的sql语句
	exec, _ := db.Raw("UPDATE xk_user SET user_name = ?,telephone = ?,address = ?,sign = ?,birth_day = ?,avatar = ?,male = ?,province=?,city=?,job=?, update_time = now() WHERE id = ? and system_id = ?").
		SetArgs(user.UserName, user.Telephone, user.Address, user.Sign, user.BirthDay, user.Avatar, user.Male, user.Province, user.City, user.Job, user.ID, user.SystemId).
		Exec()
	_, err := exec.RowsAffected()
	if err != nil {
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
func (mapper *UserMapper) CountMessage(userId uint64, systemId int) []*vo.UserMessageVo {
	vos := []*vo.UserMessageVo{}
	_, err := orm.NewOrm().Raw(`SELECT '我的消息' as label,'me' as ckey,COUNT(1) as mnum FROM xk_msg_me WHERE mark = 0 and user_id = ? and system_id = ?
	UNION ALL
	SELECT '系统消息' AS label,'system' AS ckey,COUNT( 1 ) AS mnum FROM xk_msg_system t1,xk_msg_point t2 WHERE t1.STATUS = 1 AND t1.s_type = 1 AND t2.user_id = ? and t2.system_id = ?  AND (t2.system_lasttime is null or t2.system_lasttime <= t1.update_time)
	UNION ALL
	SELECT '课程消息' AS label,'course' AS ckey,COUNT( 1 ) AS mnum FROM xk_msg_system t1,xk_msg_point t2 WHERE t1.STATUS = 1 AND t1.s_type = 2 AND t2.user_id = ?  and t2.system_id = ? AND (t2.course_lasttime is null or t2.course_lasttime <= t1.update_time)`, userId, systemId, userId, systemId, userId, systemId).QueryRows(&vos)
	if err != nil {
		return nil
	}
	return vos
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc 统计我的消息和系统消息个数
 */
func (mapper *UserMapper) CountMessageAll(userId uint64, systemId int) []*vo.UserMessageVo {
	vos := []*vo.UserMessageVo{}
	_, err := orm.NewOrm().Raw(`SELECT '我的消息' as label,'me' as ckey,COUNT(1) as mnum FROM xk_msg_me WHERE user_id = ? and system_id = ?
	UNION ALL
	SELECT '系统消息'as label,'system' as ckey, COUNT(1) as mnum FROM xk_msg_system WHERE STATUS = 1 and s_type = 2 and system_id = ?
	UNION ALL
	SELECT '课程消息'as label,'course' as ckey, COUNT(1) as mnum FROM xk_msg_system WHERE STATUS = 1 and s_type = 1 and system_id = ?`, userId, systemId, systemId, systemId).QueryRows(&vos)
	if err != nil {
		return nil
	}
	return vos
}
