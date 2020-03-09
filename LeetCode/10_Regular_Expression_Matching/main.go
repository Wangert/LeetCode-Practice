package main

import "fmt"

/**
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	fmt.Println(isMatch("aa", "a*"))

}

func isMatch(s string, p string) bool {

	//将字符串转为rune切片
	sRune := []rune(s)
	pRune := []rune(p)

	//路径矩阵
	d := make([][]bool, len(sRune) + 1)
	for i := 0; i < len(sRune) + 1; i++ {
		d[i] = make([]bool, len(pRune) + 1)
	}

	d[len(sRune)][len(pRune)] = true

	//定义'*'和'.'
	STAR := []rune("*")[0]
	POINT := []rune(".")[0]


	for i := len(sRune); i >= 0; i-- {
		for j := len(pRune) - 1; j >= 0; j-- {

			//判断 当前位置是否相等，即sRune[i]是否等于pRune[j]或者 sRune[j]是否为'.'
			flag := i < len(sRune) && (sRune[i] == pRune[j] || pRune[j] == POINT)

			if j + 1 < len(pRune) && pRune[j + 1] == STAR {

				d[i][j] = d[i][j + 2] || flag && d[i + 1][j]

			} else {

				d[i][j] = flag && d[i + 1][j + 1]

			}

		}
	}

	return d[0][0]
}

