package main

import "fmt"

/*
规律法：
1. 第0行的字符位于原串索引k * (2*numRows - 2)处；
2. 第numRows-1行的字符位于原串索引k * (2*numRows - 2) + numRows - 1处；
3. 中间行的字符位于原串索引k * (2*numRows - 2) + i处和(k + 1) * (2*numRows - 2) - i处
 */

func main()  {
	s := "PAYPALISHIRING"

	fmt.Println(convert(s, 3))
}


func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	//记录结果
	result := make([]rune, 0)
	//将字符串转换成rune切片
	runeStr := []rune(s)
	//计算总字符数
	strLen := len(runeStr)
	//用于循环的因子数
	cycleNumer := 2 * numRows -2

	for i := 0; i < numRows; i++ {
		for j := 0; j + i < strLen; j = j + cycleNumer {

			result = append(result, runeStr[j + i])

			if (i != 0 && i != numRows - 1 && j + cycleNumer - i < strLen) {

				result = append(result, runeStr[j + cycleNumer - i])

			}


		}
	}

	return string(result)

}