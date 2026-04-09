package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	var maxHari int
	if tujuan == "domestik" {
		maxHari = 3
	} else if tujuan == "mancanegara" {
		maxHari = 8
	}

	if jumlahHari > maxHari {
		return maxHari
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	makan := 2 * 35000
	penginapan := 250000
	uangSaku := 300000

	totalPerMhs := makan + penginapan + uangSaku
	return totalPerMhs * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	biayaDomestik := float64(biayaPerHari(jumlahMhs))

	if tujuan == "domestik" {
		*totalBiaya = biayaDomestik * float64(hariDitanggung)
	} else if tujuan == "mancanegara" {
		*totalBiaya = (biayaDomestik * 1.5) * float64(hariDitanggung)
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)
	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
