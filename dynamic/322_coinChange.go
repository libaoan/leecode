package main

import (
	"fmt"
	"sort"
	"time"
)

// todo: 未实现 不能在单层循环内实现 贪心算法
func coinChange2(coins []int, amount int) int {

	n := len(coins)
	if n == 0 && amount == 0 {
		return 0
	} else if n == 0 {
		return -1
	} else {
		sort.Ints(coins)
		cnt, amountT := 0, amount

		for i := n - 1; i >= 0 && amountT > 0; i-- {
			d := amountT / coins[i]
			amountT = amountT % coins[i]
			cnt += d
		}

		if amountT > 0 {
			cnt = -1
		}

		res := coinChange(coins[:n-1], amount)
		fmt.Println(coins, res, cnt)
		if res > 0 && res < cnt {
			return res
		} else if res == 0 {
			return 0
		} else {
			return cnt
		}

	}

}

func coinChange(coins []int, amount int) int {

	n := len(coins) - 1
	if amount == 0 {
		return 0
	}
	if amount < 0 || n < 1 {
		return -1
	}

	cnt := 0
	cnt1 := coinChange(coins, amount-coins[n])
	cnt0 := coinChange(coins[:n], amount)
	// fmt.Println(coins, cnt1, coins[:n], cnt0)
	if cnt1 < 0 && cnt0 < 0 {
		return -1
	}
	if cnt1 < 0 {
		return cnt0
	}
	if cnt0 < 0 {
		return cnt1 + 1
	}
	if cnt1 == 0 && cnt0 == 0 {
		return 1
	}
	if cnt0 == 0 || cnt1 == 0 {
		return cnt1 + 1
	}

	if cnt0 > cnt1+1 {
		cnt = cnt0
	} else {
		cnt = cnt1 + 1
	}

	return cnt
}

func main() {
	start := time.Now().UnixNano()
	coins := []int{1, 2, 5}
	amount := 11
	coins = []int{186, 419, 83, 408}
	amount = 6249
	sort.Ints(coins)
	fmt.Println(coins)
	fmt.Println(coinChange(coins, amount))
	fmt.Println("运行时间", time.Now().UnixNano()-start)
}
