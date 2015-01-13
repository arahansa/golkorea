// 48_심화연습_복소수세제곱근
package main

import "fmt"

func Cbrt(x complex128) complex128 {
}

func main() {
	fmt.Println(Cbrt(2))
}

/*
complex64 타입과 complex128 타입을 통해서 Go 언어의 복소수 지원 기능을 알아봅니다. 세제곱근을 얻기 위해서는, 뉴턴의 방법 (Newton's method)을 적용하여 다음을 반복 수행합니다:

z = z - (z * z * z - x) / (3 * z * z)

알고리즘이 잘 동작하는지 확인하기 위해 2의 세제곱근을 구해봅시다. math/cmplx 패키지에는 Pow 함수가 있습니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
