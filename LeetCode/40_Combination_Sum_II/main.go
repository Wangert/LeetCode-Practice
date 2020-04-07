package main

import (
	"sort"
	"fmt"
)

/**
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

//记录结果
type Result struct {
	r [][]int
}

func main()  {

	candidates := []int{10,1,2,7,6,1,5}
	fmt.Println(combinationSum2(candidates, 8))

}

func combinationSum2(candidates []int, target int) [][]int {

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

		if index != 0 {
			if candidate == candidates[index - 1] {
				continue
			}
		}

		temp := make([]int, len(currentState))
		copy(temp, currentState)

		if candidate < lack {
			temp = append(temp, candidate)
			result.addNum(temp, candidates[index + 1:], lack - candidate)
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