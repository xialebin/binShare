package main

import (
	"fmt"
)

func getKingNum(n int,m int) int{

	//声明切片
	var amount []int

	//往切片中追加猴子
	for i:=0;i<n;i++  {
		amount = append(amount,i+1)
	}
	//计算第一次要删除的猴子的索引下标
	startIndex := (m-1) % len(amount)

	for true {
		if  len(amount) <= 1{
			break
		}
		//移除操作
		amount = append(amount[:startIndex],amount[startIndex+1:]...)
		//重新计算索引下标
		startIndex = (startIndex+m-1) % len(amount)
	}
	return amount[0]
}

func main()  {
	//对切片的操作
	num := getKingNum(5,2)
	fmt.Println(num)
}