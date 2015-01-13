// 37_맵
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/*
맵은 값에 키를 지정합니다.
맵은 반드시 사용하기 전에 make 를 명시해야합니다. (주의: new 가 아닙니다)
make 를 수행하지 않은 nil 에는 값을 할당할 수 없습니다.

Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
