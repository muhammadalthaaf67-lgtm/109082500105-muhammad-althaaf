package main

import "fmt"

// Menetapkan batas maksimal data sesuai instruksi
const NMAX = 1000000

// Tipe data alias array integer
type arrInt [NMAX]int

// Melakukan pengurutan menggunakan algoritma selection sort
func SelectionSort(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[min] {
				min = j
			}
		}
		// Tukar posisi
		temp := T[i]
		T[i] = T[min]
		T[min] = temp
	}
}

// Menghitung median dari array yang sudah terurut
func median(T arrInt, n int) float64 {
	if n%2 != 0 {
		// Jika jumlah data ganjil, ambil nilai tengah
		return float64(T[n/2])
	} else {
		// Jika jumlah data genap, rerata dari dua nilai tengah
		return float64(T[n/2-1]+T[n/2]) / 2.0
	}
}

func main() {
	var A arrInt // array integer
	var x int    // variabel masukan
	var n int = 0

	// Membaca input awal
	fmt.Scan(&x)

	// Loop utama sampai bertemu marker -5313541
	for x != -5313541 && n < NMAX {
		if x == 0 {
			// Jika ditemukan 0, urutkan data dan tampilkan median
			// Buat salinan untuk diurutkan agar data asli tidak teracak
			tempArr := A
			SelectionSort(&tempArr, n)
			fmt.Printf("Median : %.1f\n", median(tempArr, n))
		} else {
			// Jika bukan 0, simpan ke dalam array
			A[n] = x
			n++
		}
		// Membaca input berikutnya
		fmt.Scan(&x)
	}
}
