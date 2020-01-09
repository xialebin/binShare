package main
import "fmt"

func bubbleSort(data []int) []int {

	for i:=0;i<len(data);i++  {

		for j:=0;j<len(data)-i-1;j++  {
			//这里的j+1决定了为什么内层循环要-1
			if  data[j]>data[j+1]{
				temp := data[j+1]
				data[j+1] = data[j]
				data[j] = temp
			}
		}
	}
	return  data
}

func main()  {
	arr := []int{7,8,9,4,5,7,4,1,2,3,5,6}
	rel := bubbleSort(arr)
	fmt.Println(rel)
}
