package main

import "fmt"

func main() {
	str := "Hello Golang!"
	revesed := ReverseString(str)
	fmt.Println("Orignal String: ", str)
	fmt.Println("Revesed string: ", revesed)
}

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
