package solution

import (
	"reflect"
	"strings"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/
func isPalindrome(s string) bool {
	ls := strings.ToLower(s) // step1: 转换为小写
	is := make([]int32, 0)
	temp := make([]int32, 0)

	for _, v := range ls { // step2: 过滤字符，仅保留数字和字母
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') {
			is = append(is, v)
		}
	}

	temp = append(temp, is[:]...) // step3: 拷贝一个新串

	reverse(temp) // step4: 反转新串

	if reflect.DeepEqual(is, temp) { // step5: 比较新串和旧串
		return true
	}

	return false

}

func reverse(is []int32) {
	for i, j := 0, len(is)-1; j > i; {
		is[i], is[j] = is[j], is[i]
		i++
		j--
	}
}

func isPalindromeWinner(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = filter(s)

	for i, j := 0, len(s)-1; j > i; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// 将特殊字符和空格过滤，大写转换为小写，保留字母和数字
func filter(s string) string {
	bytes := make([]byte, len(s))
	index := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c = c - 'A' + 'a' // 巧妙的大小写转换
		}

		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			bytes[index] = c
			index++
		}
	}

	return string(bytes[:index])
}
