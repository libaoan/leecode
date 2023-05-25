package main

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
