package controller

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gookit/validate"
	"ksd-social-api/utils"
	"regexp"
	"strconv"
	"strings"
)

var (
	NICKNAME_DATA        = "锦时璃::青山锁::薄荷蓝::跟风走::遇晚星::梦旅人::风飞扬::眉间雪::愿相衬::安若汐::陌上花::凝残月::断秋风::龙吟凤::无眉缺::醉红颜::陌潇潇::莫阑珊::孔风雷::领青瓷::暮成雪::江满样::小悸动::几世寻::了无痕::月长歌::浅时光::北末染::青无锋::断莫离::醉晨色::梦里人::清衡子::君子棱::予星河::柚绿时::槑小猫::玉添花::漫山谷::邯郸道::挽木琴::梓心花::邢馨月::长歌行::撑青伞::北星缘::小米粒::小甜甜::临晚心::忆往昔::青柠檬::湘歌简::夜如影::尘世客::望北海::苏浅依::长青诗::凉堇雨::陌念念::程染筱::终不悔::画卿颜::世中仙::叮叮当::饮晚风::七里笙::枕边书::关启燕::梁子玉::苏九凉::解相思::温卷猫::明月归::避风港::陈茵沫::时空记::顾子兮::小酒窝::意容颜::桃米水::小白兔::槐不语::白衫辞::无线侠::恰梦中::怀中喵::沐卿月::千韵惘::小兔叽::吹羌笛::忆倾城::曾何时::箫外月::暮云深::蓝蝶恋::夏雪若::夜未央::跳跳羊::冷月魄::青衫隐"
	AVATAR_DATA          = "/avatar/01.png::/avatar/02.png::/avatar/03.png::/avatar/04.png::/avatar/05.png::/avatar/06.png::/avatar/07.png::/avatar/08.png::/avatar/09.png::/avatar/10.png"
	SUCCESS_CODE         = 20000
	FAIL_CODE            = 40001
	SUCCESS_MSG          = "success"
	FAIL_MSG             = "fail"
	VALIDATOR_MAP        = map[string]string{"code": "701", "msg": "属性验证有误"}
	BINDING_PAMATERS_MAP = map[string]string{"code": "702", "msg": "参数绑定有误"}
)

/*父控制器 -- 继承beego.Controller*/
type BaseController struct {
	beego.Controller
}

/*统一返回*/
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (web *BaseController) Ok(data any) {
	data = web.TransferData(data)
	web.Data["json"] = &Response{Code: SUCCESS_CODE, Message: SUCCESS_MSG, Data: data}
	web.ServeJSON()
}

func (web *BaseController) OkCode(code int, data any) {
	data = web.TransferData(data)
	web.Data["json"] = &Response{Code: code, Message: SUCCESS_MSG, Data: data}
	web.ServeJSON()
}

func (web *BaseController) OkCodeMsg(msg string, data any) {
	data = web.TransferData(data)
	web.Data["json"] = &Response{Code: SUCCESS_CODE, Message: msg, Data: data}
	web.ServeJSON()
}

func (web *BaseController) Fail() {
	web.Data["json"] = &Response{Code: FAIL_CODE, Message: FAIL_MSG, Data: nil}
	web.ServeJSON()
}

func (web *BaseController) FailData(data any) {
	web.Data["json"] = &Response{Code: FAIL_CODE, Message: FAIL_MSG, Data: data}
	web.ServeJSON()
}

func (web *BaseController) FailCodeMsg(code int, msg string) {
	web.Data["json"] = &Response{Code: code, Message: msg, Data: nil}
	web.ServeJSON()
}

func (web *BaseController) FailCodeMsgData(code int, msg string, data any) {
	data = web.TransferData(data)
	web.Data["json"] = &Response{Code: code, Message: msg, Data: data}
	web.ServeJSON()
}

func (web *BaseController) FailWithValidatorData(validate *validate.Validation) {
	all := validate.Errors.All()
	one := validate.Errors.One()
	code, _ := strconv.ParseInt(VALIDATOR_MAP["code"], 10, 0)
	web.Data["json"] = &Response{Code: int(code), Message: one, Data: all}
	web.ServeJSON()
}

