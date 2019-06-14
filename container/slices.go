package main

import "fmt"

// []int 表示切片 不是数组，数组是值类型，slice是引用类型
func updateSlice(s []int) {
	s[0] = 100
}

// 切片
func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])

	s1 := arr[:6]
	updateSlice(s1)
	fmt.Println("s1:", s1)

	s2 := arr[:]
	updateSlice(s2)
	fmt.Println("s2:", s2)
	fmt.Println("arr:", arr) // 这里被上面updateSlice 修改了 [100 1 2 3 4 5 6 7 8 9]

	fmt.Println("Reslice")
	fmt.Println(s2) //[100 1 2 3 4 5 6 7 8 9]
	s2 = s2[:5]
	fmt.Println(s2) //[100 1 2 3 4]
	s2 = s2[2:]
	fmt.Println(s2) //[2 3 4]

	fmt.Println("Extending Slice")
	arr[0] = 0
	s1 = arr[2:6]
	s2 = s1[3:5]

	// arr=[0 1 2 3 4 5 6 7 8 9], len(arr)=10, cap(arr)=10
	fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))
	// s1=[2 3 4 5], len(s1)=4, cap(s1)=8
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	// s2=[5 6], len(s2)=2, cap(s2)=5
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Printf("s3,s4,s5-", s3, s4, s5)
	fmt.Printf("arr-", arr)

}
