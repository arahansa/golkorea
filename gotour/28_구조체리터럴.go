// 28_구조체리터럴
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}

/*
 구조체 리터럴은 필드와 값을 나열해 구조체를 새로 할당하는 방법이다.
 원하는 필드를 {Name:value} 구문을 통해 할당 가능하다.
 특별한 접두어 &를 사용하면 구조체 리터럴에 대한 포인터 생성 가능하다.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
