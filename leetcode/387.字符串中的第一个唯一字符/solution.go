package solution

import "fmt"

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.


注意事项：您可以假定该字符串只包含小写字母。
*/

func firstUniqChar(s string) int {
	bs := []byte(s)
	type counter struct {
		sub   int // 数组下标
		count int // 出现次数
	}
	m := make(map[byte]counter)

	for i, v := range bs {
		c := m[v].count
		m[v] = counter{sub: i, count: c + 1}
	}

	for i, v := range bs {
		if m[v].count == 1 {
			return i
		}
	}

	return -1
}

func firstUniqCharWinner(s string) int {
	x := make([]int, 26)
	for _, v := range s {
		fmt.Println(v)
		x[v-'a']++
	}
	for i, v := range s {
		if x[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
