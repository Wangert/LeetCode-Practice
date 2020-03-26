package main

import "fmt"

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	haystack := "ababcaababcaabc"
	needle := "ababcaabc"

	fmt.Println(strStr(haystack, needle))

}

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	//将字符串转为rune切片
	haystackRunes := []rune(haystack)
	needleRunes := []rune(needle)

	//被匹配串和匹配串游标
	i, j := 0, 0
	//记录匹配成功时，首字母位置
	resultIndex := -1

	for {

		if j >= len(needleRunes) {

			//记录匹配成功的第一个位置
			resultIndex = i - len(needleRunes)

			break
		}

		if i >= len(haystackRunes) {
			break
		}

		//第一个字母就不匹配，移动被匹配串游标
		if haystackRunes[i] != needleRunes[j] && j == 0 {
			i++
			continue
		}

		if haystackRunes[i] != needleRunes[j] && j != 0 {
			j = getNextIndex(needleRunes[:j])
			continue
		}

		if haystackRunes[i] == needleRunes[j] {
			i++
			j++
		}

	}

	return resultIndex

}

//获取匹配串下次开始匹配的位置
func getNextIndex(matchedStr []rune) int {

	if len(matchedStr) == 1 {
		return 0
	}

	//首尾位置
	i, j := len(matchedStr) - 1, 1
	//记录最长公共前后缀的长度
	maxLength := 0

	for {

		if string(matchedStr[:i]) == string(matchedStr[j:]) {
			maxLength = len(matchedStr[:i])
			break
		}

		i--
		j++

		if i == 0 || j == len(matchedStr) {
			break
		}

	}

	return maxLength
}