package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Febri Riyanto", "Riyan"))
	fmt.Println(strings.Split("Febri Riyanto", " "))
	fmt.Println(strings.ToLower("Febri Riyanto"))
	fmt.Println(strings.ToUpper("Febri Riyanto"))
	fmt.Println(strings.Trim("       Febri Riyanto       ", " "))
	fmt.Println(strings.ReplaceAll("Febri Riyan Febri", "Febri", "Hero"))
}
