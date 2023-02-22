package utils

import (
	"fmt"
	"regexp"
)

//
// MatchString
//  @Description: 验证前端返回的字符串是否符合某种正则表达式标准
//  @param pattern: 正则表达式
//  @param data: 需要验证的字符串
//  @return bool: true为成功，false为失败
//
func MatchString(pattern, data string) bool {
	matched, err := regexp.MatchString(pattern, data)
	if !matched {
		return false
	}
	if err != nil {
		fmt.Println("[api MatchString err] regexp.MatchString : ", err.Error())
		return false
	}
	return true
}
