// 52_포인터리시버를가지는메소드
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}

/*
메소드는 이름이 있는 타입 또는 이름이 있는 타입의 포인터와 연결할 수 있습니다.
방금 두 개의 Abs 메소드를 보았는데, 하나는 *Vertex 라는 포인터 타입의 메소드고,
다른 하나는 MyFloat 값 타입의 메소드 입니다.
포인터 리시버를 사용하는 이유는 두 가지 입니다. 첫째, 메소드가 호출될 때 마다 값이
복사되는 것(큰 구조체 타입인 경우 값이 복사되는 것은 비효율적이죠)을 방지하기 위함 입니다.
다른 이유는 메소드에서 리시버 포인터가 가르키는 값을 수정하기 위함 입니다.

*Vertex 타입의 리시버 대신 Vertex 를 사용하도록 메소드 Abs 와 Scale 의 선언부분을 바꿔 보세요.
v 를 Vertex 타입으로 받으면 Scale 메소드가 더 이상 동작하지 않습니다. Scale 은 v 를 바꾸는데,
v 가 (포인터가 아닌) 값 타입이기 때문에 Vertex 타입인 복사본에 작업을 하기 때문에 원래의 값은 바뀌지 않습니다.
Abs 의 경우는 다릅니다. 여기서는 v 를 읽기만 하기 때문에,
(포인터가 가르키는) 원래의 값이건 복사본이건 상관이 없게 됩니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
