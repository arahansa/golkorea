/*
    author twitter : @mmcgrana
    https://gobyexample.com/json
	Go by Example: json

	Go는 내장된 내장된 커스텀 데이터타입으로의 JSON 인코딩과 디코딩을 내장형태로 지원한다.
*/

package main

import "encoding/json"
import "fmt"
import "os"

// 우리는 이 두가지 구조체를 사용하여 커스텀타입의 인코딩과 디코딩을 해볼 것이다.
type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	//먼저 우리는 기본데이터타입을 JSON 문자열로 인코딩하는 것을 살펴볼 것이다.
	// 여기 원자값(atomic)을 위한 몇가지 예제가 있다.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 그리고 여기에 슬라이스와 맵을 위한 몇가지 것들이 있는데,
	// 당신이 기대한 JSON 배열과 오브젝트들로 인코딩해준다.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 패키지들은 자동으로 당신의 커스텀 데이터타입을 인코딩할 수있다.
	// 그것은 인코딩결과에서 오직 exported된 필드만 포함하고, 기본적으로 그것들의 이름을 JSON 키로 이용한다.
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 당신은 구조체필드선언에서 태그를 사용하여서 인코딩된 JSON 키이름을 커스터마이징할수 있다.
	// 이러한 태그의 예를 보기 위해 Response2의 정의를 체크해봐라.
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//이제 JSON데이터가 Go 값으로 디코딩되는 것을 살펴보자.
	//여기 일반적인 데이터구조를 위한 예제가 있다.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 우리는  JSON 패키지가 디코딩된 데이터를 놓는 곳에 변수를 제공할 필요가 있다
	// 이 map[string]interface{}는 무작위 데이터에 대한 문자열맵을 가지고 있다.

	var dat map[string]interface{}

	//여기 실제 디코딩이 있고 연관된 에러를 체크한다.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 디코딩된 맵에서 값을 사용하기 위해 우리는 적절한 타입으로 형변환을 해줘야 한다.
	// 예를 들자면 여기서는 우리는 `num`값을 우리가 예상한 `float64`타입으로 형변환해준다.
	num := dat["num"].(float64)
	fmt.Println(num)

	// 중첨된 데이터에 접근하는 것은 형변환의 연속을 필요로 한다.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//우리는 또한 JSON 을 커스텀 데이터 타입에 디코딩할 수 있다.
	// 이것은 추가적인 타입안전(type-safety)을 우리 프로그램에 더하는 것의 장점을 가지고,
	// 디코딩된 데이터에 접근할 떄, type assertions들에 대한 필요성을 제거해준다.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 예제에서 우리는 언제나 bytes와 string을 data와 JSON표현을 중재하는 데 썼지만,
	// 우리는 또한 JSON 인코딩을 직접적으로 `os.Stdout`나 심지어 HTTP 응답몸체같은 `os.Writer`에 데이터전송을 할 수 있다.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
