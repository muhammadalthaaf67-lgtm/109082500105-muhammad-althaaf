package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == targetNIM {

			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, targetNIM string) int {
	max := -1
	for i := 0; i < n; i++ {
		if T[i].NIM == targetNIM {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var n int
	var searchNIM string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	if n > nMax {
		n = nMax
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("\nMasukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&searchNIM)

	pertama := cariNilaiPertama(T, n, searchNIM)
	terbesar := cariNilaiTerbesar(T, n, searchNIM)

	if pertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, pertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, terbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan.\n", searchNIM)
	}
}
