package main

import "fmt"

type angka int
type kata string

type buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunPenerbit angka
	jumlahHalaman angka
}

func main() {
	var biodata buku
	fmt.Println("===input biodata buku===")
	fmt.Print("masukan judul buku:")
	fmt.Scan(&biodata.judul)
	fmt.Print("masukan penulis buku:")
	fmt.Scan(&biodata.penulis)
	fmt.Print("masukan penerbit buku:")
	fmt.Scan(&biodata.penerbit)
	fmt.Print("masukan tahun penerbit buku:")
	fmt.Scan(&biodata.tahunPenerbit)
	fmt.Print("masukan jumlah halaman buku:")
	fmt.Scan(&biodata.jumlahHalaman)
	fmt.Println("===biodata buku===")
	fmt.Println("masukan penulis buku:", biodata.judul)
	fmt.Println("masukan penerbit buku:", biodata.penerbit)
	fmt.Println("masukan tahun penerbit buku:", biodata.tahunPenerbit)
	fmt.Println("masukan jumlah halaman buku:", biodata.jumlahHalaman)
}
