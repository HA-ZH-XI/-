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

// 通过链表存储和验证白名单 appid，确保只有合法请求能通过校验
func init() {
	/*在包被加载时自动执行，向全局链表 global.List 中添加一个白名单条目 "a588bd9d01ecc75dbbaaefce"。
	PushBack 是链表操作，将值添加到链表末尾。*/
	// a588bd9d01ecc75dbbaaefce是ksd-social-web使用
	global.List.PushBack("a588bd9d01ecc75dbbaaefce")
}

/**
 * @author feige
 * @date 2023-10-26
 * @version 1.0
 * @desc 判断appid是否是白名单
 */
//// 遍历全局链表 global.List，检查传入的 appid 是否存在于白名单中
func ValidateAppid(appid string) bool {
	flag := false
	for i := global.List.Front(); i != nil; i = i.Next() {
		if i.Value == appid {
			flag = true
			break
		}
	}
	return flag
}

/*若找到匹配的 appid，设置 flag 为 true 并立即终止循环。
  返回验证结果 flag。*/

/*优化点*/

//1.当前使用链表（global.List）存储白名单，适合少量数据。若白名单条目较多，建议改用 Map（如 map[string]struct{}），将查询时间复杂度从 O(n) 降至 O(1)。
//2.白名单 appid 直接硬编码在代码中，不利于维护。建议通过配置文件或环境变量动态加载
//3.若 global.List 会被并发访问，需加锁保护（如 sync.Mutex），避免数据竞争
//4.注释问题  白名单的具体用途（如限制接口访问权限）。 global.List 的预期结构和初始化方式。
