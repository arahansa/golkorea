// 36_연습_슬라이스
package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}

/*
Pic이라는 함수를 구현합니다. 이 함수는 dy개 만큼의 길이를 가지는 슬라이스를 리턴해야 하는데, 각각의 요소들은 또한 dx 개의 8비트 부호없는 8비트 정수 타입을 가지는 슬라이스입니다. 프로그램을 실행하면 이 정수값들을 흑백 (사실은 파란색)을 나타내는 값으로 해석하여 그림을 보여줄 것입니다.

그림은 여러분이 원하는 것으로 선택할 수 있습니다. (이용할 수 있는) 흥미로운 함수로는 x^y, (x+y)/2, x*y 등이 있습니다.

(여러분은 [][]uint8 슬라이스 내에서 사용할 각각의 []uint8 슬라이스를 할당하기 위해 루프를 활용해야 할 것입니다.)

(타입 간의 변환을 위해서는 uint8(intValue)을 사용합니다.)

Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
