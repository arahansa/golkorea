package main

/*
원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo

본글을 조금 내용을 요약해서 적어보겠다.

다음과 같이 주어진 선언이 있다면.. 밑에 것을 자동으로 정렬해서 바꿔줄 것이다.(예)

	type T struct {
       name string // name of the object
       value int // its value
   }
gofmt tool 을 이용해서 gofmt 파일이름 를 실행시키면 아마 정렬이 된 것을 볼 것이다.
*/
type T struct {
	name  string // name of the object
	value int    // its value
}

/*
	들여쓰기
	줄길이
	괄호구문등을 바꿔줄 것이다.


	들여쓰기(Indentation)
들여쓰기를 위해서 tab을 사용합니다. gofmt 는 tab을 default 들여쓰기로 변환해줄 것입니다. 꼭 그래야 할 경우만 space를 사용하세요.

	줄 길이(Line Length)
Go 언어는 줄 길이의 제한이 없습니다. 오버플로우에 대해서는 걱정할 필요가 없습니다. 만약 줄이 너무 길다면 줄을 tab으로 둘러싸면 됩니다.

	괄호구문(Parentheses)
Go 언어는 아주 적은 수의 괄호구문만을 필요로 합니다: 제어문(if, for, switch)은 Go언어의 문법상 괄호구문을 필요로 하지 않습니다. 또한, 연산자 우선순위는 더 짧고 더 간단합니다.
   즉,

   x<<8 + y<<16

   위의 수식에서 space가 우선순위를 암시하는 것을 말함.
   (역주:괄호를 사용하지 않고도 space만으로 우선순위가 나뉜다는 의미임.)
*/
