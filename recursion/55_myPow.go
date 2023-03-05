package main

// 递归, 通过率 90%, 算法复杂度n
func myPow(x float64, n int) float64 {

	var res float64 = 0
	switch {
	case n == 0:
		{
			res = 1
		}
	case n == 1:
		{
			res = x
		}
	case n == -1:
		{
			res = 1 / x
		}
	case x == 1:
		{
			res = 1
		}
	case x == -1:
		{
			return -1
		}
	case n < 0:
		{
			res = 1 / x * myPow(x, n+1)
		}
	case n > 1:
		{
			res = x * myPow(x, n-1)
		}
	}

	return res
}

// 非递归 通过率90%, 时间复杂度n
func myPow2(x float64, n int) float64 {
	var res float64 = 1
	switch {
	case n == 0:
		{
			res = 1
		}
	case x == 1:
		{
			res = 1
		}
	case x == -1:
		{
			res = -1
		}
	case n >= 1:
		{
			for n >= 1 {
				res *= x
				n--
			}
			return res
		}
	case n <= -1:
		{
			for n <= -1 {
				res *= 1 / x
				n++
			}
			return res
		}
	default:
		panic(any("not a nu,ber"))
	}
	return res
}

// todo：别人的递归答案， 算法复杂度logn
func myPow3(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {

	println(myPow3(0.00001, 2147483647))
}
