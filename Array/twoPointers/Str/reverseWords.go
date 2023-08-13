package Str

import "strings"

func reverseWords(s string) string {
	// 删除首尾空格
	s = strings.TrimSpace(s)

	var res []string
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		//搜索首个空格
		for i >= 0 && s[i] != ' ' {
			i--
		}
		//添加单词
		res = append(res, s[i+1:j+1])
		//跳过单词间空格
		for i >= 0 && s[i] == ' ' {
			i--
		}
		//j 指向下个单词的尾字符
		j = i
	}

	return strings.Join(res, " ")
}
