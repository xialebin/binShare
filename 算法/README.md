## 算法

>说明：<br>

该算法列表中的题目是笔者从网络、书籍上整理而来，且用自己的理解，使用C++、Go、PHP进行实现，解析内容使用Go语言进行举例。目的是记录学习点滴，与大家共同进步。笔者能力有限，仅供参考。

>题目目录：<br>
>>（1）[求猴子大王的编号](#jump_1)<br>
>>（2）[冒泡排序](#jump_2)<br>
>>（3）[求数组的最大子数组或最大子序列和](#jump_3)<br>


>解析：<br>

<span id="jump_1">（1）求猴子大王的编号</span><br>
```
一群猴子排成一圈，按1、2、...、n依次编号。然后从第一只开始数，数到第m只把它提出圈，从它后面再开始数，再数到第m只，把它踢出去...，如此不停的进行下去，直到最后只剩下一只猴子为止，那只猴子就叫做大王。要求编程模拟这个过程，输入m,n 输出最后那个大王的编号。
```
首先，需要选用合适的数据结构，根据题目，我首先能想到是构造一个带有序号的双向链表，记录序号，每次进行链表的删除操作，然后更新序号，直到链表长度为1。而在具体的编程语言中，就可以使用具体的类型来替代这个链表的作用，在GO语言中，可以使用切片来实现，因为切片的长度是动态的，且自带append操作，可以构造出删除元素等操作。同理，C++中，使用<vector>向量，自带erase操作，可以满足需要，PHP中使用array数组即可，动态数组，且自带各种操作函数。
```go
package main

import (
	"fmt"
)

func getKingNum(n int,m int) int{

	//声明切片
	var amount []int

	//初始化，往切片中追加猴子
	for i:=0;i<n;i++  {
		amount = append(amount,i+1)
	}
	//计算第一次要删除的猴子的索引下标
	startIndex := (m-1) % len(amount)
	
	//无限循环   
	for true { 
		if  len(amount) <= 1{
			break
		}
		//移除操作 即重新拼装切片，第二个参数即从startIndex+1到切片之后所有的元素，GO支持这样的写法
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
```

<span id="jump_2">（2）冒泡排序</span><br>

```
冒泡排序，对数字序列 85496进行从小到达的冒泡排序
```

冒泡排序，顾名思义较大（或较小）的数就像泡泡一样一步步浮出水面。程序实现的过程是，双层循环，外层循环是按照顺序从第一个位置遍历到最后一个，内层循环的操作是，从开始位置对应的数据与下个位置的数据进行比较，进行移动（冒泡的过程），而循环条件的边界受到外层循环的影响，即为len-i-1,也就是，外层循环记录了已经进行了几轮比较，而每次外层循环即一轮冒泡，就会把最大（最小）的数送到最后，所以内存循环需要将已经循环的次数扣除，而内层循环是与下一个数比较的策略，即与n+1进行比较，所以循环边界需要减1，否则会下标溢出。

```go

package main
import "fmt"

func bubbleSort(data []int) []int {
	//外层循环 遍历整个数据
	for i:=0;i<len(data);i++  {
                //内层循环即一轮冒泡的过程
		for j:=0;j<len(data)-i-1;j++  {
			//这里的j+1决定了为什么内层循环要-1
			if  data[j]>data[j+1]{
                data[j+1],data[j] = data[j],data[j+1]
			}
		}
	}
	return  data
}

func main()  {
        //使用切片模拟冒泡过程
	arr := []int{8,5,4,9,6}
	rel := bubbleSort(arr)
	fmt.Println(rel)
}

```
<span id="jump_3">（3）求数组的最大子数组或最大子序列和</span><br>


场景：根据股票的折线图，求出买入与卖出的利益最大化的日期，可以根据当天的价格与前一天的价格进行比较存入数组，最后求最大数组和问题。如下：

![image](https://github.com/xialebin/binStudy/blob/master/%E7%AE%97%E6%B3%95/z_images/jump_3_1.png)





