// 메서드 만들기
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/*
내용 : 함수를 설명하는데 매개변수가 뒤에 감을 의식하자.
바로 이곳  : add(x int, y int) int
왜 그런지 궁금하면 http://blog.golang.org/gos-declaration-syntax 이곳을 참조하자.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
