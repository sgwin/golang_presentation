package main

import (
	"time"
	"math/rand"
	"fmt"
)
func main(){
	go race("Shawn")
    race ("Chris")
}

func race(name string){
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(100)
		time.Sleep(time.Duration(x) * time.Millisecond)
		fmt.Println(name)
	}
}