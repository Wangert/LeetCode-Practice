package main

import (
	"strings"
	"fmt"
)

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main() {

	s := ""

	fmt.Println(longestPalindrome(s))


}

func longestPalindrome(s string) string {

	//将字符串转换成rune切片
	runeStr := []rune(s)
	//获取字符串长度
	strLen := len(runeStr)
	//存储最长子串
	longestStr := ""
	//记录最大长度
	maxLen := 0

	for index, _ := range runeStr {

		if maxLen >= strLen - index {
			break
		}

		for i := strLen; i >= 0; i-- {

			if maxLen >= i - index {
				break
			}

			//原始子串
			oldSubStr := string(runeStr[index:i])
			//反转子串
			reverseSubStr := reverseString(oldSubStr)

			//判断是否回文
			if strings.Compare(oldSubStr, reverseSubStr) == 0 {
				//记录当前最长回文子串
				longestStr = oldSubStr
				//记录当前最大长度
				maxLen = i - index

				break
			}

		}

	}

	return longestStr

}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
