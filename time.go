package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2024, time.February, 14, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-03 15:04:05"

	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	// Kita akan mengambil salah satunya
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Second())
}
