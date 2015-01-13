// 38_맵리터럴2
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

/*
만약 가장 상위의 타입이 타입명이라면 리터럴에서 타입명을 생략해도 됩니다.
"Bell Labs": {40.68433, -74.39967} 또는
"Bell Labs": Vertex{40.68433, -74.39967}
는 같은 표현입니다.


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
