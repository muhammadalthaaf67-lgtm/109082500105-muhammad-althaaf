package main

import "fmt"

func main() {
	var n, i int = 0, 1
	fmt.Scan(&n)
	turun(n, i)
	naik(n, i)
}

func turun(n, i int) {
	if n != i {
		fmt.Print(n)
		fmt.Print(" ")
		turun(n-1, i)
	}
}

func naik(n, i int) {
	if i <= n {
		fmt.Print(i)
		fmt.Print(" ")
		naik(n, i+1)
	}
}
