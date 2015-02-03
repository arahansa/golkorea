/*
    author twitter : @mmcgrana
    https://gobyexample.com/url-parsing
	Go by Example: url-parsing

	url은 자원(resource)들을 위치(locate)시키는 한결같은 방법을 제공한다.
	여기 Go에서 URL을 어떻게 파싱하는지 나와있다.
*/
package main

import "fmt"
import "net/url"
import "strings"

func main() {

	//우리는 스키마, 인증정보, 호스트, 포스,경로 , 쿼리 파라미터와 쿼리조각등을 가지고 있는, 이 예제 URL을 파싱할 것이다.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// URL을 파싱하고 에러가 없단 것을 보장한다.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 스키마에 접근하는 것은 간단하다.
	fmt.Println(u.Scheme)

	// User는 모든 인증정보를 포함한다.
	// 여기서 개인적인 값을 위해 Username과 Password 를 호출한다.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host 는 호스트네임과 포트를 포함하고 만약 존재하면 수동으로 포트와 호스트를 추출해야한다. (배열씀)
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	// 여기 우리는 경로와 조각을 #이후로 추출한다.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// `k=v`형식의 문자열에서 쿼리 파라미터를 얻기 위해
	// RawQuery를 사용한다. 당신은 또한 쿼리파라미터들을 맵으로 파싱한다.
	// 문자열슬라이스나 문자열로부터 파싱된 쿼리파라미터맵으로 파싱되며 만약 첫번째 값을 얻고 싶으면 index값은 [0]이다.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
