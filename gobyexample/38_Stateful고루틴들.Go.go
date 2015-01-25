/*
    author twitter : @mmcgrana
    https://gobyexample.com/stateful-goroutines

	Go by Example: stateful-goroutines

	지난 예제에서 mutexes로 명백한 락킹을 하여, 멀티플고루틴에서 공유된 상태에대한 동기화된 접근을 하였다.
	또다른 선택은 내장된 고루틴과 채널의 동기화특성을 이용하여 같은 결과를 얻는 것이다.
	이 채널기반 접근법은 정확히 1개의 고루틴들이 가진 각각의 데이터 조각들을 서로 통신하고 가짐으로써
	Go의 메모리공유 목적(ideas)들을 조정한다.
	(번역이 잘 됬는지 모르겠다. 원래 그냥 소스를 보고 실행되는 모습을 봐야할것같다;;:역주)
*/

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 이 예제에서 우리의 상태는 단일 고루틴에 의해 소유될 것이다.
// 이것은 데이터가 동시성 접근에 의한 방해를 받지 않을 것이란 것을 보장한다.
// 상태를 읽거나 쓰기 위해, 다른 고루틴들은 소유하고 있는 고루틴에게 메시지를 보낼 것이며,
// 그에 상응하는 답변을 받을 것이다.
// 이러한 'readOp'과 'writeOp' 구조체는 이러한 요청들과
// 소유하고있는 고루틴이 응답할 방법을 요약(encapsulate)한다.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 전에서 우리는 얼마나 많은 작업이 수행되는지 카운트 했었다.
	var ops int64 = 0

	// reads 와 writes채널은 읽고 쓰는 요청을 각각 발행하기 위해 ,
	// 다른 고루틴들에 의해 사용될 것이다.
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 여기 상태를 소유하고있는 고루틴이 있으며, 이것(상태)은 예전의 예제와 같이 맵이지만,
	// 지금은 stateful 고루틴에 private 하다.
	// 이 고루틴은 반복적으로 reads, writes채널에서 selects하여, 요청이 도착하면 응답한다.
	// 응답은 첫번째 요청작업을 수행하며 실행되고 그후 응답채널 resp 에 성공했다고 알리는 값을 전송한다.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 이것은 100개의 고루틴을 실행하여서 reads채널을 통해 state-owning고루틴에게 reads를 발행한다.
	// 각각의 read는 readOp 를 만들어, reads채널로 보내지길 운한다.  그리고
	// 제공된 resp채널을 통해 결과를 받기를 원한다.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 우리는 10개의 writes를 또한 시작하며, 비슷한 접근을 사용할 것이다.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 1초동안 고루틴들이 일하게 하자
	time.Sleep(time.Second)

	//마지막으로, ops카운트를 캡쳐하고 보고하게 하자.

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	//우리의 프로그램을 실행하는 것은 고루틴기반한 관리예제로 초당 800,000작업을 수행했다.
	// 이 특별한 케이스에서 고루틴기반접근은 뮤텍스기반접근보다 좀 더 관여한다.(밑에 원문 남김)
	//For this particular case the goroutine-based approach was a bit more involved than the mutex-based one

	// 그러나 특정한 케이스에서 이것은 좀 더 유용할 것이다. 예를 들자면 다른 채널을 가진 곳에서나
	// mutexes들이 에러를 발생하려는 경향이 있는 멀티플한 것을 관리할 때 말이다.
	//당신은 어떠한 접근이건, 좀 더 자연스러운 접근을 해야하며 특히 당신의 프로그램에 올바른 것이 무엇인지 이해하는 것이 중요하다

}
