// 50_메소드
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

/*
고에는 클래스가 없습니다. 하지만 메소드를 구조체(struct)에 붙일 수 있습니다.

메소드 리시버(method receiver) 는 func 키워드와 메소드의 이름 사이에 인자로 들어갑니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
