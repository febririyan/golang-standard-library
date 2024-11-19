package main

import (
	"container/list"
	"fmt"
)

// Strukture data double linked

func main() {
	data := list.New()

	data.PushBack("Febry")
	data.PushBack("Riyanto")
	data.PushBack("Kaya")

	head := data.Front() // ambil data di depannya
	fmt.Println(head.Value)

	next := head.Next() // ambil data selanjutnya
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for f := data.Front(); f != nil; f = f.Next() {
		fmt.Println(f.Value)
	}
}
