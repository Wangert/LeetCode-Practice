package main

import (
	"sort"
	"fmt"
)

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

//记录结果
type Result struct {
	r [][]int
}

func main()  {

	candidates := []int{2, 3, 5}
	fmt.Println(combinationSum(candidates, 8))

}

func combinationSum(candidates []int, target int) [][]int {
	//从小到大排序
	sort.Ints(candidates)
	//存储结果
	result := &Result{}
	result.r = make([][]int, 0)
	//当前状态
	currentState := make([]int, 0)

	result.addNum(currentState, candidates, target)

	return result.r
}

//添加数字
func (result *Result) addNum(currentState, candidates []int, lack int) {

	for index, candidate := range candidates {

		temp := make([]int, len(currentState))
		copy(temp, currentState)

		if candidate < lack {
			temp = append(temp, candidate)
			result.addNum(temp, candidates[index:], lack - candidate)
		} else if candidate > lack {
			break
		} else {
			temp = append(temp, candidate)
			result.r = append(result.r, temp)
			break
		}
	}

	return
}