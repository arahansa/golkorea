// 23_연습_루프와함수
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
}

func main() {
	fmt.Println(Sqrt(2))
}

/*
함수와 루프의 사용법을 익히는 간단한 연습으로, 제곱근 함수를 뉴턴의 방법(Newton's method)을 이용하여 구현합니다.
여기서 뉴턴의 방법이란 초기값 z를 선택한 후에 다음의 공식을 이용하여 반복적으로 Sqrt(x) 함수의 근사값을 찾아가는 방법을 말합니다:
z = z - (z * z - x) / (2 * z)
처음에는 계산을 10번만 반복하여 여러분이 작성한 함수가 다양한 값들 (1, 2, 3, ...)에 대하여 얼마나 정확한 값을 찾아내는지 확인합니다.
그 다음에는, 루프의 조건을 수정하여 값이 더이상 바뀌지 않을 때 (혹은 아주 작은 차이가 발생할 때) 루프를 종료하도록 합니다. 이렇게 하면 반복문의 실행 횟수가 어떻게 달라지는지 확인합니다. 결과값이 math.Sqrt 함수의 값과 얼마나 비슷한가요?
힌트: 실수(floating point)값을 선언하고 초기화 하려면, 실수값을 표현하는 문법을 사용하거나 변환 함수를 사용합니다:
z := float64(1)
z := 1.0

 Written by arahansa
 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
