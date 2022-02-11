package main

import "fmt"

// 数组相关
func main() {
	// 数组的遍历
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// 1 for循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	// 2 for range 遍历
	for index, value := range cityArray {
		fmt.Println(index, value)
		//比如 把value或index去掉，只打印一个值，出来的只有索引
		// 如果只是想接收到数据的值，可以 _,value
	}
}
