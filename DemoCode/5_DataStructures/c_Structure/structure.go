package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	b := person{name: "Bob", age: 42} //same as b:= person{"Bob", 42}

	fmt.Println(b)
	fmt.Println(b.name)

	people := []struct{
		name string
		age int
	}{
		{"Jack", 30},
		{"Jim", 31},
		{"Joe", 32},
		{"Jen", 33},
		{"Jerry",34},
	}

	fmt.Println(people)
	fmt.Println(people[3].name)
}
