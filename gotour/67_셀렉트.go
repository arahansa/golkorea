// 67_셀렉트
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
select 구문은 고루틴이 다수의 통신 동작으로부터 수행 준비를 기다릴 수 있게 합니다.

select 는 case 구문으로 받는 통신 동작들 중 하나가 수행될 수 있을 때까지 수행을 블록합니다.
다수의 채널이 동시에 준비되면 그 중 하나를 무작위로 선택합니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
