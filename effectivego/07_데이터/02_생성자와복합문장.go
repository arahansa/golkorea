/*
때때로, zero 값으로는 충분하지 않기 때문에 여기있는 os package로 부터 파생된 예처럼 초기화 생성자가 필요할 때가 있습니다.
*/
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
/*
위 예에는 많은 정보가 들어있습니다. 우리는 매번 그 값이 계산될 때 마다 
새로운 인스턴스가 생성되는 표현인 복합문장(composite literal)을 사용해서 그것을 단순화시킬 수 있습니다.
*/
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
/*
C언어에서와는 다르게, 지역 변수의 주소를 반환하는것은 괜찮습니다. 
변수의 저장소는 함수가 반환된 뒤에도 살아있을 것입니다. 
사실, 복합문장의 주소를 가져가는 것은 그것이 계산될 때마다 새로운 인스턴스를 할당하기 때문에 
우리는 위 예에서 마지막 두 줄을 다음과 같이 섞어서 사용할 수 있습니다.
*/
    return &File{fd, name, nil, 0}
/*
복합문장의 필드는 순서대로이며 반드시 모두 있어야 합니다.
하지만, 필드:필드값 처럼 명확히 요소들을 라벨링한 경우는 초기화는 어떤 순서로든 사용될 수 있습니다. 
몇가지는 생략도 가능하며, 생략된 경우는 zero값으로 초기화됩니다. 그러므로 위의 예는 다음과 같이 사용 될 수 있습니다.
*/
    return &File{fd: fd, name: name}
/*
제한된 경우로, 만약 복합문장이 어떤 필드도 가지지 않은 경우, 그 타입에 대해 zero 값을 생성할 것입니다. 
new(File)과 &File()의 표현식은 둘다 동일합니다.

복합 문장들은 또한 array, slice, map을 위해 사용될 수 있으며, 
이 경우도 인덱싱된 필드라벨 또는 map 키를 함께 사용할 수 있습니다. 
아래 예들의 경우, Enone, Eio, Einval이 구분되어지기만 한다면 가지고 있는 값에 상관없이 정상적으로 초기화 될 것입니다.
*/
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}