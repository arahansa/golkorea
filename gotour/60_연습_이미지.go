// 60_연습_이미지
package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
)

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

/*
이전의 연습에서 당신이 작성한 그림 생성기를 기억하십니까? 다른 생성기를 만들어봅시다.
하지만 이번에는 데이터의 슬라이스 대신에 image.Image 의 구현체를 반환할 것입니다.
당신 자신의 Image 타입을 정의하시고, 필수 함수들 을 구현하신 다음,
pic.ShowImage 를 호출하십시오.
Bounds 는 image.Rect(0, 0, w, h) 와 같은 image.Rectangle 을 반환해야 합니다.
ColorModel 은 color.RGBAModel 을 반환해야 합니다.

At 은 하나의 컬러를 반환해야 합니다;
지난 그림 생성기에서 값 v 는 color.RGBA{v, v, 255, 255} 와 같습니다.

Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
