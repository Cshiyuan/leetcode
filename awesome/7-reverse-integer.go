package awesome

import "math"

func reverseInteger(x int) (rev int) {


	// 	难点在于判断
	// 需要判断反转后的数字是否超过 32 位有符号整数的范围 [-2^{31},2^{31}-1][−2

	for x != 0 {
		// 「推入数字之前检查」
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x = x/10
		rev = rev*10 + digit
	}
	return
}
