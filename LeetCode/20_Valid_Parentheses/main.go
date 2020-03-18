package main

import (
	"fmt"
	"LeetCode/20_Valid_Parentheses/stack"
	"reflect"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	s := "({[]})"

	fmt.Println(isValid(s))

}

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}

	//定义括号
	par := []rune("(){}[]")

	//将字符串转换成rune切片
	sRunes := []rune(s)
	//计算字符个数
	sLen := len(sRunes)

	//有效的括号串必须是偶数个括号
	if sLen % 2 != 0 {
		return false
	}

	//定义栈
	myStack := stack.NewStack(reflect.TypeOf([]rune("1")[0]))

	for i := 0; i < sLen; i++ {

		if sRunes[i] == par[0] || sRunes[i] == par[2] || sRunes[i] == par[4] {
			myStack.Push(sRunes[i])
			continue
		}

		if sRunes[i] == par[1] || sRunes[i] == par[3] || sRunes[i] == par[5] {

			tempValue, _ := myStack.Pop()

			if tempValue == nil {
				return false
			}

			if tempValue.(int32) == par[0] && sRunes[i] == par[1] {
				continue
			} else if tempValue.(int32) == par[2] && sRunes[i] == par[3] {
				continue
			} else if tempValue.(int32) == par[4] && sRunes[i] == par[5] {
				continue
			} else {
				return false
			}
		}
	}

	if myStack.Empty() {
		return true
	} else {
		return false
	}
}