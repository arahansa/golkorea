/*
    author twitter : @mmcgrana
    https://gobyexample.com/interfaces

    Go by Example: interfaces

 인터페이스는 이름붙여진, 메소드 모음이다.

*/

package main

import "fmt"
import "math"

// 여기 기하학모양을 위한 기본적인 인터페이스가 있다.
type geometry interface {
	area() float64
	perim() float64
}

// 우리의 예제를 위해서 우리는 이 인터페이스를 square와 circle 타입에서 구현할 것이다.
type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Go에서 인터페이스를 구현하기 위해, 우리는 단지
// 인터페이스를 모두 구현하면 된다. 여기 square에 geometry 인터페이스를 구현하는 예제가 있다.
func (s square) area() float64 {
	return s.width * s.height
}
func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

// circle에서 구현하는 것이다.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 만약 변수가 인터페이스 타입을 가지고 있다면, 우리는 그 인터페이스 안의 메소드를 호출할 수 있다.
// 여기에, 이러한 장점을 이용하여 어떤 gemetry 가 오더라도 포괄적으로 사용하는 measure function이 있다.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	// circle 과 square구조체는 둘다 geometry 인터페이스를 구현하고
	// 우리는 이 구조체를 measure 에 아규먼트로 보낼 수 있다
	measure(s)
	measure(c)
}
