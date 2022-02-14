package main

import "fmt"

var num = 10 // 全局变量

func testGlobal() {
	num := 100
	name := "沙河娜扎"
	// 可以在函数中使用变量
	// 1、先在自己函数中查找，找到了就用自己函数中的
	// 2、函数找不到变量就往外层找全局变量
	fmt.Println("变量num:", num)
	fmt.Println(name)
}
func main() {
	testGlobal()
	// fmt.Println(name) // 外层不能访问到内层函数的内部变量(局部变量)

	// 函数可以作为变量
	abc := testGlobal // 函数当做值传给变量
	fmt.Printf("%T\n", abc)
	abc() // 可以通过abc()直接调用

	// 变量i此时只在for循环的语句代码块中生效
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// fmt.Println(i) // 外层访问不了内部 for 语句块中的变量
}

// 如果一定要访问内部的变量，我们需要在函数中利用返回值的形式，将这个变量向外返回，然后在外层定义一个变量接收返回值
// for ,switch 这些语句块，外层也是访问不了内层的
