/*
    author twitter : @mmcgrana
	https://gobyexample.com/tickers
    Go by Example: tickers

	timer는 당신이 미래에 어떤 일을 하고 싶을 때 쓰고 ticker는 당신이 정기적으로
	어떤 일을 반복적으로 하기 원할 때 쓴다. 여기 우리가 멈출 때까지 정기적으로 tick하는
	ticker 예제가 있다.  (틱장애가 떠오른다...역주)

*/

package main

import "time"
import "fmt"

func main() {

	// Ticker는 타이머(값을 전송하는 채널)와 비슷한 메커니즘이다.
	// 여기 채널에서 반복하는 range문을 이용하여서, 매초 500ms 마다 들어오는 값을 반복하자.
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	//Tickers는 타이머같이 멈춰질수 있는데 ticker가 멈춰지면, 그것은 채널로부터 어떠한 더 이상의 값도 받지 못한다.
	// 1500ms초후에 우리는 이것을 멈춰보자.
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
