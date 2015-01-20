package main

import "fmt"

func main() {
	fmt.Println("문자열을 입력하시오")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
