package utils

import "regexp"

// 识别手机号码
func IsMobile(mobile string) bool {
	result, _ := regexp.MatchString(`^(1[3|4|5|6|8][0-9]\d{4,8})$`, mobile)
	return result
}
