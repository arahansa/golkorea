/*
    author twitter : @mmcgrana
    https://gobyexample.com/atomic-counters

	Go by Example: Atomic Counters

	Go에서 상태를 관리하는 주요 메커니즘은 채널을 통해 커뮤니케이션하는 것이다.
	우리는 예제 worker-pool를 통해서 이러한 것들을 보았다.
	그러나 다른 몇가지 옵션들이 있는데, 여기서 sync/atomic 패키지를 사용하는
	멀티플고루틴에 의해 접근되는 atomic 카운터를 한번 살펴보겠다.
*/

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

	//우리의 카운터를 위하여  unsigned integer 를 사용하겠다.
	var ops uint64 = 0

	//동시적 업데이터를 가정하기 위해서, 우리는 50개의 고루틴을 시작할 것이고
	//우리의 고루틴들은 1밀리세컨드 당 카운터를 각각 증가시킬 것이다.
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 원자적으로 카운터를 증가시키기 위해 우리는 AddUint64를 쓸 것이며 ops의 메모리주소를 &와 함께 줄 것이다.
				atomic.AddUint64(&ops, 1)

				// 다른 고루틴들이 처리되도록 허용한다.
				runtime.Gosched()
			}
		}()
	}

	// 몇몇 ops 가 모여지기를 허용하기 위해 1초 기다린다.
	time.Sleep(time.Second)

	// 다른 고루틴에 의해 업데이트 될 동안 안전하게 카운터를 사용하기 위해,
	// LoadUint64를 통하여 현재값의 복사본을 opsFinal로 추출한다.
	// 위에서 우리는 이 함수에게 값을 전송할 &ops메모리 주소를 주었다.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
