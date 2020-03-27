package main

import (
	"math"
	"fmt"
)

/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

 

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2
 

提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	fmt.Println(divide(7, -3))

}

func divide(dividend int, divisor int) int {

	if divisor == 0 {
		panic("divisor is zero")
	}

	//结果正负标记,false表示负
	flag := false

	if dividend >= 0 && divisor > 0 || dividend <= 0 && divisor < 0 {
		flag = true
	}

	//临时存储结果
	tempResult := 0
	//取绝对值
	posDividend := int(math.Abs(float64(dividend)))
	posDivisor := int(math.Abs(float64(divisor)))
	//存储结果递增值
	addNum := 1

	for {

		if posDividend < posDivisor && posDivisor != int(math.Abs(float64(divisor))) {
			posDivisor = int(math.Abs(float64(divisor)))
			addNum = 1
			continue
		}

		if posDividend < posDivisor {
			break
		}

		if posDividend >= posDivisor {
			posDividend = posDividend - posDivisor
			tempResult = tempResult + addNum
			//翻倍处理
			posDivisor = posDivisor + posDivisor
			addNum = addNum + addNum
		}

	}

	if flag == false {
		tempResult = -tempResult
	}

	if tempResult < math.MinInt32 || tempResult > math.MaxInt32 {
		return math.MaxInt32
	}

	return tempResult

}