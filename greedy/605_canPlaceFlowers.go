package main

// todo 通过率 20% 速度 2N ， 内存 1
func canPlaceFlowers(flowerbed []int, n int) bool {
	fn := len(flowerbed)
	if fn <= n {
		return false
	}
	i := 0
	for flowerbed[i] != 1 {
		i++
	}
	cnt := i / 2

	for i = i + 2; i < fn-1; i += 2 {

		if flowerbed[i] == 0 {
			flowerbed[i] = 1
			cnt++
		}
	}
	if i == fn-1 && flowerbed[i] == 0 {
		cnt++
	}
	return cnt >= n
}

func main() {
	nums := []int{1, 0, 0, 0, 0, 1}
	canPlaceFlowers(nums, 2)
}
