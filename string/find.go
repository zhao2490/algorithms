package string

/*
BF 算法（Brute Force 暴力匹配）
n : 主串长度
m : 模式串长度
时间复杂度: O(m*n)
return: 未找到返回-1 其它情况返回首字母下标
*/
func BruteForceSearch(main string, pattern string) int {
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return -1
	}

	for i := 0; i <= len(main)-len(pattern); i++ {
		subStr := main[i : i+len(pattern)]
		if subStr == pattern {
			return i
		}
	}

	return -1
}

/*
RK 算法（Rabin-Karp）
BF算法升级版，通过hash算法对主串中的n-m+1个子串分别求hash值，然后逐个与模式串的hash值比较大小

return: 未找到返回-1 其它情况返回首字母下标
*/
func simpleHash(s []rune, start int, end int) int32 {
	var ret int32
	for i := start; i < end+1; i++ {
		ret += s[i]
	}
	return ret
}

func RabinKarpSearch(main string, pattern string) int {
	runeM := []rune(main)
	runeP := []rune(pattern)
	n := len(runeM)
	m := len(runeP)
	if n <= m {
		if main == pattern {
			return 0
		}
		return -1
	}

	h := make([]int32, n-m+1)
	h[0] = simpleHash(runeM, 0, m-1)
	for i := 1; i < n-m+1; i++ {
		h[i] = h[i-1] - runeM[i-1] + runeM[i+m-1]
	}

	hashP := simpleHash(runeP, 0, m-1)

	for index, hashM := range h {
		if hashP == hashM {
			if pattern == string(runeM[index:index+m]) {
				return index
			}
		}
	}
	return -1
}
