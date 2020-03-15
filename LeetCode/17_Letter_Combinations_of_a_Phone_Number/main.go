package main

import (
	"strconv"
	"fmt"
)

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与九宫格电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */

func main()  {

	digits := "23"

	fmt.Println(letterCombinations(digits))

}

//结果结构体
type Result struct {
	strs []string
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	//定义映射关系
	digitToLetter := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	//存储结果
	result := &Result{}
	//将字符串转为rune切片
	digitRunes := []rune(digits)

	result.backTrack("", digitRunes, digitToLetter, 0)

	return result.strs
}

//用于回溯
func (result *Result) backTrack(tempStr string, digitRunes []rune, digitToLetter [][]string, nextIndex int) {
	//超过数组结束
	if nextIndex >= len(digitRunes) {
		result.strs = append(result.strs, tempStr)
		return
	} else {

		digit, _ := strconv.Atoi(string(digitRunes[nextIndex]))
		//遍历当前数字对应的字母
		for _, letter := range digitToLetter[digit - 2] {

			result.backTrack(tempStr + letter, digitRunes, digitToLetter, nextIndex + 1)

		}
	}
}