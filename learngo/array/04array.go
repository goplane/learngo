package main

import "fmt"

// 数组的长度固定的并且数组长度属于类型的一部分，所以数组有很多局限性。
// 切片是一个拥有相同元素的可变长度序列，他是基于数组类型做的一层封装。它非常灵活，支持自动扩容

func main() {
	var a []string
	var b []int

	var c = []bool{false, true}

	fmt.Println(a, b, c)
}
