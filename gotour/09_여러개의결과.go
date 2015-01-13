// 09_여러개의결과.go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/*
내용 : 하나의 함수가 다중리턴이 가능한데, 여기선 두개의 문자열 순서를 바꿔서 리턴함.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
