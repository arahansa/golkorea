/*

원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo


메모리 할당으로 다시 돌아가보면, 내장함수 make(T, args)는 new(T)와는 다른 기능을 제공합니다.
make()는 slice, map, channel을 만드는 용도로만 사용됩니다.
또, *T가 아니라 T 타입의 초기 값(초기값은 0이 아님)을 반환합니다.
구분해야하는 이유는 숨겨져있고, 사용전에 반드시 초기화되어야 하는 자료 구조를 참조하는 이러한 세가지 형태가 있기 때문입니다..
slice를 예로 들어보면, slice는 데이타(배열 내부)에 대한 포인터, 길이, 용량의 세가지 아이템을 가지고 있는 descriptor이고,
 아이템들이 초기화되기 전까지 slice는 nil입니다.
make()는 slice, map, channel의 내부 데이터 구조를 초기화하고 사용할 수 있도록 해줍니다. 예를 들면,

*/

make([]int, 10, 100)
/*
   위의 예문은 100개의 int 배열을 위해 메모리 할당을 하고 배열의 처음 10개 요소를 가리키는
   길이 10, 용량 100을 가지는 slice 구조를 만듭니다.
   (slice를 만들 때, 용량은 생략될 수 있습니다.; 더 많은 정보를 원하면 slice 섹션을 보기 바랍니다.) .
   반면, {{{new(int)는 새롭게 할당된, 0값의 slice 구조에 대한 포인터를 반환합니다. (즉, nil slice에 대한 포인터를 의미함)

   다음 예제는 new()와 make()의 차이점에 대해서 설명해줍니다.
*/
var p *[]int = new([]int)      // slice구조를 할당함; *p == nil; 그다지 유용하지 않음
var v []int = make([]int, 100) // slice v는 이제 100개의 int를 가지는 새로운 배열을 참조함

// 불필요하게 복잡함:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 자연스러움:
v := make([]int, 100)
/*
      꼭 기억해야할 것은 make()는 오직 map, slice, channel을 위해서만 적용가능하다는 것이며,
   포인터를 반환하지 않습니다. 따라서, 분명한 포인터 값을 얻기 위해서는 new()를 사용해야 합니다.

*/