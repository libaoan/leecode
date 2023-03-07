package main

var maps = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	} else {
		ans := make([]string, 0)
		if s, ok := maps[digits[0]]; ok {
			subList := letterCombinations(digits[1:])
			if len(subList) == 0 {
				for _, ss := range s {
					ans = append(ans, string((ss)))
				}

			} else {
				for _, ss := range s {
					for _, subS := range subList {
						ans = append(ans, string(ss)+subS)
					}
				}
			}
			return ans
		} else {
			panic(any("error" + digits[1:]))
		}
	}

}

func main() {
	s := "23"
	res := letterCombinations(s)
	for _, r := range res {
		println(r)
	}
}
