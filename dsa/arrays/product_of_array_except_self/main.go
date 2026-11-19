package main


func productExceptSelf(nums []int) []int {
var returnInt []int
var prodVal int
for key1,_:=range nums{
	prodVal=1
  for key2,val2:=range nums{
     if key2!=key1{
		prodVal*=val2
	 }
  }
  returnInt=append(returnInt,prodVal)
}
return returnInt
}

func BestSolution(nums []int) []int {
	var leftNumsProd []int

	var prodVal int = 1

	for key, _ := range nums {
		if key == 0 {
			prodVal = 1
		} else {
			prodVal *= nums[key-1]

		}
		leftNumsProd = append(leftNumsProd, prodVal)

	}

	for i := len(leftNumsProd) - 1; i >= 0; i-- {
		if i == len(leftNumsProd)-1 {
			prodVal = 1
			continue
		} else {
			prodVal *= nums[i+1]
		}
		leftNumsProd[i] *= prodVal
	}

	return leftNumsProd
}
