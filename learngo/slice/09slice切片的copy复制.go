package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5) //长度是5 ，容量是5
	c := b
	copy(b, a) // a切片数据拷贝到b切片内
	fmt.Println(a, b)
	b[0] = 100
	fmt.Println(a, b) // 证明a和b切片的内存地址值不是同一个
	fmt.Println(c)

}
