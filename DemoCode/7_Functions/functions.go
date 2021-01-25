package main 

func main() {
	println(swapString("World", "Hello"))
	_, d := swapString("Hi", "There")
	println(d)
}

func swapString(a string, b string) (string, string){
	return b, a
}