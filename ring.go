package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// Strukture data circular list

func main() {
	data := ring.New(100)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value" + strconv.Itoa(i)

		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
