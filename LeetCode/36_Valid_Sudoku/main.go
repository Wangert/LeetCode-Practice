package main

/**
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

}

func isValidSudoku(board [][]byte) bool {

	//定义行的计数map数组
	rowMaps := make([]map[byte]bool, 9)
	//定义列的计数map数组
	cowMaps := make([]map[byte]bool, 9)
	//定义子数独的计数map数组
	subSudokuMaps := make([]map[byte]bool, 9)
	//初始化map
	for k := 0; k < 9 ; k++ {
		map1 := make(map[byte]bool)
		map2 := make(map[byte]bool)
		map3 := make(map[byte]bool)

		rowMaps[k] = map1
		cowMaps[k] = map2
		subSudokuMaps[k] = map3
	}

	//是否有效
	isValid := true

	for i, rowArray := range board {
		for j, num := range rowArray {

			if num == []byte(".")[0] {
				continue
			}

			//判断行是否重复
			if (rowMaps[i])[num] == true {
				isValid = false
				goto END
			} else {
				(rowMaps[i])[num] = true
			}
			//判断列是否重复
			if (cowMaps[j])[num] == true {
				isValid = false
				goto END
			} else {
				(cowMaps[j])[num] = true
			}

			//判断该数位置属于哪一个子数独
			subSudoku := (j / 3) * 3 + i / 3
			//判断子数独中是否重复
			if (subSudokuMaps[subSudoku])[num] == true {
				isValid = false
				goto END
			} else {
				(subSudokuMaps[subSudoku])[num] = true
			}
		}
	}

END:
	return isValid

}
