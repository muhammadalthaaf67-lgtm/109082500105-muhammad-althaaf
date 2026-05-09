package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 0)
}

func bintang(n, i int) {
	if i != n {
		baris(i, 0)
		fmt.Println("")
		bintang(n, i+1)
	} else {
		baris(i, 0)
	}
}

func baris(i, a int) {
	if a != i {
		fmt.Print("*")
		baris(i, a+1)
	}
}
