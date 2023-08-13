package Str

func strStr(haystack string, needle string) int {
	// 如果 needle 为空字符串，则返回 0。
	if len(needle) < 1 {
		return 0
	}

	// 从 haystack 的开头开始遍历。
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		// 如果 haystack 的第 i 个字符与 needle 的第一个字符相等，则从第 i + 1 个字符开始比较。
		if haystack[i] == needle[0] {
			j := 1
			for j < len(needle) && haystack[i+j] == needle[j] {
				j++
			}
			// 如果比较了 needle 的全部字符，则返回 i。
			if len(needle) == j {
				return i
			}
		}
	}

	// 如果没有找到 needle，则返回 -1。
	return -1
}
