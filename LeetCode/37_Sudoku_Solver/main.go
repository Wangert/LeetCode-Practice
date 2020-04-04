package main

import (
	"fmt"
)

/**
编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。

Note:

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sudoku-solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Sudoku:")
	for _, rowArray := range board {
		for _, num := range rowArray {
			fmt.Print(string(num))
			fmt.Print(" ")
		}
		fmt.Println()
	}

	solveSudoku(board)

	fmt.Println("Result:")
	for _, rowArray := range board {
		for _, num := range rowArray {
			fmt.Print(string(num))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func solveSudoku(board [][]byte)  {

	//定义行的计数map数组
	rowMaps := make([]map[int]bool, 9)
	//定义列的计数map数组
	cowMaps := make([]map[int]bool, 9)
	//定义子数独的计数map数组
	subSudokuMaps := make([]map[int]bool, 9)
	//初始化所有行列，以及子数独的状态
	initState(rowMaps, cowMaps, subSudokuMaps, board)

	placeNum(0, rowMaps, cowMaps, subSudokuMaps, board)

}

//初始化数独未填数字状态
func initState(rowMaps, cowMaps, subSudokuMaps []map[int]bool, board [][]byte) {

	//初始化map
	for k := 0; k < 9 ; k++ {
		map1 := make(map[int]bool)
		map2 := make(map[int]bool)
		map3 := make(map[int]bool)

		rowMaps[k] = map1
		cowMaps[k] = map2
		subSudokuMaps[k] = map3
	}

	for i, rowArray := range board {
		for j, num := range rowArray {

			if num == []byte(".")[0] {
				continue
			} else {

				(rowMaps[i])[int(num) - 48] = true
				(cowMaps[j])[int(num) - 48] = true

				//判断该数位置属于哪一个子数独
				subSudoku := (j / 3) * 3 + i / 3
				(subSudokuMaps[subSudoku])[int(num) - 48] = true
			}
		}
	}
}

//填充数字
func placeNum(count int, rowMaps, cowMaps, subSudokuMaps []map[int]bool, board [][]byte) bool {

	if count == 81 {
		return true
	}
	//计算行列位置，属于哪个子数独
	row := count / 9
	cow := count % 9
	subSudoku := (cow/3)*3 + row/3

	if board[row][cow] == byte('.') {
		for k := 1; k <= 9; k++ {
			//满足数独要求
			if (rowMaps[row])[k] == false && (cowMaps[cow])[k] == false && (subSudokuMaps[subSudoku])[k] == false {

				(rowMaps[row])[k] = true
				(cowMaps[cow])[k] = true
				(subSudokuMaps[subSudoku])[k] = true

				board[row][cow] = byte(48 + k)

				if placeNum(count+1, rowMaps, cowMaps, subSudokuMaps, board) {
					return true
				} else {

					(rowMaps[row])[k] = false
					(cowMaps[cow])[k] = false
					(subSudokuMaps[subSudoku])[k] = false
					board[row][cow] = byte('.')
				}

			}
		}
	} else {
		return placeNum(count + 1, rowMaps, cowMaps, subSudokuMaps, board)
	}

	return false
}