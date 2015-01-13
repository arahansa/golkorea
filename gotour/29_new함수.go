// 29_new함수 : 모든필드가 0이 할당된 T의 반환
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}

/*
new 는 모든 필드가 0이 할당된 T타입의 포인터를 반환함.
var t *T = new(T) 또는 t := new(T)

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
