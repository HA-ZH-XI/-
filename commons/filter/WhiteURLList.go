package filter

import (
	"ksd-social-api/commons/global"
)

/*
*
  - @author feige
  - @date 2023-10-26
  - @version 1.0
  - @desc 白名单

进程-----内存控制（堆、（a588bd9d01ecc75dbbaaefce）栈）
*/

func init() {

	// 因为这些接口地址都是微信服务器，支付宝服务，微信服务器来调用的
	global.List.PushBack("/api/pay/weixin/native/success/callback")
	global.List.PushBack("/api/pay/alipay/native/success/return")
	global.List.PushBack("/api/pay/alipay/native/success/notify")
	global.List.PushBack("/api/login/weixin/login")
	global.List.PushBack("/api/login/weixin/callback")
}

/**
 * @author feige
 * @date 2023-10-26
 * @version 1.0
 * @desc 判断url是否是白名单
 */

func ValidateURL(url string) bool {
	flag := false
	//通过 Front() 获取链表头节点，循环遍历直到链表末尾（nil）
	for i := global.List.Front(); i != nil; i = i.Next() {
		if i.Value == url {
			flag = true
			break
		}
	}
	return flag
}
