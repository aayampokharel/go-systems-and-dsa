package main

import (
	"fmt"
	"strconv"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var returnStr string
	for _, val := range strs {
		returnStr += fmt.Sprintf("%v%v%v", len(val), "#", val)
	}
	return returnStr
}

func (s *Solution) Decode(encoded string) []string {
	var count int
	var returnStrSlice []string
	for key := 0; key < len(encoded); key++ {

		if string(encoded[key]) == "#" {
			num, err := strconv.Atoi(string(encoded[key-count : key]))
			if err != nil {
				continue
			}
			returnStrSlice = append(returnStrSlice, string(encoded[key+1:key+1+num]))
			key = key + num
			count = 0
		} else {
			count++
		}
	}
	return returnStrSlice
}
