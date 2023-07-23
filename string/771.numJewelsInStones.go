package main

// T(O)=100% S(O) = 45%
func numJewelsInStones(jewels string, stones string) int {

	if len(jewels) == 0 {
		return 0
	}

	cnt, maps := 0, map[rune]struct{}{}
	for _, c := range jewels {
		maps[c] = struct{}{}
	}

	for _, c := range stones {
		if _, ok := maps[c]; ok {
			cnt++
		}
	}

	return cnt

}
