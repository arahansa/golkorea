// 14_상수 선언
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "안녕"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/*
 상수는 const 키워드와 함께 변수처럼 선언가능하다.
 상수는 character, string, boolean, 숫자타입중의 하나 가능하다
 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
