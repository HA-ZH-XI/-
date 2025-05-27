package vo

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserMessageVo struct {
	// 用户uuid
	Label string `json:"label"`
	// 用户名
	Mnum string `json:"mnum"`
	// 用户名
	Ckey string `json:"ckey"`
}
