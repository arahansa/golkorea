/*
    author twitter : @mmcgrana
    https://gobyexample.com/errors

    Go by Example: Errors

 Go 에서는 명시적이고 분리된 리턴값으로 에러를 처리하는 것이 자연스럽다.
 이것은 	자바나 루비에서 사용되는 예외, 혹은 과적된 단일 결과와는 대조적이다.
에러값은 때때로 C에서 사용되었다. Go의 접근은 어떤 함수가 에러를 리턴하는지 보기 쉽게 하고,
에러를 같은언어구조를 사용하여 에러를 다루는 것이다.

*여긴 같은 언어구조에 대한 설명 : constructs employed for any other, non-error tasks.
same language 뒤에 바로 나온 말인데 무슨 말인지 그냥 패쓰함; 별로 안 중요한 것같기도하고;
요새 좀 바쁘다 ㅠ흙.

역주 개요: f1 과 f2를 만들어서 에러를 테스트 해본다.
f1은 errors.New로 에러를 만들어 리턴하고,
f2는 &접두어를 이용하여 새로운 구조체를 만들어 리턴한다.
*/
package main

import "errors"
import "fmt"

// 관습적으로, 에러는 마지막으로 리턴된 값이고. 내장된 인터페이스 타입 'error'를 가짐.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` 는 기초적인 'error'값을 주어진 메시지와 함께 생성한다.
		return -1, errors.New("can't work with 42")

	}
	// 에러 포지션에서 nil값은 에러가 없다는 것을 가리킨다.
	return arg + 3, nil
}

// 에러의 커스텀타입을 사용하는 것은 가능하며, Error() 메소드를 구현하면 된다.
// 여기에 명시적으로 아규먼트 에러를 나타내는 커스텀 타입을 사용하는 다양한 예제가 있다.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// 이러한 경우 우리는 &argError 문법을 사용하여서 새로운 구조체를 만든다.
		// 구조체는 두개의 필드 arg와 prob를 제공한다.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// 이 밑에 있는 두개의 반복문은 각각 에러-리턴 function을 test한다.
	// if문에서 인라인에러체크의 사용은 Go code에서 흔한 사용이라는 것을 알아두자.
	// if문에서 한줄로 e!=nil 이것 말하는 것임(역주)
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	// 만약 커스텀에러에서 프로그래밍적으로 데이터를 사용하고 싶다면
	// type assertion을 통해서 커스텀에러의 인스턴스를 에러로서 받기를 필요로 할 것이다.
	//역주 :  ae, ok := e.(*argError);이 type assertion인듯하다.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println("여기가 에러 출력부분(역주: 더 넣음) ")
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	// If you want to programmatically use the data in
	// a custom error, you'll need to get the error  as an
	// instance of the custom error type via type
	// assertion.
}
