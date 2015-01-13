// 66_Range와Close
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
데이터 송신측은 더이상 보낼 값이 없다는 것을 알리기 위해 채널을 close 할 수 있습니다.
수신측은 다음과 같이 수신 코드에 두번째 인자를 줌으로써 채널이 닫혔는지 테스트 할 수 있습니다.

v, ok := <-ch
채널이 이미 닫혔고 더이상 받을 값이 없다면 ok 는 false 가 됩니다.

for i := range c 반복문은 채널이 닫힐 때까지 계속해서 값을 받습니다.

주의: 송신측만 채널을 닫을 수 있습니다. 수신측에선 불가능합니다.
이미 닫힌 채널에 데이터를 보내면 패닉이 일어납니다.

또하나의 주의: 채널은 파일과 다릅니다; 항상 닫을 필요는 없습니다.
채널을 닫는 행위는 오로지 수신측에게 더이상 보낼 값이 없다고 말해야 할때만 행해지면 됩니다.
range 루프를 종료시켜야 할 때처럼요.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
