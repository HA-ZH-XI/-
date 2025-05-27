package web

import (
	"github.com/mojocn/base64Captcha"
	"ksd-social-api/modules/login/service"
	service2 "ksd-social-api/modules/msg/service"
	uservice "ksd-social-api/modules/user/service"
)

var loginService = service.LoginService{}
var userService = uservice.UserService{}
var messageMeService = service2.MessageMeService{}
var messagePointService = service2.MessagePointService{}

// 用于验证码验证的时候使用
var store = base64Captcha.DefaultMemStore
var driver = base64Captcha.DefaultDriverDigit
