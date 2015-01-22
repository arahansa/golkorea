package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//res, err := http.Get("http://www.google.com/robots.txt")
	//현재 시간 계산 시작
	start := time.Now().Nanosecond()

	res, err := http.Get("http://localhost:9000") //Go로 게시판 서버 열었다.
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	//작업후 시간 계산 시작
	end := time.Now().Nanosecond()
	fmt.Println("실행시간 :", (end-start)/1000)

}
