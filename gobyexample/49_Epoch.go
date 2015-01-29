/*
    author twitter : @mmcgrana
    https://gobyexample.com/epoch
	Go by Example: epoch

	프로그램에서 보통 요구사항은 초,밀리세컨, 나노세컨의 개수(the numbers)를 얻는 것이다.
	[Unix epoch](http://en.wikipedia.org/wiki/Unix_time) 때문이다.
	(역주: 역주는 epoch가 익숙하지가 않다)
	여기 Go에서 어떻게 하는지 예제가 있다.
*/

package main

import "fmt"
import "time"

func main() {

	// time.Now를 Unix나 UnixNano와 같이 사용하여서 경과시간을 얻을 수 있다.
	// Unix Epoch 가 초나 나노초에 각각 있기 때문이다?
	// since the Unix epoch in seconds or nanoseconds, respectively.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// `UnixMillis`가 없다는 것을 명심하자. 그러니 밀리세컨초를 얻자. since 이어서...
	//(죄송합니다;)밤인데 졸리네요; 이것만 하면 오늘 끝인지라^^.. 역주..
	// since epoch you'll need to manually divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	//당신은 또한 정수초나 나노세컨초로 변환할 수 있다.
	//왜냐하면 epoch는 'time'에 상응할 수 있기 때문이다.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
