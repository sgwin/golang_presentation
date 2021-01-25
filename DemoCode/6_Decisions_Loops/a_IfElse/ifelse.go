package main 

import (
	"math/rand"
	"fmt"
)

func main() {
	if x:=rand.Intn(100); x%2 == 0 {
		fmt.Println(x, "is Even")
	}else {
		fmt.Println(x, "is Odd")
	}
}