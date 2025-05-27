package service

import (
	cmapper "ksd-social-api/mods/course/mapper"
	mmapper "ksd-social-api/modules/msg/mapper"
	smapper "ksd-social-api/modules/social/mapper"
	"ksd-social-api/modules/user/mapper"
)

var userFocusMapper = mapper.UserFocusMapper{}
var userMapper = mapper.UserMapper{}
var userBuyBbsMapper = mapper.UserBuyBbsMapper{}
var userBuyCourseMapper = mapper.UserBuyCourseMapper{}
var userBuyVipMapper = mapper.UserBuyVipMapper{}
var userWalletMapper = mapper.UserWalletMapper{}
var userVipMapper = mapper.UserVipMapper{}
var userCenterMapper = mapper.UserCenterMapper{}
var userVipCodeMapper = mapper.UserVipCodeMapper{}
var userWalletCodeMapper = mapper.UserWalletCodeMapper{}

var courseMapper = cmapper.CourseMapper{}
var courseUserFavMapper = cmapper.CourseUserFavMapper{}
var courseUserLikeMapper = cmapper.CourseUserLikeMapper{}
var courseUserListMapper = cmapper.CourseUserListMapper{}

var bbsTopicMapper = smapper.BbsTopicMapper{}
var bbsUserLikeMapper = smapper.BbsUserLikeMapper{}
var bbsUserFavMapper = smapper.BbsUserFavMapper{}

var messageSystemMapper = mmapper.MessageSystemMapper{}
var messageMeMapper = mmapper.MessageMeMapper{}
