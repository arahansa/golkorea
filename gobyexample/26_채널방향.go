/*
    author twitter : @mmcgrana
    https://gobyexample.com/channel-directions
    Go by Example: channel-directions

	만약 채널을 함수의 파라미터로 이용할때, 당신은 채널이 오직 값을 받거나, 수신할지 명시할 수 있다.
	이러한 명시는 프로그램의 타입안전(type-safety)를 증가시킨다.

*/

package main

import "fmt"

//이 핑 함수는 오직 값을 보내는 것만 받는다,
//이 채널로 받기를 시도하려고 한다면 이것은 compile 에러를 낼 것이다.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 이 퐁 함수는 오직 (pings)수신을 위한 채널과 ('pongs')전송을 위한 채널만을 받아들인다.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
