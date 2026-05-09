package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n, " ")
	for n > 1 {
		if n%2 == 0 {
			n = genap(n)
		} else {
			n = ganjil(n)
		}
	}
}

func genap(n int) int {
	fmt.Print(n * 1 / 2)
	fmt.Print(" ")
	return n * 1 / 2
}

func ganjil(n int) int {
	fmt.Print(3*n + 1)
	fmt.Print(" ")
	return 3*n + 1
}
