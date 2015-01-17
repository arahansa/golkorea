/*
    author twitter : @mmcgrana
    https://gobyexample.com/closures

    Go by Example: Closures

 Go는 익명함수를 지원한다. [_anonymous functions_](http://en.wikipedia.org/wiki/Anonymous_function),
 이것은 클로저로부터 유래되었다. http://en.wikipedia.org/wiki/Closure_(computer_science)
 익명함수는 당신이 이름을 짓지않고 inline으로 함수를 선언할 때 유용하다.
*/
package main

import "fmt"

// 이 intSeq함수는  intSeq안에 익명으로 정의한 다른 함수를 리턴한다.
// 리턴된 함수는 변수 i 위에서 클로져를 생성하기  위해 닫혀진다? (원문 참고)
//The returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	//우리는 intSeq 이라는 것을 호출하고 결과를 (함수) nextInt 에 할당하였다.
	//이 함수는 자신이 가진 i 값을 가지고서 nextInt 가 호출될때마다 그것을 수정한다.
	nextInt := intSeq()

	// nextInt를 몇번 호출함으로써  클로저의 효과를 보시라~
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 상태가 특정 함수에서 유일하다는 것을 확인하자. 새로운 하나를 만들어 테스트해보자.
	newInts := intSeq()
	fmt.Println(newInts())
}
