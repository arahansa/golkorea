/*
    author twitter : @mmcgrana
	https://gobyexample.com/closing-channels
    Go by Example: closing-channels

	채널을 닫는 것은 더이상 어떤 값도 그것으로 보내지지 않을 것이란 것을 가리킨다.
	이것은 채널의 리시버와 통신은 완료하는데 유용하다.
	여기서 worker는 따로 일하는 고루틴을 말하는 듯하다(역주)
*/

package main

import "fmt"

// 이 예제에서, 우리는 jobs 채널을 이용하여 main() 고루틴으로부터 work고루틴으로 작업이 완료되었다고
// 통신할 것이다. worker로 할일이 더 이상 없다면 우리는 jobs채널을 닫을 것이다.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//여기예 worker고루틴이 있다. 이것은 반복적으로 jobs로부터 j를, more는 jobs를 수신한다.
	// 이 특별할 2개의 값수신을 하면서 모든 작업을 완료되어 jobs가 close됬을 때
	//more의 값은 false가 될 것이다.
	//우리는 우리의 모든 작업을 했을 때, done'채널에 이것을 알리기 위해 이것을 사용한다.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 이것은 세개의 jobs를 jobs채널을 이용하여 worker에게 전송하고 jobs채널을 닫는다.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	//worker가 이전에 본 [synchronization](channel-synchronization)접근을 이용하여
	//기다리는 것을 볼 수 있다.
	<-done
}
