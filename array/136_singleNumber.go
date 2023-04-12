package main

import "sort"

// 通过， 速度 12% 内存62%
func singleNumber(nums []int) int {
	sort.Ints(nums)
	i := 0
	for ; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[i]
}

// hashmap 速度 2N 25%, 内存 N 18%
func singleNumber2(nums []int) int {
	n := len(nums)
	maps := make(map[int]bool, n/2+1)
	for _, i := range nums {
		if _, ok := maps[i]; !ok {
			maps[i] = false
		} else {
			maps[i] = true
		}
	}

	for k, v := range maps {
		if !v {
			return k
		}
	}
	return -1
}

// hashmap优化 速度 N 66%, 内存 N 18%
func singleNumber3(nums []int) int {
	n := len(nums)
	maps := make(map[int]struct{}, n/2+1)
	for _, i := range nums {
		if _, ok := maps[i]; !ok {
			maps[i] = struct{}{}
		} else {
			delete(maps, i)
		}
	}

	for k, _ := range maps {
		return k
	}
	return -1
}

// 参考，别人的算法，通过位运算实现 a ^ 0 = 0 and a ^ a = 0
// 看来位运算还是一个常考的点：
// 1、 二进制一样的位为0， 不一样的位为1
// 2、相同的数字异或为0，0与任何数字异或结果是它本身
// 3、异或运算的法则类似乘法：满足交换律、结合律。a ^ b = b ^ a  (a ^ b) ^ c = a ^ (b ^ c)
// 数学算法， 速度 N 96%， 内存 1 98%
func singleNumber4(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {

}
