package main

import (
	"fmt"
	"strings"
)

/**
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main() {

	//s := "LEETCODEISHIRING"
	s := "PAYPALISHIRING"
	//s := "A"

	fmt.Println("\nResult String:\n" + convert(s, 3))

}

func convert(s string, numRows int) string {

	if s == "" {
		return ""
	}

	//将字符串转换成rune切片
	runeStr := []rune(s)
	//获取字符总数量
	strLen := len(runeStr)
	//定义结果数组
	result := make([]rune, 0)
	//定义二维切片
	zigZag := make([][]rune, numRows)
	for i := 0; i < len(zigZag); i++ {
		zigZag[i] = make([]rune, 0)
	}
	//定义行列游标,j=1避免零除
	i, j := 0, 1
	//定义原字符串字符位置
	index := 0

	//存储最后一个字符的位置
	cow := 0

	for {
		//每行都填字符
		if 1 == (j % (numRows - 1)) {

			for {

				if i != numRows && index >= strLen {

					zigZag[i] = append(zigZag[i], []rune(" ")[0])

					i++
					continue

				} else if i == numRows && index >= strLen {

					cow = j - 1
					goto PRINT

				} else {

					zigZag[i] = append(zigZag[i], runeStr[index])

					i++
					index++
				}


				if i == numRows {
					i = 0
					j++
					break
				}

			}
		} else {

			for {
				//判断是否是字符位置
				if i == numRows - 2 - j%(numRows-1){

					//fmt.Println(i)
					zigZag[i] = append(zigZag[i], runeStr[index])

					i++
					index++

					continue
				}

				zigZag[i] = append(zigZag[i], []rune(" ")[0])

				i++

				if i == numRows && index == strLen  {
					cow = j - 1
					goto PRINT
				}

				if i == numRows {
					i = 0
					j++
					break
				}

			}

		}

	}

PRINT: index = 0

	fmt.Println("ZigZag:")

	for i = 0; i < numRows; i++ {

		for j = 0; j <= cow; j++ {

			if strings.Compare(" ", string(zigZag[i][j])) != 0 {

				result = append(result, zigZag[i][j])

			}

			fmt.Print(string(zigZag[i][j]))

		}

		fmt.Println()

	}

	return string(result)
}
