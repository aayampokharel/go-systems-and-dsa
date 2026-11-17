package main

func hasDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, val := range nums {
		m[val]++
	}
	for key := range m {
		if m[key] > 1 {
			return true
		}
	}
	return false

}

// func bestSolution(nums []int) bool {    //!Easier solution for day-to-day coding.
//      m:=make(map[int]int,len(nums))
// 	for _,val:=range nums{
//         m[val]++
// 		if m[val]>1{
// 			return true
// 		}
// 	}
// 	return false

// }

func BestSolution(nums []int) bool {
	seen := make(map[int]struct{}, len(nums)) // Hint capacity for performance
	for _, val := range nums {
		if _, exists := seen[val]; exists {
			return true // Exits immediately on the first duplicate found!
		}
		seen[val] = struct{}{}
	}
	return false
}
