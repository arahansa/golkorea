/*
    author twitter : @mmcgrana
	https://gobyexample.com/select
    Go by Example: select

	Go의 셀렉트는 여러 채널 운영에서 당신을 기다리게 해준다.
	고루틴과 채널을 셀렉트와 조합은 Go의 파워풀한 특징이다.
*/

package main

import "time"
import "fmt"

func main() {

	// 우리의 예제를 위해 우리는 이 두 채널에서 셀렉트할 것이다.
	c1 := make(chan string)
	c2 := make(chan string)

	// 각각의 채널은 일정시간이후 값을 받을 것이다.
	// 예를 들자면 동시에 발생하는 고루틴에서 실행되는 원격프로시저 호출을 블락킹할 것이다.
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//우리는 동시에 이러한 값 둘다 기다리기 위해 select 를 사용할 것이며
	// 값이 도착하면 각각의 값들을 출력할 것이다.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
