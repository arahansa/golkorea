/*
원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo


Slices
slice는 연속적인 자료에 대해 더 일반화하고, 강력하며, 편리한 인터페이스를 제공하기 위해 배열을 둘러쌉니다. 
변환 행렬 같은 명확한 크기를 갖는 아이템을 제외하고, Go에서 대부분의 배열 프로그램은 단순 배열 보다는 slice를 사용합니다.

slice는 참조 타입입니다. 즉, 만약 하나의 slice를 다른쪽에 지정하면 둘다 같은 배열을 참조하는 것을 의미합니다. 
예를 들어, 만약 함수가 slice 인자를 받는다면, 값이 변경되는 경우 slice의 요소들을 호출하는 쪽에서 
볼 수 있을 것입니다.(배열의 포인터를 넘겨주는 것과 유사합니다.) 

따라서, Read 함수는 포인터나 count를 인자로 받기 보다 slice를 인자로 받는 것이 좋습니다. 
slice 내부의 길이는 얼마나 많은 data를 읽을 지에 대한 제한이 있습니다. 
여기에 예로, os package에서 File 타입의 Read 메소드의 특징이 있습니다.
*/
func (file *File) Read(buf []byte) (n int, err os.Error)
/*
위의 메소드는 읽은 byte의 숫자와 error값을 반환합니다. 더 큰 buffer b로 처음 32byte를 읽기위해 buffer를 분할합니다.
*/
    n, err := f.Read(buf[0:32])
/*
이런 분할은 매우 일반적이며 효율적입니다. 사실, 잠시 효율성을 떠나서, 이 조각은 buffer의 처음 32byte를 읽을 것입니다.
*/
    var n int
    var err os.Error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        if nbytes == 0 || e != nil {
            err = e
            break
        }
        n += nbytes
    }
/*
slice의 길이는 배열의 크기 제한이 허용하는 한도 내에서 변화될 것 입니다. ; 
slice를 지정하십시오. 내장함수인 cap을 사용해 slice의 최대 용량을 알 수 있습니다. 
여기 slice에 data를 추가하는 예가 있습니다. 만약 data가 slice의 용량을 초과한다면, slice는 재할당될 것입니다. 
그 결과로 재할당된 slice가 반환됩니다. 아래 예제 함수에서는 
len()과 cap()이 nil slice에도 사용이 가능하다는 점을 사용하고 있습니다.(nil slice인 경우 0을 반환함)
*/
func Append(slice, data[]byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // 재할당
        // 공간을 더 키우기 위해 두배로 할당
        newSlice := make([]byte, (l+len(data))*2)
        // copy 함수는 미리 선언되어져있으며 어떤 slice 타입이나 동작함
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    for i, c := range data {
        slice[l+i] = c
    }
    return slice
}
/*
반드시 뒤에 slice를 리턴해야만 합니다. 왜냐하면 Append()가 slice의 요소들을 수정할 수 있다하더라도 
slice 자체는(포인터, 길이, 용량을 가지고 있는 run-time 자료구조) 값으로 넘겨져야 합니다.

slice에 data를 추가하는 아이디어는 꽤 유용하며, append 함수는 내장 함수로 포함되어 있습니다.
append 함수의 설계를 이해하기 위해서는 좀 더 많은 정보가 필요하기 때문에 추후에 다시 살펴보겠습니다.

*/