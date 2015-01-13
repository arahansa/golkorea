// 44_연습_피보나치
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
함수를 가지고 놀아봅시다.

fibonacci 함수를 구현합니다. 이 함수는 이어지는 피보나치 수를 반환하는 함수 (클로져)를 반환해야 합니다.


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
