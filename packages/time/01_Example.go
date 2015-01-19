// 01_Example.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1초마다 계속 울림 ")
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	t2 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t2.Local())

	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}

	for {
		select {
		case <-time.After(1000 * time.Millisecond):
			fmt.Println("timed out : 약 1 초 지남. 지금시간은 :", time.Now())

		}
	}

}
