//출처 : http://golang-examples.tumblr.com/post/44089080167/get-set-a-field-value-of-a-struct-using-reflection
package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	N int
}

//리플렉션을 이용한 게터세터를 이용할 때,
//겟할 때는 reflect.ValueOf(n)
//셋할 때는 Value(&변수명).Elem()
func main() {

	n := MyStruct{1}

	// 겟해보기
	immutable := reflect.ValueOf(n)
	val := immutable.FieldByName("N").Int()
	fmt.Printf("N=%d\n", val) // prints 1

	// set
	mutable := reflect.ValueOf(&n).Elem()
	mutable.FieldByName("N").SetInt(7)
	fmt.Printf("N=%d\n", n.N) // prints 7
}
