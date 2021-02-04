package main

import (
	"fmt"
	"labs/lab1/myadder"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(myadder.Add(5, 6))
}
