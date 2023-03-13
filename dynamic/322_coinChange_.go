package main

import (
	"fmt"
	"math"

	// "sort"
	"time"
)

// 递归回朔算法 结果正确，但是运行超时 用例通过率30%
func coinChange(coins []int, amount int) int {

	return doCoinChange(coins, amount, 0)
}

func doCoinChange(coins []int, amount int, cnt int) int {
	n := len(coins) - 1
	if amount == 0 {
		return cnt
	}
	if amount < 0 || n < 0 {
		return -1
	}

	cnt0, cnt1 := -1, -1
	if amount-coins[n] >= 0 {
		cnt1 = doCoinChange(coins, amount-coins[n], cnt+1)
	}
	if n > 0 {
		cnt0 = doCoinChange(coins[:n], amount, cnt)
	}

	if cnt0 < 0 && cnt1 < 0 {
		return -1
	}

	if cnt1 < 0 {
		return cnt0
	}

	if cnt0 < 0 {
		return cnt1
	}

	if cnt0 < cnt1 {
		cnt = cnt0
		fmt.Println(coins, coins[:n], amount, cnt)
	} else {
		cnt = cnt1
		fmt.Println(coins, coins[:n], amount, cnt)
	}

	return cnt
}

/*
确定dp数组以及下标的含义
dp[j]：凑足总额为j所需钱币的最少个数为dp[j]

确定递推公式
凑足总额为j - coins[i]的最少个数为dp[j - coins[i]]，那么只需要加上一个钱币coins[i]即dp[j - coins[i]] + 1就是dp[j]（考虑coins[i]）

所以dp[j] 要取所有 dp[j - coins[i]] + 1 中最小的。

递推公式：dp[j] = min(dp[j - coins[i]] + 1, dp[j]);

dp数组如何初始化
首先凑足总金额为0所需钱币的个数一定是0，那么dp[0] = 0;

作者：代码随想录
链接：https://leetcode.cn/problems/coin-change/solutions/975451/dai-ma-sui-xiang-lu-dai-ni-xue-tou-wan-q-80r7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
// 版本一, 先遍历物品,再遍历背包
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp,j,i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 版本二,先遍历背包,再遍历物品
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 遍历背包,从1开始
	for j := 1; j <= amount; j++ {
		// 初始化为math.MaxInt32
		dp[j] = math.MaxInt32
		// 遍历物品
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] && dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	start := time.Now().UnixNano()
	coins := []int{1, 2, 5}
	amount := 11
	coins = []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}
	amount = 9864
	fmt.Println(coinChange1(coins, amount))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
	start = time.Now().UnixNano()
	fmt.Println(coinChange2(coins, amount))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
