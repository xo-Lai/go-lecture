package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

// no:2

/**
- bool,string
- (u)int,(u)int8,(u)int16,(u)int64,uintptr  默认长度根据操作系统来
- byte,rune(指针，char,32字节)
- float32 ,float64,complex64,complex128
 */
func main() {
	euler()
	enums()
}

//1.欧拉公式   complex复数
func euler() {
	//复数绝对值
	c := 3 + 4i
	// |3+4i|=5
	fmt.Println(cmplx.Abs(c))
	fmt.Println(
		cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

// 2. 强类型转换 没有隐式转换
func triangle() {
	a, b := 3, 4
	calculateTriangle(a, b)
}

func calculateTriangle(a, b int) int {
	//c=math.Sqrt(a*a+b*b)  编译错误  没有隐式转换
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 3. 常量

const filename = "abc.txt"

func consts() {
	const a, b = 3, 4
	c := int(math.Sqrt(a*a + b*b))
	fmt.Print("c = %d", c)

}

//4.枚举 就是用const常量定义

func enums() {
	//const (
	//	cpp=0
	//	java=1
	//	python=2
	//	golang=3
	//)

	// 底下下划线是指跳过
	const (
		cpp    = iota
		_
		java
		python
		golang
	)

	//b,kb,mb,gb,pb
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, pb)
}
