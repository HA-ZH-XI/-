package utils

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  星座
 */
func GetZodiac(month, day int) string {
	month = month - 1
	var (
		DAY_ARR = [12]int{20, 19, 21, 20, 21, 22, 23, 23, 23, 24, 23, 22}
		ZODIACS = [13]string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	)

	if day < DAY_ARR[month] {
		return ZODIACS[month]
	} else {
		return ZODIACS[month+1]
	}
}

/**
 * @author feige
 * @date 2023-10-10
 * @version 1.0
 * @desc  生肖
 */
func GetChineseZodiac(year int) string {
	var CHINESE_ZODIACS = [12]string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	if year > 1900 {
		return CHINESE_ZODIACS[(year-1900)%len(CHINESE_ZODIACS)]
	} else {
		return ""
	}
}
