package main

import "fmt"

func main() {
	var kendaraan string
	var total int
	i := 1
	for kendaraan != "--" {
		fmt.Printf("kendaraan %v\n", i)
		fmt.Print("kendaraan jenis apa? (motor/mobil/-- untuk selesai): ")
		fmt.Scan(&kendaraan)
		switch kendaraan {
		case "motor":
			total += motor(i)
		case "mobil":
			total += mobil(i)
		}
		i++
	}
	fmt.Printf("***total pendapatan parkir %v***", total)
}

func motor(i int) int {
	var masuk, keluar, hasil, sisa int
	fmt.Println("format jam 12-24")
	fmt.Print("jam masuk: ")
	fmt.Scan(&masuk)
	fmt.Print("jam keluar: ")
	fmt.Scan(&keluar)
	sisa = keluar - masuk
	if keluar > 17 {
		if masuk == 12 {
			masuk = 0
		}
		hasil = (17 - masuk) * 4000
		hasil += (keluar - 17) * 5000
	} else {
		hasil = sisa * 4000
	}
	fmt.Printf("biaya parkir kendaraan %v :%v\n", i, hasil)
	fmt.Println("=========================================================")
	return hasil
}

func mobil(i int) int {
	var masuk, keluar, hasil, sisa int
	fmt.Println("format jam 12-24")
	fmt.Print("jam masuk: ")
	fmt.Scan(&masuk)
	fmt.Print("jam keluar: ")
	fmt.Scan(&keluar)
	sisa = keluar - masuk
	if keluar > 17 {
		if masuk == 12 {
			masuk = 0
		}
		hasil = (17 - masuk) * 6000
		hasil += (keluar - 17) * 7000
	} else {
		hasil = sisa * 6000
	}
	fmt.Printf("biaya parkir kendaraan %v :%v\n", i, hasil)
	fmt.Println("=========================================================")
	return hasil
}
