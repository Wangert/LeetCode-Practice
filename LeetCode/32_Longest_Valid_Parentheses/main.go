package main

import (
	"LeetCode/32_Longest_Valid_Parentheses/stack"
	"reflect"
	"fmt"
)

/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main() {

	s := "(()"
	fmt.Println(longestValidParentheses(s))

}

func longestValidParentheses(s string) int {

	//定义左右括号
	par := []rune("()")
	//将字符串转为rune切片
	sRunes := []rune(s)
	//定义一个栈
	myStack := stack.NewStack(reflect.TypeOf(int(0)))
	//记录最大长度
	maxLen := 0

	for i := 0; i < len(sRunes); i++ {
		//此时栈里没有左括号，即当前右括号无效
		if sRunes[i] == par[1] {
			//为空直接将右括号入栈，不为空，判断是否有左括号匹配
			if myStack.Empty() {
				myStack.Push(i)
			} else {
				index, _ := myStack.Pop()
				if sRunes[index.(int)] != par[0] {
					myStack.Push(index.(int))
					myStack.Push(i)
				}
			}
			continue
		}

		//左括号入栈
		if sRunes[i] == par[0] {
			myStack.Push(i)
		}
	}

	if myStack.Empty() {
		return len(sRunes)
	}

	rearIndex := len(sRunes)

	//将栈中的位置pop出，求相邻距离的最大，即为最长有效括号的长度
	for {

		frontIndex, _ := myStack.Pop()

		fmt.Println(frontIndex.(int))

		if maxLen < rearIndex - frontIndex.(int) - 1 {
			maxLen = rearIndex - frontIndex.(int) - 1
		}

		rearIndex = frontIndex.(int)

		if myStack.Empty() {

			if maxLen < rearIndex {
				maxLen = rearIndex
			}

			break
		}
	}

	return maxLen
}