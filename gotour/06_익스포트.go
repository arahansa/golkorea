// 대문자를 써야 외부참조 가능!! 주의
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
	//fmt.Println(math.Pi)
}

/*
내용 : pi를 Pi 로 변경해야 한다. 외부참조시에 대문자로 쓰여져 exported 된 것들만 접근 가능하다.
 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
