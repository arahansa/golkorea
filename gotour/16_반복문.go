// 16_반복문
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
내용 : Go 에서는 반복문이 For 밖에 없다.
 하지만 소괄호() 가 필요없고 중괄호는 무조건 필요 {}

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
