/*
	author twitter : @mmcgrana
	https://gobyexample.com/maps

	Go by Example: Maps
	맵은 Go의 내장된 연관데이터타입이다(때론 다른 언어에서 해쉬나 dicts 로 불린다)
*/
package main

import "fmt"

func main() {

	// 빈 맵을 생성하기위해서
	// `make(map[키타입]값타입)`.
	m := make(map[string]int)

	// 키/값 쌍 값 저장은 다음과 같이 변수명[키]값 문법.
	m["k1"] = 7
	m["k2"] = 13

	// Println을 하면 맵의 모든 키/값이 나온다.
	fmt.Println("map:", m)

	// `name[key]`로 해당 키에 대한 값을 받는다.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// len 은 키/값 쌍의 갯수를나타낸다.
	fmt.Println("len:", len(m))

	//  `delete` 는 맵으로 부터 해당 키/값을 지운다.
	delete(m, "k2")
	fmt.Println("map:", m)

	// 맵으로 부터 다중 리턴을 받는 경우 두번째 값은 이 키에대해서 값이 있는지 없는 지의 여부이다.
	// 이 키(k2)에 대한 값이 당장 안쓰기 때문에 생략을 위해 '_'로 쓸 수 있다.
	// 이 방법을 사용하므로써, 0이나 "" 같은 zero-value인지 아니면 아예 값이 없는 것인지 구별할 수 있다.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 또한 한줄에 선언과 초기화를 할 수 있다.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
