/*
    author twitter : @mmcgrana
   	https://gobyexample.com/signals
	Go by Example: signals

	때때로 우리는 Go프로그램이 지능적으로 Unix시그널들을 처리하기를 원한다.
	예를 들면, 우리는 SIGTERM이나, 입력처리를 그만두게하는 커맨드라인툴을 받았을 때 서버가 우아하게 꺼지길 원한다.
	(when it receives a `SIGTERM`, or a command-line tool to stop processing input if it receives a `SIGINT`.)
	여기에 Go에서 채널과 함께 어떻게 그런 신호를 다루는지 나와있다.
*/

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	// Go 신호 알림은 'os.Signal'값을 채널에서 전송함으로써 동작한다.
	// 우리는 저러한 알림들을 받을 채널을 만들 것이다.
	// (우리는 또한 프로그램이 종료될 수 있을 때 우리에게 알릴 채널또한 만들 것이다)
	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 'signal.Notify' 는 명시된 시그널알림들을 수신할 채널을 등록한다.
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 이 고루틴은 시그널을 위해 블락킹수신을 실행한다.
	// 이것이 수신을 할때 이것은 그것을 print하고 그리고 프로그램에게 마칠 수 있단 것을 알린다.
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 여기서 프로그램이 예상된 시그널(위의 done에게 값을 보내는 고루틴이 보내는)을 받을 때까지 기다린다.
	//그리고 종료된다.
	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
