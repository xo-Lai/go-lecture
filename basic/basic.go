package main

import "fmt"

// no:1

// 1. 打印
func main() {
	fmt.Println("Hello world")
	variableZeroValue()
}

// 2. 定义变量
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 3. 变量赋初值
func variableInitialValue() {
	var a, b int = 3, 5
	var s string
	fmt.Println(a, b, s)
}

// 4. 自动类型推断
func variableTypeDeduction() {
	var a, b, c, d, e = 1, 2, true, "def", 2
	fmt.Println(a, b, c, d, e)
}

// 5. var 缩写
func variableShorter() {
	a, b, c := 1, 2, 3
	a = 2
	fmt.Println(a, b, c)
}

// 6. 函数外变量不是全局变量  是package 变量
//var aa = 3
//var ss = "kkk"
//bb :=ture  函数外 var不能用: 缩写
//var bb = true
// 可以统一声明
var (
	aa = 3
	ss = "kkk"
	bb = true
)



