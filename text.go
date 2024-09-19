package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30

	fmt.Println("Siapa yang terbesar?")

	if a > b && a > c {
		fmt.Println("a yang terbesar :", a)
	} else if b > a && b > c {
		fmt.Println("b yang terbesar :", b)
	} else {
		fmt.Println("c yang terbesar :", c)
	}
}
