// 13_변수의짧은선언 := 로 바로 선언함.
package main

import "fmt"

func main() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}

/*
 := 를 선언하면 var와 타입생략가능하다 !  로컬변수 정도에 쓰기좋다.
 함수밖에서는 := 불가.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
