/*
	author twitter : @mmcgrana
	https://gobyexample.com/functions

	Go by Example: Functions
	_Functions_ 은 Go의 중심이다. . 몇가지 예제와 함께 function을 알아보자.
*/
package main

import "fmt"

// 여기 두개의 int를 받아서 합계를 int로 돌려주는 예제가 있다.
func plus(a int, b int) int {

	// Go 분명한 리턴을 필요로 한다.
	//다시말하자면 자동으로 표현식의 마지막값을 리턴하지 않는다.
	return a + b
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}

// todo: coalesced parameter types
