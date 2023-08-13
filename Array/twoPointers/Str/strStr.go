package Str

func strStr(haystack string, needle string) int {
	// 如果 needle 为空字符串，则返回 0。
	if len(needle) < 1 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			j := 1
			for j < len(needle) && haystack[i+j] == needle[j] {
				j++
			}
			if j == len(needle) {
				return i
			}
		}
	}

	// 如果没有找到 needle，则返回 -1。
	return -1
}
