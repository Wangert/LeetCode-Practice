package main

import (
	"fmt"
	"math"
)

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

进阶:

你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main() {

	fmt.Println(isPalindrome(123321))
	fmt.Println(numberOfDigits(1234321))

}

//不用字符串判断
func isPalindrome(x int) bool {
	//小于零为负数，肯定不是回文数
	if x < 0 {
		return false
	}

	//高低位游标, 高低位数字
	var low, high, lowNum, highNum int
	low = 1
	high = numberOfDigits(x)

	for {

		lowNum = x % int(math.Pow10(low)) / int(math.Pow10(low - 1))
		highNum = x % int(math.Pow10(high)) / int(math.Pow10(high - 1))

		fmt.Println(lowNum)
		fmt.Println(highNum)

		low++
		high--

		if (lowNum == highNum) {

			if (low >= high) {
				return true
			}

			continue

		}else {

			return false
		}

	}

}

//判断整数是几位数
func numberOfDigits(x int) int {

	//计数
	count := 1

	for {
		x = x / 10
		if x <= 0 {
			break
		}

		count++
	}

	return count
}

