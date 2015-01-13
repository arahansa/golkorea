// 63_고루틴
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
_고루틴_은 Go 런타임에 의해 관리되는 경량 쓰레드입니다.

go f(x, y, z)
위의 코드는 새로운 고루틴을 시작시킵니다.

f(x, y, z)
현재의 고루틴에서 f , x , y , z 가 평가(evaluation)되고,
새로운 고루틴에서 f 가 수행(execution)됩니다.

고루틴은 동일한 주소 공간에서 실행되므로, 공유되는 자원으로의 접근은 반드시 동기화 되어야 합니다.
[[http://golang.org/pkg/sync/][sync]] 패키지가 이를 위해 유용한 기본 기능을 제공합니다.
Go 에서는 그외에도 다양한 기본 기능을 제공하니 크게 필요치 않을 테지만요. (다음 슬라이드를 보세요.)

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
