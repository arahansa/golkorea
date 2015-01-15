package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines, err := readLines("sample.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
}

// readLines 은 파일을 메모리로 읽어들여,
// 파일들 라인의 슬라이스를 리턴한다. 참고 http://bit.ly/golangreadwrite
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() //defer 는 다른 언어문법에서 finally 같은 것을 의미함.

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
