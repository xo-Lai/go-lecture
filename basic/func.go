package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

// no:5
func main() {
	//-----------
	q, r := div2(13, 3)
	fmt.Println(q, r)

	//-----------
	if result, err := eval2(3, 4, "+"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	//-----------
	fmt.Println(apply(pow, 2, 4))
	// 使用匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	//-----------
	fmt.Println(sum(1, 2, 3, 4, 5))

	//-----------
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b) //4 3

	c, d := 5, 6
	c, d = swap2(c, d)
	fmt.Println(c, d) //6 5

}

// 1. 函数计算
func eval(a, b int, op string) int {

	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation" + op)
	}
}

// 2. 带余除法 13/3=4....1
func div(a, b int) (int, int) {
	return a / b, a % b
}

// q(quotient):商 r(remainder):余数
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

// 2. 函数计算 不中断执行
func eval2(a, b int, op string) (int, error) {

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation:%s " + op)
	}
}

// 3. 函数式编程,这里的op是一个返回值int的函数
func apply(op func(int, int) int, a, b int) int {
	// 反射
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args  (%d,%d)\n", opName, a, b)
	return op(a, b)
}

// a的b次方
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 4.可变参数列表，参数无数个
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

// 5. 指针 ，交换参数
func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}