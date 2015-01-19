/*
    author twitter : @mmcgrana
    https://gobyexample.com/recursion

    Go by Example: Recursion

 Go는 재귀함수를 지원하며 여기 잘 알려진 팩토리얼 예제를 보여주겠다~
*/
package main

import "fmt"

// 이 fact function은 func(0)d이 될때까지 자기 자신을 호출한다.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
