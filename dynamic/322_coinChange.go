package main

import (
	"fmt"
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

func main() {
	start := time.Now().UnixNano()
	coins := []int{1, 2, 5}
	amount := 11
	coins = []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}
	amount = 9864
	//sort.Ints(coins)
	fmt.Println(coinChange(coins, amount))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
