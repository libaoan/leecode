package main

// 递归算法 超时，通过率90%
func numTrees(n int) int {

	if n == 0 {
		return 1
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += numTrees(i-1) * numTrees(n-i)
	}
	return sum
}

// 备忘录 速度 7%， 内存 98%
func numTrees2(n int) int {
	maps := make(map[int]int, n)
	return count(n, maps)
}

func count(n int, maps map[int]int) int {
	if n == 0 {
		return 1
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if _, ok := maps[i-1]; !ok {
			maps[i-1] = count(i-1, maps)
		}
		if _, ok := maps[n-i]; !ok {
			maps[n-i] = count(n-i, maps)
		}

		sum += maps[i-1] * maps[n-i]
	}
	return sum
}
