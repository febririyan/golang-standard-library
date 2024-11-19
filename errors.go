package main

import (
	"errors"
	"fmt"
)

// membuat var error
var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Febry" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("Febry")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found Error")
		} else {
			fmt.Println("unknow error")
		}
	}
}
