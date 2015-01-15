/*
	author twitter : @mmcgrana
	https://gobyexample.com/arrays

	Go by Example: Arrays

   Go에서 배열은 특정 길이를 가진, 번호가 붙은 요소들의 나열이다.
*/
package main

import "fmt"

func main() {

	// a는 상수배열이고 크기 5를 가진다. 요소타입과 길이는 둘다 배열타입의 일부다..
	// int배열의 기본값은 zero-value 즉 0을 가진다.
	var a [5]int
	fmt.Println("emp:", a)

	// 배열[순번] = 값 으로 배열의 값을 지정하고, 배열[순번]으로 값을 불러온다.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 내장된 len으로 배열의 길이 가져오기
	fmt.Println("len:", len(a))

	//다음과 같이 한줄로 배열을 선언하고, 초기화할수도 있다.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 배열은 1차원일수도있고, 다차원이 될 수도 있다(그냥 짧게 번역함)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
