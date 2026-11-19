package main

func isValidSudoku(board [][]byte) bool {
	columnMap := make(map[string]int, 10)
	rowMap := make(map[string]int, 10)
	threeIntoThreeMap := make(map[string]int, 10)
	for _, val := range board {
		for _, val2 := range val {
			if string(val2) == "." {
				continue
			}
			rowMap[string(val2)]++
			if rowMap[string(val2)] > 1 {
				return false
			}
		}
		clear(rowMap)

	}
	for i := 0; i < 9; i++ {
		for _, val := range board {
			if string(val[i]) == "." {
				continue
			}
			columnMap[string(val[i])]++
			if columnMap[string(val[i])] > 1 {
				return false
			}

		}
		clear(columnMap)
	}

	for i := 0; i < 9; i += 3 {
		for key, val := range board {
			threeIntoThreeMap[string(val[i+0])]++
			threeIntoThreeMap[string(val[i+1])]++
			threeIntoThreeMap[string(val[i+2])]++
			if key%3 == 2 {
				for key1, val1 := range threeIntoThreeMap {
					if key1 != "." && val1 > 1 {
						return false
					}
				}
				clear(threeIntoThreeMap)
			} else {
			}

		}
	}

	return true

}

// func main() {
//   board := [][]byte{
//     {'1', '2', '.', '.', '3', '.', '.', '.', '.'},
//     {'4', '.', '.', '5', '.', '.', '.', '.', '.'},
//     {'.', '9', '8', '.', '.', '.', '.', '.', '3'},
//     {'5', '.', '.', '.', '6', '.', '.', '.', '4'},
//     {'.', '.', '.', '8', '.', '3', '.', '.', '5'},
//     {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
//     {'.', '.', '.', '.', '.', '.', '2', '.', '.'},
//     {'.', '.', '.', '4', '1', '9', '.', '.', '8'},
//     {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// }
//   fmt.Println(isValidSudoku(board))
// }
