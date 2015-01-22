package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
	//윈도우환경에서 하고 싶다면 다음으로 실행하면 된다:역주
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("c:/go"))))

}
