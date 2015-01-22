/*덧붙임(Append)

원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo


이제 내장함수인 append 함수에 대해 설명하도록 하겠습니다. 
append 함수의 특징은 위에 나온 기존의 append 함수와는 좀 다릅니다. 대략 다음과 같습니다:
*/

func append(slice []T, elements...T) []T
/*
T는 어떤 주어진 타입에 대한 위치표시(placeholder) 입니다. 
사실 Go언어에서 호출자에 의해 T 타입이 결정되어지는 함수는 사용할 수 없습니다. 
이 것이 append가 내장함수인 이유입니다.: 그것은 컴파일러의 지원이 필요합니다. 
append가 하는 것은 slice의 끝에 요소(element)를 추가하고 결과를 반환하는 것 입니다. 
결과는 반환되어질 필요가 있습니다. 왜냐하면, 우리가 손으로 써서 글에 덧붙이는 것 처럼(append), 
원본 배열이 변경될지도 모르기 때문입니다. 여기 간단한 예제가 있습니다.
*/

x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
/*
위의 예제는 [1 2 3 4 5 6]을 출력 합니다. 따라서, append는 조금은 Printf의 동작과도 비슷하며, 
임의의 개수의 인자들을 취합합니다. 하지만, 만약 slice와 slice를 append하기를 원한다면 
어떻게 해야할까요? 쉽습니다: 위의 예제에서 Output을 호출하듯이, 호출하는 쪽에서 ...을 사용하면
 됩니다. 이렇게 하는 것은 위의 예에서 한것과 동일한 결과를 출력합니다.
*/
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
/*
... 없이는 컴파일이 되지 않습니다. 왜냐하면 타입이 잘못되었기 때문입니다. y는 int 타입이 아닙니다.
*/