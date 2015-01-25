package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//out, err := exec.Command("cmd", "/C", "dir").Output()
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
