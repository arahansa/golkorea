// 59_이미지
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

/*
Package image 는 Image 인터페이스를 정의합니다.

package image

type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}
(모든 세부사항에 대한 것은 이 문서 를 참고하십시오.)
http://golang.org/pkg/image/#Image
또한, color.Color 와 color.Model 는 인터페이스이지만,
미리 정의된 구현체인 color.RGBA 와 color.RGBAModel 을
사용함으로써 그 인터페이스를 무시할 수 있습니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
