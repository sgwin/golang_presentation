package main

func main() {
	num1 := 123
	modNoPointer(num1)
	println(num1, "<-- After modNoPointer")
	modWithPointer(&num1)
	println(num1, "<-- After modWithPointer")
}

func modNoPointer(x int) {
	x += 50
	println(x, "<-- Inside modNoPointer")
}

func modWithPointer(xPtr *int) {
	*xPtr += 77
	println(*xPtr, "<-- Inside modWithPointer")
}
