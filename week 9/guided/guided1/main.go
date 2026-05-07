package main

import "fmt"

type arrInt [2023]int

func terkecil_1(tabInt arrInt, n int) int {
	var min int = tabInt[0]
	var j int = 1

	for j < n {
		if min > tabInt[j] {
			min = tabInt[j]
		}
		j = j + 1
	}
	return min
}

func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0
	var j int = 1

	for j < n {
		if tabInt[idx] > tabInt[j] { // [9, 8, 5, 6]
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var n int
	var arr arrInt

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan elemen array :")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	min := terkecil_1(arr, n)
	idx := terkecil_2(arr, n)

	fmt.Println("Nilai terkecil: ", min)
	fmt.Println("Indeks nilai terkecil: ", idx)
}
     


soal 2




package main

import "fmt"

// Mendefinisikan konstanta nMax
const nMax = 51

// Mendefinisikan struct mahasiswa
type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

// Mendefinisikan tipe arrayMahasiswa
type arrayMahasiswa [nMax]mahasiswa

// Fungsi untuk mencari nilai pertama seorang mahasiswa dengan NIM tertentu
func cariNilaiPertama(T arrayMahasiswa, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == targetNIM {
			// Mengembalikan nilai pertama yang ditemukan (indeks terkecil)
			return T[i].nilai
		}
	}
	return -1 // Penanda jika NIM tidak ditemukan
}

// Fungsi untuk mencari nilai terbesar seorang mahasiswa dengan NIM tertentu
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

	// Bagian a: Menerima masukan jumlah data N
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	// Validasi agar n tidak melebihi kapasitas array nMax
	if n > nMax {
		n = nMax
	}

	// Bagian a: Menyimpan N data mahasiswa ke dalam array
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("\nMasukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&searchNIM)

	// Bagian b & c: Pemanggilan fungsi pencarian
	pertama := cariNilaiPertama(T, n, searchNIM)
	terbesar := cariNilaiTerbesar(T, n, searchNIM)

	// Bagian d: Menampilkan hasil pencarian
	if pertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, pertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, terbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan.\n", searchNIM)
	}
}