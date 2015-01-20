/*
    author twitter : @mmcgrana
    https://gobyexample.com/channel-synchronization
    Go by Example: channel-synchronization

	우리는 채널을 Go루틴들사이에서 동기화실행을 위해 사용할 수 있다.
	여기에 Go루틴이 마쳐질때까지 기다리기위해 리시브를 블락킹하는 예제가 있다.
*/

package main

import "fmt"
import "time"

// 여기 이 함수는 Go루틴에서 실행될 함수인데,
// 다른 고루틴에게 이 함수의 작업이 완료되었단 것을 알리기 위해 done 채널이 사용될 것이다.
func worker(done chan bool) {
	fmt.Println("working...(역주:여기서 1초 쉴겁니다)")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 우리가 다 됬단 것을 알려주기 위한 값을 전송한다.
	done <- true
}

func main() {

	// worker Go루틴을 실행하고 알림을 주기위해 채널을 전달한다.
	done := make(chan bool, 1)
	go worker(done)

	//우리가 채널이 있는 worker로부터 알림을 받을 때까지 블락된다.
	<-done
}
