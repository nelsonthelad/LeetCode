func strStr(haystack string, needle string) int {
	match := false

	for i := 0; i < (len(haystack) - len(needle) + 1); i++ {
		for j := 0; j < len(needle); j++ {
			match = true
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}