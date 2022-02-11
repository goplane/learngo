package main

import "fmt"

// 数组相关
func main() {
	var a [4]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)
	//int 类型默认初始值就是0,[4] 4代表长度

	// 数组的初始化
	// 1、定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "广州", "南京"}
	fmt.Println(cityArray)
	fmt.Println(cityArray[0])
	fmt.Println(cityArray[2])

	// 2、编译器推导数组的长度
	var boolArray = [...]bool{true, false, false, true} // 不加...也可以
	fmt.Println(boolArray)

	// 3、使用索引值方式初始化
	var langArray = [...]string{1: "golang", 3: "python", 7: "java"}
	fmt.Println(langArray) //空的字符串占一个空字符串

	fmt.Printf("%T\n", langArray) //%T查看类型，打印出来长度为8

}
