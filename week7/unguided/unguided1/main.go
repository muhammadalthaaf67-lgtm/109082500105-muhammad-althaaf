package main

import "fmt"

type suhu float64

func main() {
	var n suhu
	fmt.Print("masukan suhu :")
	fmt.Scan(&n)
	fmt.Println("reamur: ", reamur(n))
	fmt.Println("fahrenheit: ", fahrenheit(n))
	fmt.Println("kelvin: ", kelvin(n))
}

func reamur(c suhu) suhu {
	return 0.8 * c
}

func fahrenheit(c suhu) suhu {
	return (1.8 * c) + 32
}

func kelvin(c suhu) suhu {
	return c + 273
}
