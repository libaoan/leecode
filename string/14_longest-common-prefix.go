package main

func main() {
	strs := []string{"flower", "flow", "flight"}
	strs = []string{"dog", "racecar", "car"}
	strs = []string{"a"}
	println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	prefix := ""
	i := 0

	for {
		if i > len(strs[0]) {
			return prefix
		}

		prefix = strs[0][0:i]
		for _, str := range strs {
			if i > len(str) || prefix != str[0:i] {
				return prefix[0 : i-1]
			}
		}
		i++
	}

}
