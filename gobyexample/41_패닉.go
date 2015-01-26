/*
    author twitter : @mmcgrana
    https://gobyexample.com/panic

	Go by Example: panic

	패닉은 일반적으로 어떤 것인가가 예상치못하게 잘못되었단 것을 의미한다.
	대부분 보통의 작업동안 발생하지 않아야하거나, 우아하게 다룰(handle) 준비가 안 된 에러가 발생했을 때
	빠르게 실패(fail)시키기 위해 패닉을 사용한다.

*/

package main

import "os"

func main() {

	//우리는 패닉을 이 곳(site)에서 예상치 못한 에러를 체크하기 위해 사용할 것이다.
	// 이 곳은 패닉을 발생시키기 위한 위치(site)이다.
	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.
	panic("a problem")

	//  패닉을 보통 사용하는 방법은 함수가 우리가 어떻게 처리할지 모르거나 (혹은 아예 모르는) 에러값을 리턴했을 때 그것을
	// 무시하는 것이다.
	//여기 새로운 파일을 생성하면서 익숙하지 않은 예외를 얻었을 경우 panic 을 사용하는 예제가 있다.

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	//이 프로그램을 실행하면 패닉에 이를 것이고 에러 메시지와 고루틴 트레이스를 출력할 것이다.
	// 그리고 non-zero status와 함께 프로그램이 종료될 것이다.
	// 많은 에러를 예외를 사용해서 처리하는 다른 언어와 달리 Go에서는 가능하면 에러를 리턴값에서 받는 것이 흔하다.
}
