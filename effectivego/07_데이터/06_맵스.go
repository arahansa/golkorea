/*Maps
원글 링크 : https://code.google.com/p/golang-korea/wiki/EffectiveGo


Map은 서로 다른 타입의 값들에 접근하기 위한 편리하고 강력한 자료구조 입니다. 
Key는 동등연산자(==)이 가능한 integer, float, 복소수, string, 포인터, 
그리고 인터페이스(dynamic 타입이 동등연산을 지원한다면)와 같이 어떤 타입이든 될 수 있습니다. 

반면, struct와 배열 그리고 slice는 동등연산을 할 수 없기 때문에 map key로는 사용할 수 없습니다. 
slice처럼 map은 참조타입 입니다. 함수에 map을 값으로 넘기고 함수 안에서 변경하게 되면, 
map 값의 변경사항을 호출자가 확인할 수 있을 것 입니다.

map은 일반적으로 콜론으로 구분되어지는 key값들을 사용해 생성하게 됩니다. 
즉, map을 초기화하여 만드는 것은 쉽습니다.
*/

var timeZone = map[string] int {
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
//map의 값을 지정하고 불러오는 것은 인덱스에 정수를 사용하지 않는 것 말고는 배열과 동일한 것 처럼 보입니다.

offset := timeZone["EST"]
/*map에서 map에 존재하지 않는 key를 사용해 값을 불러오려고 시도하면 0 값을 반환할 것 입니다. 
예를 들면, map이 integer값을 포함하고 있다면, 존재하지 않는 key 값으로 검색을 하면 0을 반환하게 됩니다. 
집합(set)은 bool 타입의 map으로 구현될 수 있습니다. 
집합에 포함시키기 위해 map 항목의 값을 true로 정하면 됩니다. 그리고 인덱스로 테스트 해보십시오.*/

attended := map[string] bool {
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
/*
때로는 zero 값을 사용해서 빠진 항목을 구분해야할 필요가 있습니다. 
(위 예제에서) UTC라는 항목이 있는지 또는 그것이 map 안에 전혀 없기 때문에 값이 zero 값인지 
다중 할당(multiple assignment)을 사용해서 구분할 수 있습니다.
*/
var seconds int
var ok bool
seconds, ok = timeZone[tz]
/*
이 것은 "comma ok"라고 불려집니다. 예를 들어, 만약 tz가 존재한다면, seconds는 적절한 값이 
들어갈 것이고 ok는 true가 될 것 입니다. 만약 그렇지 않다면, seconds는 zero값이 될 것이고 
ok는 false가 될 것 입니다. 여기 그 두가지를 함께 사용한 error report 함수 예제가 있습니다.
*/

func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
/*
실제 값에 대해 map 안에 존재하는지 여부를 시험하기 위해서, blank 식별자나 underscore(_)를 사용할 수도 있습니다. 
blank 식별자는 지정되거나(assign) 또는 어떤 타입의 값으로든 선언될 수 있습니다. 
map에 존재하는지 테스트하기 위해서 일반적인 변수를 사용하는 곳에 blank 식별자를 사용해 보십시오.
*/

_, present := timeZone[tz]
/*
map의 항목을 삭제하기 위해서는, 다중 할당(multiple assignment)에서 위치를 서로 바꾸고, 
boolean 값을 오른쪽에 위치시키면 됩니다.(아래 예와 같이) 만약 boolean 값이 false라면 
그 항목은 삭제될 것입니다. 심지어, map에서 이미 그 key가 삭제되었다고 하더라도 이렇게 하는 것은 안전합니다.
*/

timeZone["PDT"] = 0, false  // Now on Standard Time