func (web *BaseController) FailWithValidatorMap(validatorMsg map[string]string) {
	code, _ := strconv.ParseInt(VALIDATOR_MAP["code"], 10, 0)
	web.Data["json"] = &Response{Code: int(code), Message: VALIDATOR_MAP["msg"], Data: validatorMsg}
	web.ServeJSON()
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  获取登录的token中的用户id
 */
func (web *BaseController) GetUserId() uint64 {
	data := web.Ctx.Input.GetData("userId")
	if data != nil {
		return data.(uint64)
	}
	return 0
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 获取登录的UUID
 */
func (web *BaseController) GetUuid() string {
	data := web.Ctx.Input.GetData("uuid")
	if data != nil {
		return data.(string)
	}
	return ""
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  获取登录的token中的用户名字
 */
func (web *BaseController) GetUserName() string {
	data := web.Ctx.Input.GetData("username")
	if data != nil {
		return data.(string)
	}
	return ""
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  获取登录的token中的用户头像
 */
func (web *BaseController) GetUserAvatar() string {
	data := web.Ctx.Input.GetData("avatar")
	if data != nil {
		return data.(string)
	}
	return ""
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  获取用户手机
 */
func (web *BaseController) GetUserPhone() string {
	data := web.Ctx.Input.GetData("phone")
	if data != nil {
		return data.(string)
	}
	return ""
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  获取用户地址
 */
func (web *BaseController) GetUserAddress() string {
	data := web.Ctx.Input.GetData("address")
	if data != nil {
		return data.(string)
	}
	return ""
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  获取用户昵称
 */
func (web *BaseController) GetUserNickname() string {
	data := web.Ctx.Input.GetData("nickname")
	if data != nil {
		return data.(string)
	}
	return ""
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  获取系统
 */
func (web *BaseController) GetSystemId() int {
	data := web.Ctx.Input.GetData("systemId")
	if data != nil {
		atoi, _ := strconv.Atoi(data.(string))
		return atoi
	}
	return 0
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 对param参数进行转换
 */
func (web *BaseController) ParseParamInt(key string, def ...int) (int, error) {
	strv := web.Ctx.Input.Param(":" + key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.Atoi(strv)
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc 对param参数进行转换
 */
func (web *BaseController) ParseParamUint64(key string, def ...uint64) (uint64, error) {
	strv := web.Ctx.Input.Param(":" + key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseUint(strv, 10, 64)
}

/**
 * @author feige
 * @date 2023-10-08
 * @desc  对param参数进行转换
 */
func (web *BaseController) ParseParamInt64(key string, def ...int64) (int64, error) {
	strv := web.Ctx.Input.Param(":" + key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseInt(strv, 10, 64)
}

/**
 * @author feige
 * @date 2023-10-08
 * @version 1.0
 * @desc json参数的绑定
 */
func (web *BaseController) BindJSONToStruct(obj interface{}) error {
	return json.Unmarshal(web.Ctx.Input.RequestBody, obj)
}

/*
*
  - 获取用户IP地址
  - @author feige
  - @date 2023-11-07
  - @version 1.0
  - @desc
    由于前端使用了 nginx 做反向代理，所以通过 beego 框架内 提供的方法
    Input.IP()获取到的IP是内网IP。不满足需求
    nginx 代理中已经设置
    location / {
    proxy_set_header   X-real-ip $remote_addr;
    proxy_pass http://upstream/;
    }
*/
func (web *BaseController) GetIpAddr() string {
	header := web.Ctx.Request.Header
	ip := header.Get("x-forwarded-for")

	if ip == "" || len(ip) == 0 || "unknown" == ip {
		ip = header.Get("Proxy-Client-IP")
	}

	if ip == "" || len(ip) == 0 || "unknown" == ip {
		ip = header.Get("X-Forwarded-For")
	}

	if ip == "" || len(ip) == 0 || "unknown" == ip {
		ip = header.Get("WL-Proxy-Client-IP")
	}

	if ip == "" || len(ip) == 0 || "unknown" == ip {
		ip = header.Get("X-Real-IP")
	}

	if ip == "" || len(ip) == 0 || "unknown" == ip {
		ip = web.Ctx.Input.IP()
	}

	if "0:0:0:0:0:0:0:1" == ip || strings.Contains(ip, "::1") {
		return "127.0.0.1"
	} else {
		return ip
	}
}

/**
 * @desc 解决精度丢失的问题
 * @author feige
 * @date 2023-11-17
 * @version 1.0
 */
func (web *BaseController) TransferData(data any) any {
	j, _ := json.Marshal(data)
	reg := regexp.MustCompile(`id\":(\d{16,20}),"`)
	l := len(reg.FindAllString(string(j), -1)) //正则匹配16-20位的数字，如果找到了就开始正则替换并解析
	if l != 0 {
		var mapResult map[string]interface{}
		str := reg.ReplaceAllString(string(j), `id": "${1}","`)
		json.Unmarshal([]byte(str), &mapResult)
		data = &mapResult
	}
	return data
}

/**
 * 获取雪花算法的id
 * @author feige
 * @date 2023-11-22
 * @version 1.0
 * @desc
 */
func (web *BaseController) GetSnowWorkerId(workId int64) uint64 {
	worker, _ := utils.NewWorker(workId)
	return worker.NextNumId()
}

/**
 * 获取雪花算法的ID
 * @author feige
 * @date 2023-12-04
 * @version 1.0
 * @desc
 */
func (web *BaseController) GetSnowWorkerIdString(workId int64) string {
	worker, _ := utils.NewWorker(workId)
	return worker.NextId()
}

/**
 * 过滤html标签
 * @author feige
 * @date 2023-11-30
 * @version 1.0
 * @desc
 */
func (web *BaseController) FilterHtmlTag(str string) (string, error) {
	reader, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if nil != err {
		return "", err
	}
	text := reader.Text()
	return text, nil
}

/**
 * 字符串截取
 * @author feige
 * @date 2023-11-30
 * @version 1.0
 * @desc
 */
func (web *BaseController) SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	var substring = ""
	for i := start; i < end; i++ {
		substring += string(r[i])
	}

	return substring
}
