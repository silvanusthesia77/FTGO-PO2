package main

import "fmt"

type Person struct {
	name  string
	age   int
	hobby string
}

func main() {
	var person = Person{name: "thoby", age: 20, hobby: "Football"}
	fmt.Println("Name :", person.name)
	fmt.Println("age :", person.age)
	fmt.Println("Hobby :", person.hobby)
}
