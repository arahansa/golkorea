// 58_연습_HTTP핸들러
package main

import (
	"net/http"
)

func main() {
	// your http.Handle calls here
	http.ListenAndServe("localhost:4000", nil)
}

/*
아래 나오는 타입을 구현하고 그 타입의 ServeHTTP 메소드를 정의하십시오.
그 메소드를 당신의 웹 서버에서 특정 경로를 처리할 수 있도록 등록하십시오.

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}
예컨대, 당신은 아래와 같이 핸들러를 등록할 수 있어야 합니다.

http.Handle("/string", String("I'm a frayed knot."))
http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
