/*
	author twitter : @mmcgrana
	https://gobyexample.com/constants

	Go by Example: Constants
*/
package main

import "fmt"
import "math"

// `const` 는 상수를 선언한다.
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` 는 var가 나타날 수 있는 어디든지 등장할 수 있다.
	const n = 500000000

	// 상수표현식은 임의적인 정확도의 연산을 수행한다.
	const d = 3e20 / n
	fmt.Println(d)

	// 숫자상수는 명시적형변환같은 타입이 주어질때까지 어떤 타입도 갖지 않는다.
	fmt.Println(int64(d))

	//숫자는 그것을 필요로 하는 맥락속에서 변수할당이나 함수호출로써 사용되므로써 타입이 주어질 수 있다
	//예를 들자면 여기서 `math.Sin` 은 `float64` 를 필요로 한다.
	fmt.Println(math.Sin(n))
}
