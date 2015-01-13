// 12_변수의초기화
package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(x, y, z, c, python, java)
}

/*
내용 : 초기화 선언을 하면서 타입을 생략할 수 있으며, 초기화값에 따라타입 결정됩니다.

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
