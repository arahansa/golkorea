/*
    author twitter : @mmcgrana
	https://gobyexample.com/non-blocking-channel-operations
    Go by Example: non-blocking-channel-operations

	채널에서 기초적인 전송과 수신은 블록킹이다.
	그러나, 우리는 select를 'default' 문과 같이 써서, 논블록킹 전송과 수신,
	심지어 논블록킹 멀티 select를 구현할 수 있다.

*/

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 여기 논블록킹 수신이 있다. 만약 값이 message에서 사용가능하면
	// 셀렉트는 <-message 케이스를 값과 함께 선택할것이지만 그렇지 않다면
	//이것은 즉시 default 를 선택할 것이다.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 논블록킹 전송은 비슷하게 작동한다.
	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//우리는 멀티웨이 논블록킹 셀렉트를 구현하기 위해
	// default 위에 멀티플 case를 사용할 수 있다.
	//여기에선 message와 signals 둘다에서 논블록킹 수신을 시도한다.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
