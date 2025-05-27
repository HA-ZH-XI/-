package context

/**
 * @desc 用户充值
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserWalletContext struct {
	PageNo   int64  `json:"pageNo"`   // 第几页
	PageSize int64  `json:"pageSize"` // 每页显示多少条
	UserId   uint64 `json:"userId"`   // 用户ID
	SystemId int    `json:"systemId"`
}

/**
 * @desc 用户身份升级
 * @author feige
 * @date 2023-11-14
 * @version 1.0
 */
type UserVipContext struct {
	PageNo   int64  `json:"pageNo"`   // 第几页
	PageSize int64  `json:"pageSize"` // 每页显示多少条
	UserId   uint64 `json:"userId"`   // 用户ID
	SystemId int    `json:"systemId"`
}
