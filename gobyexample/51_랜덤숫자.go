/*
    author twitter : @mmcgrana
    https://gobyexample.com/random-numbers
	Go by Example: random-numbers

	Go의 math/rand 패키지는 수도랜덤숫자 일반화를 제공한다.
 	[pseudorandom number](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)
*/

package main

import "fmt"
import "math/rand"

func main() {

	//예를 들자면, `rand.Intn`은 랜덤정수 n을 리턴한다. n은 이 밑에서는 0이상100미만 값이다.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// randFloat64는 float64 타입의 f를 리턴한다. 여기선 0.0 <=f < 1.0 값이다.
	fmt.Println(rand.Float64())

	// 이것은 랜덤 실수를 다른 범위에서 생성해본다. 예를 들자면 `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// deterministic인 수도랜덤 생성기(pseudorandom generator)를 만들기 위해, 잘 알려진 seed를 주자.
	// 역주: 번역이 매끄럽지 못해 원문을 남깁니다. 좋은 번역 제시해주시면 감사하겠구요.
	// 그냥 어떤 값을 넣어서 숫자를 뽑아내는 것같습니다~
	// To make the pseudorandom generator deterministic,
	// give it a well-known seed.
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	// `rand.Source`은 rand패키지의 함수같은 결과를 호출한다.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//만약 당신이 같은 숫자를 가지고서 랜덤을 돌리면
	//그것은 같은 랜덤숫자의 같은 나열을 만들 것이다.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
}
