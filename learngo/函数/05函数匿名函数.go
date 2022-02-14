package main

import "fmt"

//匿名函数和闭包
func main() {
	// func() {
	// fmt.Println("匿名函数")
	// }() //直接在后面接入()，不能换行，可以直接调用匿名函数

	// 上面是匿名函数第一种运行方式，下面是第二种运行方式，变量赋值形式，将函数当值传递给变量

	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()
}
