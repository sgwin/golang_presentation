package main

import f "fmt" //Notice the alias for the imported library

func main(){
	type Coord struct {
		Lat, Long float64
	}
	var m map[string]Coord

	m = make(map[string]Coord)

	m["Albany_NY"]=Coord{42.6526, 73.7562}
	f.Println(m["Albany_NY"])
	f.Println(m["Albany_NY"].Long)
}