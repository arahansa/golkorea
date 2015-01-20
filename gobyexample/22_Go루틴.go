/*
    author twitter : @mmcgrana
    https://gobyexample.com/goroutines

    Go by Example: Goroutines

  Go루틴은 실행가능한(of execution) 경량화쓰레드다.
*/

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	//우리가 f(s)라는 함수를 가지고 있다고 생각해보자.
	// 우리는 저 함수를 동시에 일어나게(synchronously) 실행하며 보통 방법으로 실행시킬 것이다.
	f("direct")

	// 이 함수를 Go 루틴에서 적용(invoke)하기 위해, go f(s)를 사용하자.
	// 이 새로운 Go루틴은 함수를 호출하면서 동시에(concurrently) 시작될 것이다.
	// 역주) new Thread 같은 개념인 것같다고 말하고 싶다.
	go f("goroutine")

	// 당신은 또한 익명함수호출로 Go루틴을 시작할 수 있다.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//우리의 두 함수호출은 지금 다른 Go루틴속에서 비동기적으로 실행될 것이다.
	// so execution falls through to here.
	// 이 Scanln 은 메시지를 받아서 사용자가 키보드값을 입력해야 프로그램이 종료됨.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
