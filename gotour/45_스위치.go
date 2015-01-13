// 45_스위치
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

/*
다른 일반적인 언어를 아는 분이라면 switch 에 대해서 잘 알 것입니다.

다른 언어와 다른점은 case의 코드 실행을 마치면 알아서 break를 한다는 점입니다.

( fallthrough 로 끝나는 case는 스스로 break를 하지 않습니다 )


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
