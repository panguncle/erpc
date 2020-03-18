package main

import "fmt"

func main() {
	var i int
	i = 2
	switch i {
	default:
		fmt.Println("DEFAULT")
	case 2:
		fmt.Println(2)
	}
}
