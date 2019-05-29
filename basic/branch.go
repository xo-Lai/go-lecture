package main

import (
	"io/ioutil"
	"fmt"
)

// no:3
func main() {

	readfile()
}

// 1. if else 使用
func readfile() {
	const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename);
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	// 简写
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}

// 2. switch 语句 自动break
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		// 打印错误
		panic(fmt.Sprintf("Wrong Score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
