package main

import "fmt"

func main() {
	var satu, dua, tiga, temp string
	fmt.Print("Masukkan input : ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input : ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input : ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal : " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir : " + satu + " " + dua + " " + tiga)
}
