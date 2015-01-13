// 56_연습_에러
package main

import (
	"fmt"
)

func Sqrt(f float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/*
당신의 Sqrt 함수를 이전 연습에서 복사하고 error 값을 반환하도록 수정하십시오.

Sqrt 함수는 복소수를 지원하지 않기 때문에, 음수가 주어지면 nil 이 아닌 에러 값을 반환해야 합니다.

새로운 타입을 만드십시오.

type ErrNegativeSqrt float64
and make it an error by giving it a 그리고 아래 메소드를 구현함으로써
그 타입이 error 가 되게 하십시오.

func (e ErrNegativeSqrt) Error() string
이는 ErrNegativeSqrt(-2).Error() 가 "cannot Sqrt negative number: -2" 를 반환하는 그러한 메소드입니다.

Note: Error 메소드 내에서 fmt.Print(e) 를 호출하면 이 프로그램을 무한루프에 빠질 것입니다.
e 를 바꿈으로써 이 문제를 피할 수 있습니다. 왜 그럴까요?

음수가 주어졌을 때 ErrNegativeSqrt 값을 반환하도록 당신의 Sqrt 함수를 바꾸십시오.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
