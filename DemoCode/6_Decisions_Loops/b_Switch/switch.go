package main

func main() {
	x := "go"

	switch x {
	case "not":
		println("Word = 'not'")
	case "blue":
		println("Word = 'blue'")
	case "go":
		println("Word = 'go'")
	default:
		println("No match")
	}
}
