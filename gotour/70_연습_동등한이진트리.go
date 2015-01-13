// 70_연습_동등한이진트리
package main

import "code.google.com/p/go-tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int)

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
}

/*
노드(leaf)들에 있는 값들의 정렬 순열는 같지만 생김새가 다른 이진트리가 있을 수 있습니다.
예를들어, 다음 그림의 두 이진 트리를 정렬 순열는 1, 1, 2, 3, 5, 8, 13 으로 같습니다.

대부분의 프로그래밍 언어에서 두 이진 트리가 같은 순열인지를 검사하는 함수의 구현은 복잡합니다.
이제 고의 동시성과 채널을 사용한 단순한 방법으로 해결해 봅시다.

이 예제는 다음의 Tree 구조체가 정의된 tree 패키지를 사용합니다.

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

1. Walk 함수를 구현하세요.

2. Walk 함수를 테스트 해 보세요.

함수 `tree.New(k)`는 k , 2k , 3k , ..., 10k 의 값을 가지는, 무작위로 구성된 이진트리를 만들어 냅니다.

채널 ch 를 만들고, 작성한 Walk 함수의 인자로 넣어 줍니다.

go Walk(tree.New(1), ch)
이제 채널에서 10개의 값을 읽어 봅니다. 읽힌 값은 1, 2, 3, ..., 10 이어야 합니다.

3. Walk 함수를 사용해 두 트리 t1 과 t2 이 값은 값들을 가지고 있는지 비교하는 Same 함수를 구현해 보세요.

4. Same 함수를 테스트 해 보세요.

`Same(tree.New(1),`tree.New(1))`의 수행결과는 true, `Same(tree.New(1),`tree.New(2))`의 수행 결과는 false 이어야 합니다.


Written by arahansa
https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
*/
