package main

import "fmt"

func Mutate(twoDSlice [][]int) {
	s := make([][]int, 0, len(twoDSlice))

	for _, sliceAsVal := range twoDSlice {
		s = append(s, sliceAsVal)
	}
	s[0][0] = 999
}
func main() {
	var twoDSlice [][]int = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	Mutate(twoDSlice)
	fmt.Println(twoDSlice)
}
