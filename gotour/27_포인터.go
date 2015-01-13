// 27_포인터
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}

/*
Go에는 포인터가 있지만 포인터 연산은 불가
 구조체 변수는 구조체 포인터를 이용해 접근 가능.
포인터를 이용하는 간접적 접근은 실제 구조체에도 영향을 미침

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
