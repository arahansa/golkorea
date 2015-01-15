/*
	author twitter : @mmcgrana
	https://gobyexample.com/if-else

	Go by Example: IfElse

   Go에서의 If else 분기문은 간단하다
*/
package main

import "fmt"

func main() {

	// 이것은 기본적인 예제다.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// else 없이 사용 가능하며
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 문장 하나가 조건식 앞에 나올 수 있다.; 이 문장에 쓰여진 변수들은 이
	// if else 문 어디서도 사용 가능하다..
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
