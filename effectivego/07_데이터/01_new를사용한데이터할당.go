/*
원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo

new()를 사용한 메모리 할당(Allocation with new())
Go언어에서는 new()와 make(), 두가지 메모리 할당 원형이 있습니다. 
new()와 make()는 서로 다르게 동작하며 다른 형태로 적용되어 지기 때문에 혼란스러울 수 있습니다. 
하지만 규칙은 단순합니다. 먼저 new()에 대해서 얘기하자면, new()는 다른 언어들에서와 마찬가지로 
기본적으로 탑재되어 있는 함수 입니다.: 
new(T)는 새로운 아이템 T를 위해 크기 0의 저장공간을 할 당할 것이며, 
*T 타입의 값을 가지고 있는 할당된 공간의 주소를 리턴할 것입니다. 

Go 용어로 말하자면, new(T)는 새로 할당된 0 값의 타입 T에 대한 포인터를 리턴할 것입니다.

new()에 의해 반환된 메모리는 0 이기 때문에 객체는 추가적인 초기화없이 사용될 수 있습니다. 
이것은 자료구조의 사용자가 new()로 하나를 새롭게 만들 수 있고 바로 작업할 수 있다는 것을 의미합니다. 
예를 들어, bytes.Buffer문을 위해 주석을 단다면 "0값의 Buffer는 사용준비가 되어진 비어있는 buffer 입니다"라고 할 수 있습니다. 
비슷하게, sync.Mutex는 분명한 생성자나 초기화 메소드를 가지고 있지 않습니다. 
대신 sync.Mutex를 위한 0값은 mutex를 해제하기 위해 정의되어져 있습니다.

zero 값이 유용하다는 속성은 타동적으로 동작합니다. 이런 형태의 정의를 고려해보십시오.
*/
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
/*
SyncedBuffer? 타입의 값들은 할당 또는 정의된 즉시 사용할 준비가 되어져 있습니다. 
다음에서 p와 v는 추가적인 처리 없이 정확히 동작할 것입니다.
*/
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer