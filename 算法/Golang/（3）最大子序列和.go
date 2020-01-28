package main

import "fmt"

func getMaxNum(slice []int,left int,right int)  int{

	if left == right {
		return slice[left]
	}

	midden := (left + right) / 2

	//求包含中点向左序列的最大值
	var leftMax int
	var sum int
	sum = 0
	for i := midden; i >= left; i-- {
		sum += slice[i]
		if sum > leftMax {
			leftMax = sum
		}
	}

	//包含中点向右的最大序列值
	var rightMax int
	sum = 0
	for i := midden + 1; i <= right; i++ {
		sum += slice[i]
		if sum>rightMax {
			rightMax = sum
		}
	}

	leftNum := getMaxNum(slice,left,midden)
	rightNum := getMaxNum(slice,midden+1,right)

	var max int
	if leftNum > rightNum{
		max = leftNum
	}else {
		max = rightNum
	}

	if max > (leftMax+rightMax) {
		return max
	}else {
		return leftMax+rightMax
	}

}

func main()  {

	slice := []int{1,2,-3}
	num := getMaxNum(slice,0,len(slice) - 1)
	fmt.Println(num)
}