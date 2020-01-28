package main

import (
	"fmt"
)

func getRepeatNum(numArr []int)  []int{

	var tagMap map[int]int
	tagMap = make(map[int]int)

	//or tagMap := make(map[int]int)

	var result []int
	for _,num := range numArr {

		if tagMap[num] == 0  {
			tagMap[num] = 1
		}else {
			result = append(result,num)
		}
	}
	return result
}


func main()  {

	numArr := []int{2,3,1,0,2,5,3}
	result := getRepeatNum(numArr)

	fmt.Println(result)
}