package main

import "fmt"

func main() {
	arr := []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	arr = []int{1, -1, 1, -1}
	fmt.Println(canThreePartsEqualSum3(arr))
}

// O(n3) 速度 96%， 内存 64%
func canThreePartsEqualSum(arr []int) bool {
	sum, midIndexs, n := 0, []int{}, len(arr)
	if n < 3 {
		return false
	}
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	if sum%3 != 0 {
		return false
	}
	sum, ssum := sum/3, 0
	for i := 0; i < n-2; i++ {
		ssum += arr[i]
		if ssum == sum {
			midIndexs = append(midIndexs, i)
		}
	}
	for _, i := range midIndexs {
		ssum = 0
		for j := i + 1; j < n-1; j++ {
			ssum += arr[j]
			if ssum == sum {
				ksum := 0
				for k := j + 1; k < n; k++ {
					ksum += arr[k]
				}
				if ksum == sum {
					return true
				}
			}
		}
	}
	return false
}

// todo: 别人的思路, 贪心算法 O(n) 速度 96%， 内存 64%
func canThreePartsEqualSum2(arr []int) bool {
	n, target := len(arr), 0
	for i := 0; i < n; i++ {
		target += arr[i]
	}
	if target%3 != 0 {
		return false
	}

	target /= 3
	cur, i := 0, 0
	for ; i < n-2; i++ {
		cur += arr[i]
		if cur == target {
			break
		}
	}
	if cur != target {
		return false
	}

	for ; i < n-2; i++ {
		cur += arr[i+1]
		if cur == 2*target {
			return true
		}
	}
	return false

}

// todo: 别人的思路, 双指针 O(n) 速度 96%， 内存 43%
func canThreePartsEqualSum3(arr []int) bool {
	n, target := len(arr), 0
	for i := 0; i < n; i++ {
		target += arr[i]
	}
	if target%3 != 0 {
		return false
	}

	target /= 3
	isum, jsum, i, j := arr[0], arr[n-1], 1, n-2

	for i < j {
		if isum != target {
			isum += arr[i]
			i++
		}

		if jsum != target {
			jsum += arr[j]
			j--
		}

		if i <= j && isum == target && jsum == target {
			return true
		}

	}
	return false

}
