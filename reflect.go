package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name     string `required:"true" max:"10"`
	Email    string `required:"true" max:"10"`
	Password string `required:"true" max:"10"`
}

type Person struct {
	Name      string `required:"true" max:"10"`
	Pekerjaan string `required:"true" max:"10"`
	Age       int    `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, structField.Type)

		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

// Disini akan membuat validation
func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Febry", "", ""})
	readField(Person{"Yuan", "", int(90)})

	sample := Sample{
		Name:     "Berliana",
		Email:    "berliana.b@gmail.com",
		Password: "",
	}
	fmt.Println(isValid(sample))

	people := Person{
		Name:      "Berlian",
		Pekerjaan: "Go Developer",
		Age:       20,
	}
	fmt.Println(people)
}
