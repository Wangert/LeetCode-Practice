package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	fmt.Println(reverse(-21474))

}

//整数反转
func reverse(x int) int {

	const MAX = math.MaxInt32
	const MIN = math.MinInt32

	//是否为负数
	flag := false

	//判断是否为负数
	if x < 0 {
		flag = true
	}

	xAbs := int(math.Abs(float64(x)))
	//获取整数字符串
	xStr := strconv.Itoa(xAbs)

	reverseXStr := reverseString(xStr)

	var reverseInt int

	if flag == false {

		reverseInt, _ = strconv.Atoi(reverseXStr)

	}else {

		tmp, _ := strconv.Atoi(reverseXStr)
		reverseInt = -tmp

	}

	if reverseInt > MAX || reverseInt < MIN{
		return 0
	}

	return reverseInt
}


// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}