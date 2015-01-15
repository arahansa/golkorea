/*
	author twitter : @mmcgrana
	Go by Example: Switch

	Go by Example: Switch

  스위치 는 많은 경우들을 표현할 수 있다.(의역)
*/
package main

import "fmt"
import "time"

func main() {

	// 기본적인 스위치문
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//  콤마들(,)을 사용해서 여러 조건들을 하나의 case 문에 사용가능하다.
	// default 로 아무조건도 걸리지 않는 경우를 처리할 수 있다(의역)
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 표현식이 없는 swich는 ifelse 처럼 쓸 수 있다.
	// 상수가 아닌 것도 case 표현식에 적을 수 있다.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}
