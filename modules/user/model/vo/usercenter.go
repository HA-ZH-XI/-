package vo

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserStateCountVo struct {
	Bbscount    int `json:"bbscount"`
	Coursecount int `json:"coursecount"`
	Fanscount   int `json:"fanscount"`
	Gzcount     int `json:"gzcount"`
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserStateCountAllVo struct {
	Snum   int    `json:"snum"`
	Sname  string `json:"sname"`
	Sfield string `json:"sfield"`
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserStateCountAllChildVo struct {
	Snum int `json:"snum"`
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  返回用户信息
 */
type UserStateModelDataVo struct {
	Name     string                        `json:"name"`
	Model    string                        `json:"model"`
	SystemId int                           `json:"systemId"`
	Value    []*UserStateModelDataChildren `json:"value"`
}

type UserStateModelDataChildren struct {
	Datestr  string `json:"datestr"`
	Snum     int    `json:"snum"`
	SystemId int    `json:"systemId"`
}
