package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 1000
	var duration3 time.Duration = time.Minute * 60
	var duration4 time.Duration = duration1 + duration3 - duration2

	fmt.Println(duration4)
	fmt.Printf("%.3f seconds\n", duration4)
}
