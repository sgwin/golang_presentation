package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[6])
	fmt.Println("Hello " + "World")
	fmt.Println("Hello\n\tWorld")
	fmt.Println(`-----
Hello
	World
Line3`)
}
