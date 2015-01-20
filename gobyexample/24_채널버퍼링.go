/*
    author twitter : @mmcgrana
    https://gobyexample.com/channels
    Go by Example: channels

  기본적 채널은 버퍼되지(unbuffered)않았다. 이것은 그들이 오직 채널들이
  그에 상응하는 받는 (`<- chan`) 이 보내는 값을 받을 준비가 되어야만, 전송(`chan <-`)하는 것을 수락한다는 것이다.
  Buffered channels 은 그에 상응하는 값들을 위한 리시버없이 제한된 개수의 값을 수락한다.
*/

package main

import "fmt"

func main() {

	// 여기서 make로 문자열채널을 만들고, 2개의 값까지 버퍼링 up 한다.
	messages := make(chan string, 2)

	//이 채널은 버퍼되었기 때문에, 우리는 이러한 값들을 채널에 보낼 수 있으며
	//이때는, 이에 상응하는 리시버가 없어도 된다.
	messages <- "buffered"
	messages <- "channel"

	//후에 우리는 평소대로 이러한 두가지 값을 받을 수 있다.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
