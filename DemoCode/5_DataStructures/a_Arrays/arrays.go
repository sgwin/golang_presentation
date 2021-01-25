package main

import "fmt"

func main() {
	var a [3]string
	fmt.Println(a)
	a[0] = "Hello,"
	a[1] = "Let's"
	a[2] = "Go!"

	b := [5]int{1, 2, 3, 4, 5}
	c := [2]string{"Hello", "World"}

	println(a[0])
	println(a[1])
	println(a[2])

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("size:", len(a))

}
