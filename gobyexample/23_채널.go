/*
    author twitter : @mmcgrana
    https://gobyexample.com/channels
    Go by Example: channels

  채널은 동시에 공존하는 Go루틴들을 연결하는 파이프다.
  하나의 go루틴에서 채널로 값을 전송할 수 있으며 그러한 값들을 다른 Go루틴을 통해 받을 수 있다.
*/

package main

import "fmt"

func main() {

	// `make(chan 값-타입)` 로 새로운 채널을 만든다.
	// 채널은 자신이 운반하는 값에 의해 타입이 정해진다.
	messages := make(chan string)

	// `channel <- 문법을 이용해서 채널에 값을 전송할 수 있다.
	// 여기 우리가 'message'라는 위에서 만든 채널에 "ping"을 보내는 작업을 Go루틴으로부터 하고 있다.
	go func() { messages <- "ping" }()

	// <-channel 문법은 채널로부터 값을 받는다.
	// 여기 우리는 위에서 보낸 ping 메시지를 받아서 그것을 출력한다.
	msg := <-messages
	fmt.Println(msg)
}
