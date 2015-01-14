/*
	author twitter : @mmcgrana
	https://gobyexample.com/variables

	Go by Example: Variables
*/
package main

import "fmt"

func main() {

	// `var` 는 한개나 여러 개의 변수를 선언한다.
	var a string = "initial"
	fmt.Println(a)

	// 동시에 여러개의 변수선언이 가능하다.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 는 변수가 초기화될때 타입을 추론한다.
	var d = true
	fmt.Println(d)

	// 초기화값이 없는 변수는 zero-value. 숫자의 경우는0
	var e int
	fmt.Println(e)

	// `:=` 는 초기화와 선언의 요약이다. 로컬에 쓰임.
	// f:="short"는 var f string = "short"임.
	f := "short"
	fmt.Println(f)
}
