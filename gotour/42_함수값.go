// 42_함수값
package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

/*
함수도 값 입니다.
(번역자 : 맨 아래 hypot(3,4) 의 hypot 함수를 Println함수의 인자값 처럼 사용 하고 있습니다.)


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
