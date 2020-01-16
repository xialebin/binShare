package main

import (
	"fmt"
)

func fiboRecurve(n int) int{

	if n < 2 {
		if n == 0 {
			return 0
		}else {
			return 1
		}
	}

	return  fiboRecurve(n-1)+fiboRecurve(n-2)
}


func fibonacii(n int)  int{

	//边界数据处理
	if n<3 {
		if n == 0 {
			return 0
		}else {
			return 1
		}
	}
	numOne := 1
	numTwo := 1
	amount := 2

	for i := 2; i < n; i++{
		amount = numOne+numTwo
		numTwo = numOne
		numOne = amount
	}
	return amount
}

func main()  {
	var n int
	for i := 0; i < 30; i++ {
		n = fiboRecurve(i)
		fmt.Println(n)
	}
}