// 55_에러
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

/*
에러 문장(string)으로 자신을 표현할 수 있는 것은 모두 에러입니다.
이 아이디어는 문자열(string)을 반환하는 하나의 메소드 Error 로 구성된 내장 인터페이스 타입 error 에서 나왔습니다.

type error interface {
	Error() string
}
fmt 패키지의 다양한 출력 루틴들은 error 의 출력을 요청받았을 때 자동으로 이 메소드를 호출합니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
