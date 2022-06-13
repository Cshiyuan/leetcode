package awesome

import (
	"fmt"
)

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

func multiply(num1 string, num2 string) string {



	sums := make([][]byte, 0, 0)

	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	if len(num1) == 1 && num1[0] == '0' {
		return "0"
	}
	if len(num2) == 1 && num2[0] == '0' {
		return "0"
	}

	for i := len(num1) - 1; i >= 0; i-- {

		var supVal uint8 = 0
		sum := make([]byte, 0, 0)

		for j := 0; j < len(num1)-1-i; j++ {
			sum = append(sum, '0')
		}

		for j := len(num2) - 1; j >= 0; j-- {

			val1 := num1[i] - '0'
			val2 := num2[j] - '0'
			baseVal := (val1 * val2) + supVal
			if baseVal >= 10 {
				supVal = baseVal / 10
				baseVal = baseVal % 10
			} else {
				supVal = 0
			}
			sum = append(sum, byte(baseVal)+'0')
		}

		if supVal != 0 {
			sum = append(sum, byte(supVal)+'0')
			supVal = 0
		}
		sums = append(sums, sum)

		//fmt.Printf("sum append: %v \n", string(sum))
	}

	// 数字相加
	var result []byte

	for _, num := range sums {
		str := string(num)
		fmt.Printf("result: %v, num: %v \n", string(result), str)
		result = addNums(result, num)
		str = string(result)
		fmt.Printf("result: %v \n", str)
	}

	result = reverseB(result)
	return string(result)
}

func addNums(num1, num2 []byte) []byte {

	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}
	supVal := 0
	var nums []byte
	for i, j := 0, 0; i < len(num1) || j < len(num2) || supVal != 0; i, j = i+1, j+1 {

		var val1, val2 byte = 0, 0
		if i < len(num1) {
			val1 = num1[i] - '0'
		}
		if j < len(num2) {
			val2 = num2[j] - '0'
		}
		val := val1 + val2 + byte(supVal)
		if val >= 10 {
			supVal = int(val / 10)
			val = val % 10
		} else {
			supVal = 0
		}
		nums = append(nums, val+'0')
		str := string(nums)
		fmt.Printf("nums: %v \n",str)
	}
	return nums
}

func reverseB(num []byte) []byte {

	var result []byte
	for i := len(num) - 1; i >= 0; i-- {
		result = append(result, num[i])
	}
	return result
}
