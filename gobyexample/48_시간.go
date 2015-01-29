/*
    author twitter : @mmcgrana
    https://gobyexample.com/time
	Go by Example: time

	Go는 광범위한 시간과 간격(시간차)에 관한 지원을 제공하며, 여기 예제가 있다.
*/

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	//우리는 현재 시간을 얻음으로써 시작한다.
	now := time.Now()
	p(now)

	// 우리는 time구조체를 년, 월, 날짜 등등을 제공하면서 빌드할 수 있으며,
	// 시간은 언제나 위치(time zons)과 연관된다.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	//당신은 또한 예상한대로 시간값에서 다양한 컴포넌트들을 추출할 수 있다.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 월요일-일요일 `Weekday` 또한 사용가능하다.
	// The Monday-Sunday `Weekday` is also available.
	p(then.Weekday())

	// 두 시간대를 비교하는 이 메소드들이 있는데, 첫번째것은 전에 발생, 후에 발생했는지 테스트하고
	// Equals은 동시에 발생했는지 각각 테스트한다.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// `Sub`메소드는 두 시간대 사이의 간격을 나타내는 `Duration`을 리턴한다.
	diff := now.Sub(then)
	p(diff)

	// 우리는 간격의 차이를 다양한 단위로 계산할 수 있다
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 당신은 Add를 사용할 수 있다. 주어진 시간만큼 시간을 당길 수 있다.
	// 아니면 - 를 줘서 시간을 이전으로 당길 수 있다.
	p(then.Add(diff))
	p(then.Add(-diff))
}
