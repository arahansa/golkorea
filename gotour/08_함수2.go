// 메서드 타입정하는 것
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/*
내용 : 두개 이상의 매개변수가 같은 타입일 때(type) 같은 타입을 취하는 마지막 매개변수만 타입 명시하고 생략 가능하다.
예를 들어서 x int, y int 를 x, y int 로 생략가능하다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
