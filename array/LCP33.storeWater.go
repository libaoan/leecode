package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	bkt, vat := []int{1, 3}, []int{6, 8}
	bkt, vat = []int{9, 0, 1}, []int{0, 2, 2}
	bkt = []int{
		3710, 6067, 2993, 70, 2340, 2748, 9385, 3027, 3456, 5246, 9739, 1220, 9539, 9074, 4729, 7051, 8462, 6908, 3649, 9996, 8890, 2980, 4350, 7350, 6344, 6759, 4420, 269, 9341, 648, 7737, 8133, 3717, 2766, 5807, 4338, 2077, 5775, 4905, 7262, 1258, 613, 3837, 3475, 437, 3739, 9814, 4790, 2075, 7722, 3290, 5685, 3499, 6992, 4421, 934, 6004, 5763, 3463, 6138, 8818, 445, 4778, 4979, 126, 3969, 2994, 87, 3739, 8582, 9559, 8326, 9132, 257, 8928, 9147, 1615, 4665, 9828, 3925, 6435, 5326, 836, 519, 298, 600, 5503, 273, 9580, 5383, 8966, 4810, 1386, 7207, 8060, 678, 8837, 6946, 1210, 945,
	}
	vat = []int{
		6304, 6509, 4276, 9645, 6455, 8167, 9667, 4385, 8872, 7889, 9936, 4413, 9922, 9894, 8065, 7627, 9225, 9907, 7055, 9996, 9439, 3351, 9317, 8363, 9383, 6850, 4621, 5389, 9508, 3391, 9650, 8363, 8719, 5594, 8770, 5403, 7107, 9941, 9254, 9355, 4614, 4640, 4896, 8759, 4397, 8441, 9870, 9906, 2396, 8092, 6939, 9432, 8182, 9090, 8029, 4930, 7772, 7066, 7279, 7778, 9529, 2947, 6552, 6930, 5260, 8470, 8478, 1371, 9453, 9767, 9888, 9964, 9960, 3990, 9391, 9377, 3063, 5374, 9880, 7684, 7864, 7078, 2622, 3754, 617, 9773, 9415, 8026, 9883, 5735, 9233, 6715, 9105, 7932, 9178, 1081, 9340, 7284, 6621, 1965,
	}

	bkt = []int{
		9988, 5017, 5130, 2445, 9896, 9151, 3625, 7801, 608, 3283, 1386, 979, 5209, 4182, 8234, 9870, 8714, 6435, 3800, 956, 4006, 5620, 7474, 1205, 6993, 3320, 1201, 7593, 905, 3816, 4522, 4560, 8027, 8219, 6686, 3779, 2141, 1240, 6504, 6612, 6921, 7329, 8145, 5745, 7652, 4340, 7933, 6246, 5157, 9447, 107, 9665, 3653, 2978, 9832, 4945, 4312, 2199, 449, 8432, 3230, 8163, 800, 6547, 1110, 1194, 9384, 632, 3275, 1229, 7230, 8643, 7613, 8256, 5043, 1288, 3088, 8997, 4554, 4755, 7433, 8146, 9722, 3469, 8863, 5831, 7816, 5058, 4316, 7946, 8402, 975, 2450, 4958, 9811, 9336, 21, 9309, 8999, 56,
	}
	vat = []int{
		9991, 6973, 7192, 9876, 9910, 9549, 3700, 8814, 1308, 9981, 9234, 7292, 7732, 8458, 8441, 9939, 9621, 7285, 7452, 2718, 6589, 7555, 8788, 3202, 7832, 4781, 8798, 9299, 2112, 9963, 8755, 7240, 9217, 8587, 6782, 9703, 8954, 3759, 6907, 7218, 7333, 8020, 8323, 5750, 9510, 8571, 8664, 8510, 9363, 9741, 8643, 9825, 4227, 8530, 9961, 8511, 8949, 7486, 9086, 9690, 5316, 9581, 9314, 8817, 7234, 8998, 9485, 5394, 7365, 1501, 7984, 9802, 9778, 8314, 7482, 7117, 5117, 9609, 8732, 9728, 9330, 8800, 9775, 6210, 8966, 7700, 8802, 7607, 8950, 9730, 9855, 1231, 5228, 5329, 9982, 9532, 3230, 9951, 9034, 8299,
	}

	bkt = []int{
		6851, 2583, 9351, 5443, 496, 2050, 8923, 6176, 6756, 418, 2522, 7241, 7492, 5014, 34, 6664, 1298, 4205, 4196, 1908, 1649, 1779, 8906, 4359, 934, 1932, 4146, 2185, 4306, 9244, 9485, 357, 3065, 4448, 1430, 8311, 9045, 2600, 7336, 8778, 8883, 4019, 152, 9208, 3665, 3852, 5121, 756, 9407, 7310, 3851, 192, 3052, 3958, 4715, 9012, 814, 8474, 858, 7187, 5402, 7705, 7465, 1802, 3376, 8813, 3015, 7915, 1359, 8856, 6551, 4763, 8602, 2276, 4613, 4255, 8058, 1471, 8620, 8323, 8774, 4667, 1172, 298, 9843, 1025, 1876, 7519, 6490, 9829, 1061, 5183, 5200, 5818, 6424, 495, 2439, 2674, 2082, 7369,
	}
	vat = []int{
		7806, 8779, 9796, 7344, 5216, 3759, 9630, 7758, 9135, 7443, 8698, 7450, 8851, 7268, 2168, 9096, 6284, 4715, 5915, 3447, 6346, 4909, 9135, 8686, 4325, 9703, 6434, 3207, 5511, 9575, 9921, 9881, 3207, 8251, 2460, 8982, 9598, 9897, 9620, 9275, 8965, 6337, 8343, 9709, 8917, 5853, 7158, 9835, 9852, 7914, 9710, 4343, 6290, 6960, 9549, 9652, 3626, 9780, 1243, 9768, 8022, 8830, 9169, 5544, 8613, 9875, 8361, 8667, 3165, 9012, 8888, 8366, 9602, 7676, 7961, 4579, 8874, 2566, 8772, 8845, 9888, 6790, 6622, 8475, 9908, 1704, 8012, 9949, 8186, 9908, 1828, 5461, 7352, 8804, 9744, 8724, 8180, 6926, 3001, 9736,
	}

	fmt.Println(storeWater(bkt, vat))
	fmt.Println(storeWater2(bkt, vat))
}

