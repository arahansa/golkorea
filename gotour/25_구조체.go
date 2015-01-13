// 25_구조체
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

/*
 구조체는 필드와 데이터들의 조합이다.
 type 선언으로 struct 의 이름을 정할수 있다.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
