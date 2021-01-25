package main

import "fmt"

func main() {

    nums := []int{1, 3, 5, 7, 9}

    for i, v := range nums {
        fmt.Println("index:", i, "value:", v)
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

}