// todo: 通过率 82%
func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	sli := make([][3]int, n)

	ck := 0
	for i := 0; i < n; i++ {
		if bucket[i] == 0 && vat[i] != 0 {
			bucket[i] = 1
			ck++
		}
	}

	for i := 0; i < n; i++ {
		sli[i][0], sli[i][1] = bucket[i], vat[i]
		if vat[i] == 0 {
			sli[i][2] = 0
		} else if vat[i]%sli[i][0] == 0 {
			sli[i][2] = vat[i]/sli[i][0] + ck
		} else {
			sli[i][2] = 1 + vat[i]/sli[i][0] + ck
		}
	}

	for {
		sort.Slice(sli, func(i, j int) bool {
			return sli[i][2] > sli[j][2] || (sli[i][2] == sli[j][2] && sli[i][0] > sli[j][0])
		})

		fmt.Println(sli[0], sli[1], ck)

		mCnt, cnt := sli[0][2], 0
		cap := sli[0][0] + 1
		ck++
		if sli[0][1]%cap == 0 {
			cnt = sli[0][1]/cap + ck
		} else {
			cnt = sli[0][1]/cap + 1 + ck
		}
		if cnt <= mCnt && cap <= sli[0][1] {
			sli[0][2] = cnt
			sli[0][0] = cap
		} else {
			return mCnt
		}
	}

}

// todo 参考
func storeWater2(bucket []int, vat []int) int {
	n := len(bucket)
	maxk := 0
	for _, v := range vat {
		if v > maxk {
			maxk = v
		}
	}
	if maxk == 0 {
		return 0
	}
	res := math.MaxInt32
	for k := 1; k <= maxk && k < res; k++ {
		t := 0
		for i := 0; i < n; i++ {
			t += max(0, (vat[i]+k-1)/k-bucket[i])
		}
		res = min(res, t+k)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}