/*
    author twitter : @mmcgrana
    https://gobyexample.com/structs

    Go by Example: Structs

 Go의 구조체는 필드의 모음형식입니다.
 구조체는 기록을 형성하기 위하여 데이터들을 그룹핑하는 데 유용합니다.
 당신의 프로그램안에서 값과 기록에게 참조를 전달할 수 있다.

*/

package main

import "fmt"

// 이 person 구조체는 name 과 age 필드를 가집니다.
type person struct {
	name string
	age  int
}

func main() {

	//이 문법은 새로운 구조체를 생성합니다.
	fmt.Println(person{"Bob", 20})

	// 구조체를 초기화할 때 당신은 필드명을 사용할 수 있습니다.
	fmt.Println(person{name: "Alice", age: 30})

	// 생략된 필드는 제로값이 됩니다.
	fmt.Println(person{name: "Fred"})

	// `&` 접두어는 포인터를 구조체에게 넘겨준다. (&prefix yields a pointer to the struct)
	fmt.Println(&person{name: "Ann", age: 40})

	//.과 함께 구조체 필드에 접근한다.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// .을 또한 구조체 포인터와 사용할 수 있다. 포인터는 자동으로 참조되게된다.(dereference)
	sp := &s
	fmt.Println(sp.age)

	// 구조체는 변경 가능하다.
	sp.age = 51
	fmt.Println(sp.age)
}
