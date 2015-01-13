// 51_메소드2
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

/*
사실 메소드는 구조체(struct) 뿐 아니라 아무 타입(type)에나 붙일 수 있습니다.

다른 패키지에 있는 타입이나 기본 타입들에 메소드를 붙이는 것은 불가능합니다.


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
