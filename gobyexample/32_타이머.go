/*
    author twitter : @mmcgrana
	https://gobyexample.com/timers
    Go by Example: timers

	우리는 나중에 특정 시점이나 어느정도의 시간이 흐른뒤에 반복적으로  Go코드를 실행시키기를 원할 것이다.
	Go의 내장된 timer와 ticker는 이러한 일들을 쉽게 해준다.
	처음에 timer를 알아보고 그다음에 ticker를 알아보자.
*/

package main

import "time"
import "fmt"

func main() {

	//타이머는 미래의 단일 이벤트를 나타낸다.
	// 당신은 타이머에게 얼마나 오랫동안 당신이 기다리는지 말할 수 있다.
	// 그리고 타이머는 그때에 알릴수 있는 채널을 제공한다.
	//이 타이머는 2초동안을 기다릴 것이다.

	timer1 := time.NewTimer(time.Second * 2)

	// <-timer1.C 는 타이머의 채널 C가 타이머가 종료되었다는 것을 가리키는 값을 보낼때까지 블락된다.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	//만약 네가 단지 기다리기를 원한다면, 당신은 아마 time.Sleep 을 사용할 수 있을 것이다.
	// 타이머가 유용한 또 하나의 이유는 타이머가 종료되기 전에 그것을 취소할 수 있다는 것이다.
	//여기 예제가 있다.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
