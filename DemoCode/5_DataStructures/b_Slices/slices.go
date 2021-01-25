package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7, 9} //simple array
	b := a[1:4] //take a slice of the array
	fmt.Println(a) //[1 3 5 7 9]
	fmt.Println(b) //[3 5 7]

	b[0] = 0
	fmt.Println("-------")
	fmt.Println(a) //[1 0 5 7 9]
	fmt.Println(b) //[0 5 7]

	a[1] = 3
	fmt.Println("-------")
	fmt.Println(a) //[1 3 5 7 9]
	fmt.Println(b) //[3 5 7]

	fmt.Println("-------\n-------")
	d := make([]int, len(a), len(a)+5) //Make an empty array the same size
	fmt.Println(d)
	fmt.Println("Length of d:", len(d))
	fmt.Println("Capacity of d:", cap(d))
	d = append(d, 1, 2, 3)
	e := d[4:]
	fmt.Println(d)
	fmt.Println(e)
	copy(d, a) //Copy a into d
	d[2] = 99   //Modify d
	fmt.Println("-------")
	fmt.Println(a)
	fmt.Println(d)

}
