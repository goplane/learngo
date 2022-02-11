package main

import "fmt"

// 数组的长度固定的并且数组长度属于类型的一部分，所以数组有很多局限性。例如
func main() {
	x := [3]int{1, 2, 3}
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	fmt.Println(sum)
}
