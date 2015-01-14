/*
	author twitter : @mmcgrana
	https://gobyexample.com/values

	Go by Example: Values
*/
package main

import "fmt"

func main() {

	// 문자열은 + 로 연결될 수 있다. 이건 자바와 비슷하군?!
	fmt.Println("go" + "lang")

	// 숫자와 소수들
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 참거짓 연산자들.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
