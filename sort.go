package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// Membuat type Deklarasion
type UserSlice []User

// Membuat contract untuk User nya
func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"Febry", 25},
		{"Berliana", 24},
		{"Selicia", 20},
		{"Umma", 18},
	}

	sort.Sort(UserSlice(users)) // Harus memasukkan interface dari userslice nya
	fmt.Println(users)
}
