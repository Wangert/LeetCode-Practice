package main

import "fmt"

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	strs := []string{"ggg", "ggr", "ggggrer"}

	fmt.Println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	//定义rune二维切片
	runeArrays := make([][]rune, len(strs))

	for i := 0; i < len(strs); i++ {

		if strs[i] == "" {
			return ""
		}

		runeArrays[i] = []rune(strs[i] + " ")
	}

	//定义游标
	index := 0
	//定义空格
	blank := []rune(" ")[0]

	for {

		for i := 0; i < len(strs) - 1; i++ {
			//表示至少有两个串该位置字符不同
			if runeArrays[i][index] == blank || runeArrays[i+1][index] == blank || runeArrays[i][index] != runeArrays[i+1][index] {
				goto END
			}
		}

		index++

	}

END:
	if index == 0 {
		return ""
	}

	return string(runeArrays[0][:index])
}
