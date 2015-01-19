/*
    author twitter : @mmcgrana
    https://gobyexample.com/methods

    Go by Example: Methods

 Go는 구조체 타입 위에서 메소드를 만들 수 있다.

*/

package main

import "fmt"

type rect struct {
	width, height int
}

// 이 `area` 메소드는 `*rect`의 리시버타입을 가지고 있다. (포인터 형식이란 얘기: 역주)
func (r *rect) area() int {
	return r.width * r.height
}

// 메소드는 포인터나 값 전달 타입으로 생성할 수 있다. 여기에 값전달 타입이 있다.
//(역주: 포인터 안 붙임)
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 우리의 구조체에서 생성된 두개의 메소드를 불러본다.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go는 자동으로 메소드호출을 위하여 값과 포인터의 변환을 처리한다.
	// 당신은 메소드가 호출될때 복사를 피하거나 메소드가 receiving 구조체를 변경시키기 위해
	//포인터 리시버 타입을 사용할 수도 있다.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
