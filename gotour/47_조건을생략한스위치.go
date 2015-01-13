// 47_조건을생략한스위치
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

/*
스위치에서 조건을 생략하면 " switch true " 와 같습니다.
만약 긴 if-then-else 를 작성해야 할 때, 이 구조를 사용하면 코드를 깔끔하게 작성할 수 있습니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
