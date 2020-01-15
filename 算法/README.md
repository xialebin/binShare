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
```
场景：根据股票的折线图，求出买入与卖出的利益最大化的日期，可以根据当天的价格与前一天的价格进行比较存入数组，最后求最大数组和问题。如下：
```
![image](https://github.com/xialebin/binStudy/blob/master/%E7%AE%97%E6%B3%95/z_images/jump_3_1.png)

求最大子序列的和可以使用暴力法求解，即将所有的子序列展开，判断和的最大值，这里不再赘述。下面使用分治法的思想和递归来处理。<br>

思路：<br>

将数组序列平分成两段<br>
数组的最大子序列只会出现下面三种情况：<br>
1、最大子序列全部在左边 <br>
2、最大子序列全部在右边<br>
3、最大子序列横跨中点<br>

所以我们可以这样求解：<br>

求得左边最大子序列的和A<br>
求得右边最大子序列的和B<br>
根据左边最后一个值向左求最大子序列 a<br>
根据右边最后一个值向右求最大子序列 b<br>

比较返回 A,B,a+b 三者中最大的一个即为最大子序列的和<br>
而A、B的求解可以通过递归求解。<br>

tip:a+b即是横跨中点的子序列，因为从中点开始往两边遍历，肯定包含中点。<br>

```go
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

	if max > (leftNum+rightNum) {
		return max
	}else {
		return leftNum+rightNum
	}
}

func main()  {
	slice := []int{1,2,-3}
	num := getMaxNum(slice,0,len(slice) - 1)
	fmt.Println(num)
}
```


