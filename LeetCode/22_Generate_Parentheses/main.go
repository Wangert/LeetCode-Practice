package main

import "fmt"

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */

//存储结果
type Result struct {
	rs []string
}

func main()  {

	fmt.Println(generateParenthesis(4))

}

func generateParenthesis(n int) []string {

	if n == 0 {
		return []string{}
	}

	//定义结果集
	result := &Result{}

	result.depthF("", 0, 0, n)

	return result.rs

}

//深度优先遍历
func (result *Result) depthF(s string, numOfLeft, numOfRight, n int) {

	if numOfLeft > n {
		return
	}

	if numOfLeft < numOfRight {
		return
	}

	if numOfLeft + numOfRight == n * 2{
		result.rs = append(result.rs, s)
		return
	}

	result.depthF(s + "(", numOfLeft + 1, numOfRight, n)
	result.depthF(s + ")", numOfLeft, numOfRight + 1, n)

}
