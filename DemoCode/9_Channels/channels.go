package main

func addSeries(n int, c chan int){
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	c <- sum
}

func main(){
	ch := make(chan int)
	go addSeries(100, ch)
	go addSeries(200, ch)
	a, b := <-ch, <-ch
	println(a)
	println(b)
}