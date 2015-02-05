/*
    author twitter : @mmcgrana
   	https://gobyexample.com/reading-files
	Go by Example: reading-files

	많은 Go프로그램들에서 기본적으로 파일을 읽고 쓰는 작업이 필요하다.
	먼저 우리는 파일을 읽는 몇가지 예제들을 살펴보겠다.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 파일을 읽는 것은 에러를 위한 검사를 먼저 필요로 한다.
// 이 helper 는 우리의 에러체크를 간소화하게 해줄 것이다.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 아마 가장 기본적인 파일 읽기 작업은 파일전체를 메모리로 읽어들이는 것이다.
	// 윈도우용일 경우를 따로 추가했다(역주)
	//dat, err := ioutil.ReadFile("/tmp/dat")
	dat, err := ioutil.ReadFile("c:/go/text.txt") //역주
	check(err)
	fmt.Print(string(dat))

	// 당신은 종종 파일의 어떤 부분을 어떻게 읽을지 컨트롤 하고 싶을 것이다.
	// 이러한 작업을 위해, 파일을 Open하여서 os.File 값을 얻도록 하자
	//f, err := os.Open("/tmp/dat")
	f, err := os.Open("c:/go/text.txt")
	check(err)

	// 파일의 시작으로부터 몇 바이트를 읽는다.
	// 5까지 읽는 것을 허용한다. 얼마나 실제로 읽어들이는지 보자.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 당신은 또한 'Seek' 을 써서 파일에서 그 부분부터 Read읽을 수 있다.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// io패키지는 파일읽는데 도움이 되는 몇가지 함수를 제공한다.
	// 예를 들자면, 여기 있는 것들 처럼 읽는 것은 ReadAtLeast 와 함께 구현되어서 좀 더 튼튼(?robustly)하다
	// The `io` package provides some functions that may
	// be helpful for file reading. For example, reads
	// like the ones above can be more robustly
	// implemented with `ReadAtLeast`.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 내장된 되돌리는 기능없지만, Seek(0, 0)은 이러한 일들을 하게 해준다.
	_, err = f.Seek(0, 0)
	check(err)

	// bufio 패키지는 버퍼된 리더를 구현한다.
	// 버퍼된 리더는 추가적으로 제공하는 읽기메소드와 그것의 효율성 덕분에 유용하다.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 모든 작업을 마치면 파일을 닫는다 (보통 이것은 Open되면 defer와 함께 예약된다.)
	f.Close()

}
