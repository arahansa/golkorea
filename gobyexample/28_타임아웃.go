/*
    author twitter : @mmcgrana
	https://gobyexample.com/timeouts
    Go by Example: timeouts

	Timeout은 프로그램이 외부자원에 접근하거나 실행시간을 bound하기 위해 중요하다.
	채널과 셀렉트 덕분에 Go에서 timeout을 구현하는 것은 쉽고 우아하다.

*/

package main

import "time"
import "fmt"

func main() {

	//예제를 위해 2촛후에 c1 채널에 그것의 값을 리턴하는 외부호출을 실행하는 것으로 가정한다.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 여기 timeout을 구현하는 select문이 있다.
	// res:= <-c1 은 값을 기다리고, <-Time.After 는 1초후에 값이 전송되기를 기다린다.
	// select 가 첫번째 receive와 함께 진행되기 때문에,
	// 프로그램이 허용된 1초를 넘을 경우 우리는 timeoutcase를 실행할 수 있다.
	// (역주 : 기다리기도하면서 타임아웃도 잰다는 뜻같다)
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	//만약 우리가 더 긴시간인 3초의 타임아웃을 허용할 경우는 c2로부터 값을 받는 것은 성공할 것이고.
	//그 결과를 출력할 것이다.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

// todo: cancellation?
