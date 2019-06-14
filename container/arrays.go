package main

import "fmt"

// 打印数值，这里是用来演示 数组是值类型，因为当arr2 当参数会报错
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func main() {

	// 数组是值类型

	var arr1 [5] int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	// 四行五列
	var grid [4][5] int
	fmt.Println(arr1, arr2, arr3)
	fmt.Print(grid)

	// 数组遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// range 可以获得下标和值
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Printf("i=%d,v=%d.", i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(arr1)
	printArray(arr3)

}
