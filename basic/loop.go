package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"text/scanner"
)

// no：4
func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), //1011 -> 1101
		convertToBin(0),
	)
}

// 1. for  使用， 数字转二进制
func convertToBin(n int) string {

	if n == 0 {
		return "0"
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		// 数字不能直接加字符串 需要转字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// for 使用 读取文本逐条打印  for 没有条件  类似foreach
func printFile(fileName string) {
	// 打开文件
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

}

// for 死循环
func forover() {
	for {
		fmt.Println("123")
	}
}
