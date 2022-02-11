package main

import "fmt"

func main() {
	// 多维数组最外层[4]可以用[...]表示，内层不可以使用
	// [4]是内部小数组的个数，[2]是小数组内部值的个数，string 是类型
	cityArray := [4][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}
	fmt.Println(cityArray)
	fmt.Printf(cityArray[2][1]) //打印 成都
	// 二维数组遍历
	for _, v1 := range cityArray {

		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 数组在go语言中是值类型，是完整的值拷贝

}
