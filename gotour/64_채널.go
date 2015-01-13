// 64_채널
package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*
채널은 채널 연산자 <- 를 이용해 값을 주고 받을 수 있는, 타입이 존재하는 파이프입니다.

ch <- v    // v 를 ch로 보냅니다.
v := <-ch  // ch로부터 값을 받아서
           // v 로 넘깁니다.
(데이터가 화살표 방향에 따라 흐릅니다.)

맵이나 슬라이스처럼, 채널은 사용되기 전에 생성되어야 합니다:

ch := make(chan int)
기본적으로, 송/수신은 상대편이 준비될 때까지 블록됩니다.
이런 특성이 고루틴이 명시적인 락이나 조건 없이도 동기화 될 수 있도록 돕습니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
