/*
    author twitter : @mmcgrana
   	https://gobyexample.com/writing-files
	Go by Example: writing-files

	GO에서 파일을 쓰는 것은 이전에 본 읽기와 비슷하다.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//시작하기 위해, 여기 어떻게 문자열을 파일로 저장하는지 나타있다
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 더 다른 방법의 쓰기를 위해서, 쓰기를 위해 파일을 열어보겠다
	f, err := os.Create("/tmp/dat2")
	check(err)

	// 관용적으로 defer를 써서 파일을 연 후에 닫아주도록 한다.
	defer f.Close()

	// 당신이 예상한대로 바이트슬라이스를 Write한다.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString또한 사용가능하다.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// `Sync` 는 안정된 공간에 쓸 것을 플러쉬한다.
	f.Sync()

	// bufio는 이전에 본 버퍼된 리더같은 버퍼된 라이터를 제공한다.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	//Flush를 써서 버퍼된 작업들이 writer에 적용될 수 있도록 보장하자.
	w.Flush()

}
