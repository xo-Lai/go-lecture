package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d \n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating Slice")
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		// 查看怎么扩容
		printSlice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)

	// 数组需要初始值
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// 使用内建make创建数组，可以声明数组长度和容量，且不用初始化值，默认值0
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Coping Slice")
	// 复制数组（目标对象，数据源）
	copy(s2, s1)
	fmt.Println(s2)

	fmt.Println("Deleting elements from slice")
	// 系统没有内建的函数
	// s2 =[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]  删除数组8元素
	// s2[:3]+s2[4:]  就删除8 元素
	s2 = append(s2[:3], s2[4:]...)

	//[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0],len=15,cap=16
	printSlice(s2)

	fmt.Println("popping form front")
	front := s2[0]
	s2 = s2[1:]
	// 2
	fmt.Println(front)

	fmt.Println("popping form back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	// 0
	fmt.Println(tail)
	// [4 6 0 0 0 0 0 0 0 0 0 0 0],len=13,cap=15
	printSlice(s2)

}
