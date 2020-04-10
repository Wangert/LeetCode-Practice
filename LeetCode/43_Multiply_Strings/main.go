package main

import "fmt"

/**
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	fmt.Println(multiply("123", "123"))

}

func multiply(num1 string, num2 string) string {

	//有一个为0即结果为0
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1Len, num2Len := len(num1), len(num2)
	result := make([]byte, num1Len + num2Len)

	for i := num1Len; i > 0; i-- {
		//计算当前乘积存储的位置
		index := num2Len + i - 1
		//获取当前num1的值
		num := num1[i - 1] - '0'

		for j := num2Len; j > 0; j-- {
			//两数乘积结果
			temp := result[index] + num * (num2[j - 1] - '0')
			//未进位值
			result[index] = temp % 10
			//前一位需要有进位
			index--
			result[index] += temp / 10
		}
	}

	notZeroIndex:= -1
	for i := 0; i < num1Len + num2Len; i++ {

		if result[i] != 0 && notZeroIndex == -1 {
			notZeroIndex = i
		}

		result[i] += '0'
	}

	return string(result[notZeroIndex:])
}