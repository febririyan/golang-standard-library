package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	host := flag.String("host", "localhost", "database host")
	port := flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username: ", *username)
	fmt.Println("Password: ", *password)
	fmt.Println("Host: ", *host)
	fmt.Println("Port: ", *port)
}
