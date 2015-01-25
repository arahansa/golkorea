/*
    author twitter : @mmcgrana
    https://gobyexample.com/mutexes

	Go by Example: mutexes

	우리의 지난 예제에서 우리는 어떻게 간단한 카운터상태를 atomic운영을 이용하여 관리하는지 알아보았다.
	더 복잡한 상태를 위해서는, 멀티플고루틴에서 안전하게 데이터에 접근하기 위한 mutex를 쓴다.
*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	//우리의 예제에서 state는 map이 될 것이다.
	var state = make(map[int]int)

	//이 'mutex'는 상태에의 접근을 동기화할 것이다.
	var mutex = &sync.Mutex{}

	// 나중에 볼 다른 것과 mutex-based접근을 비교하기 위해
	//'ops'는 상태(state)에 관하여, 나중에 얼마나 많은 작업을 우리가 수행할지 카운트해줄 것이다.
	var ops int64 = 0

	// 여기 우리는 100개의 고루틴을 실행시켜서 반복된 상태(state)조회를 실행시킬 것이다.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// 각각의 조회를 위해 우리는 접근을 위한 key를 하나 집을 것이다.
				// mutex의 Lock() 은 상태(state)에 관한 독점적인 접근을 보장한다.
				// 주어진 키에 있는 값을 읽고 mutex의 Unlock()을 하여 ops카운트를 증가시킨다.

				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 이 고루틴이 스케쥴러를 굶어죽이지않는다는 것(원문:doesn't starve)을 보장하기 위해
				// 우리는 runtime.Gosched()의 각각의 작업 후에 명백하게 넘겨줘야(yield)한다.
				//이러한 넘겨줌은 자동으로 다뤄진다. 예를 들자면 모든 채널작업과 time.Sleep 같은 블로킹호출같은 것말이다.
				//그러나 이러한 경우 우리는 수동으로 해줄 필요가 있다.
				runtime.Gosched()
			}
		}()
	}

	// 우리는 또한 10개의 고루틴을 실행시켜서 글쓰기를 가정해본다.
	// 우리가 조회를 했었을 때와 같은 패턴을 사용할 것이다.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	// 10개의 고루틴이 상태(state)에서 일을 하게 하고
	// Let the 10 goroutines work on the `state` and
	// `mutex` for a second.
	time.Sleep(time.Second)

	// 마지막 작업 카운트 값을 받아 보고한다.
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 마지막 상태(state)를 락걸면서, 어떻게 이것이 끝나는지 보여준다.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
	// 프로그램실행은 우리가 mutex동기화된 state에 대하여 3,500,000작업을 실행했다는 것을 보여준다.
	//다음은 우리는 이와 같은 상태관리를 오직 고루틴과 채널을 이용하여 구현해볼 것이다.
}
