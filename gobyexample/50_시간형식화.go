/*
    author twitter : @mmcgrana
    https://gobyexample.com/time-formatting-parsing
	Go by Example: time-formatting-parsing

	Go는 패턴기반한 레이아웃을 통해서 시간을 형식화하고 파싱하는 것을 지원한다.
*/

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// 여기 RFC3339를 따라서 시간을 형식화하는 예제가 있다.
	// 그에 상응하는 레이아웃 상수를 사용한다.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 시간을 파싱하는 것은 같은 레이아웃 값을 `Format`으로 이용한다.
	// Time parsing uses the same layout values as `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format`과 `Parse` 는 예제기반한 레이아웃을 사용한다.
	// 보통 당신은 이러한 레이아웃들을 위해 time에서 온 상수를 사용할 것이지만, 커스텀 레이아웃도 사용할 수 있다.
	// 레이아웃은 무조건 참조 시간(`Mon Jan 2 15:04:05 MST 2006`)을 사용하여 주어진 시간/문자열이 형식/파싱될 패턴을 보여줘야 한다.
	// 이 예제 시간은 정확히 다음과 같이 보여질 것이다.(시간은 2006 같이, 시간을 위한 15, 요일을 위한 Monday 같은 형식으로 말이다)
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 순수한 숫자 표현을 위해 당신은 또한 표준문자열형식을 시간값에서 추출한 요소들과 함께 사용할 수 있다.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` 는 잘못된 입력이 있을때 에러를 리턴하며 파싱문제를 설명할 것이다.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
