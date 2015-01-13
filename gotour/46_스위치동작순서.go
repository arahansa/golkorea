// 46_스위치동작순서
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

/*
스위치의 각 조건은 위에서 아래로 평가합니다. 만약 조건이 참인 case를 찾으면 평가를 마칩니다.

(예를 들어
switch i {
case 0:
case f():
}
에서 i==0 이라면 f() 는 실행하지 않습니다)


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
