/*
    author twitter : @mmcgrana
    https://gobyexample.com/pointers

    Go by Example: Pointers

 Go는 포인터를 지원하며
 당신의 프로그램안에서 값과 기록에게 참조를 전달할 수 있다.

*/

package main

import "fmt"

// 우리는 당신에게 포인터가 값에 대해서 어떻게 작동하는지 2개의 function과 함께 보여줄 것이다.
// 하나의 func는 zeroval 이고 두번째 변수는 zeroptr 이다.
// zeroval 은 int형의 파라미터를 가지며, 값형태로 아규먼트가 전달된다.
// zeroval은 그것을 호출한 function과는 구별된 ival의 복사본을 얻게된다.
func zeroval(ival int) {
	ival = 0
}

// 그에 반해서 `zeroptr` 은 `*int` 파라미터를 가진다. 이것은 Int형 포인터를 가진다는 것을 의미하며,
// 이 함수 몸체에서 *iptr은 포인터의 참조를 찾아가서 실제 메모리 주소로 접근하여
// 값을 할당합니다.  (생략번역함)
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// `&i` 는 'i'에 대한 메모리주소를 준다.
	// i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 포인터는 또한 출력될 수 있다.
	fmt.Println("pointer:", &i)
}
