package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func SelectionSort(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[min] {
				min = j
			}
		}
		temp := T[i]
		T[i] = T[min]
		T[min] = temp
	}
}

func median(T arrInt, n int) float64 {
	if n%2 != 0 {
		return float64(T[n/2])
	} else {

		return float64(T[n/2-1]+T[n/2]) / 2.0
	}
}

func main() {
	var A arrInt
	var x int
	var n int = 0

	fmt.Scan(&x)

	for x != -5313541 && n < NMAX {
		if x == 0 {
			tempArr := A
			SelectionSort(&tempArr, n)
			fmt.Printf("Median : %.1f\n", median(tempArr, n))
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}
