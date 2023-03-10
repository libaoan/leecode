package main

// 动态规划  算法复杂度n， 空间复杂度 n
// 推倒公式 f(n) = f(n-1) + f(n-1)
func climbStairs(n int) int {

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	cnts := make([]int, n+1)
	cnts[1] = 1
	cnts[2] = 2

	for i := 3; i <= n; i++ {
		cnts[i] = cnts[i-1] + cnts[i-2]
	}
	return cnts[n]
}

//func main() {
//	fmt.Println(climbStairs(6))
//}
