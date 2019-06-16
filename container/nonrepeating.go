package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	//lastOccurred := make(map[byte]int)
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println("start",start)
	return  maxLength
}

func main() {
	s:="abcabcbb"
	fmt.Println(lengthOfNonRepeatingSubStr(s))

	s1:="abcabcd"
	fmt.Println(lengthOfNonRepeatingSubStr(s1))

	s2:="12344ss"

	fmt.Println(lengthOfNonRepeatingSubStr(s2))

}
