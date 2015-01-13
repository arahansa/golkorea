// 68_셀렉트의디폴트케이스
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

/*
select 의 default 케이스는 현재 수행 준비가 완료된 케이스가 없을 때 수행됩니다.

블로킹 없이(비동기적인) 송/수신을 하고자 할 때 default 케이스를 사용하세요.

select {
case i := <-c:
	// i를 사용
default:
	// c로부터의 수신은 블록된 상태
}

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
