// 33_빈슬라이스
package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

/*
슬라이스의 zero value는 nil 입니다.

nil 슬라이스는 길이와 최대 크기가 0입니다.

(슬라이스에 대해 더 알고 싶다면 다음 글을 읽어보세요.)
http://golang.org/doc/articles/slices_usage_and_internals.html


 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
