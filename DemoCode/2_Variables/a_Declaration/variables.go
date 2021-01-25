package main

func main() {
	var num1 int = 1 //explicity typed
	var num2 = 2     //type inferred
	num3 := 3        //declare and init
	num4 := num3
	num3 = 4

	println(num1)
	println(num2)
	println(num3)
	println(num4)
}
