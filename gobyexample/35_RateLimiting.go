/*
    author twitter : @mmcgrana
	https://gobyexample.com/rate-limiting
    Go by Example: rate-limiting

	(http://en.wikipedia.org/wiki/Rate_limiting)
	RateLitming 은 리소스활용 컨트롤과 서비스퀄리티유지에 중요한 메커니즘이다.
	Go는 go루틴, 채널, tickers와 함께 우아하게 rate limiting을 지원한다.
*/

package main

import "time"
import "fmt"

func main() {

	// 먼저 기본적인 rate limiting을 보도록 하자.
	// 우리가 들어오는 요청 핸들링을 제한하고 싶다고 하자.
	// 우리는 이러한 요청들을 같은 이름의 채널에서 다루게?(serve off) 할 것이다?
	// We'll serve these requests off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//이 limiter채널은 매 200밀리초마다 값을 받을 것이다.
	//이것은 우리의 rate limiting 계획에서 조절장치이다.
	limiter := time.Tick(time.Millisecond * 200)

	// 각각의 요청을 다루기 전에, limiter채널에서 수신을 블락킹하기전
	//우리는 우리자신을 1개의 요청당 200밀리초로 제한할 것이다.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//  전체의 rate 한계를 보존하는동안, 우리는 그 rate 한계속에서 요청들의 short bursts를 허락할 것이다.
	// We may want to allow short bursts of requests in
	// our rate limiting scheme while preserving the
	// overall rate limit.
	// 우리는 우리의 리미터채널에 버퍼링함으로써 이것을 할 수 있다.
	//We can accomplish this by buffering our limiter channel.
	// 이 burstyLimiter 채널은 3개의 이벤트까지 bursts를 허락한다.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 매 200밀리초마다 우리는 새로운 값들을 burstyLimiter에 그것의 한계치인 3개까지 추가할 것이다.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 여기서 5개 이상의 요청을 가정해보자. 이 요청들에서 3번째것까지만 burstyLimiter 의 burst 수용능력안에 들어갈 것이다.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
