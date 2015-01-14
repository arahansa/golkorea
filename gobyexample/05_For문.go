/*
	author twitter : @mmcgrana
	https://gobyexample.com/constants

	Go by Example: For

   `for` 는 Go의 유일한 반복문이고, 여기 for문의 세가지 유형이 있다.
*/
package main

import "fmt"

func main() {

	// 가장 기본적 타입인 단일조건문
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 많이 본 for문이죠..? ㅎㅎ 초기식/조건식/후처리 for문
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 조건식이 없는 for문은 break를 만나거나
	//enclosing(?동봉된?둘러쌓여진?) 함수에서 return될때까지 계속 반복합니다.
	for {
		fmt.Println("loop")
		return
	}

}
