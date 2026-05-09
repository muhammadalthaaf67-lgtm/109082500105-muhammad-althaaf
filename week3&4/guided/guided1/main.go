package main

import "fmt"

func volumetabung(jari_jari, tinggi int) float64 {
	var luasAalas, volume float64
	luasAalas = 3.14 * float64(jari_jari)
	volume = luasAalas * float64(tinggi)
	return volume
}

func main() {
	var r, t int
	var v1, v2 float64
	r = 5
	t = 10

	v1 = volumetabung(r, t)
	v2 = volumetabung(r, t) + volumetabung(15, t)

	fmt.Println("cara pemangilan 1")
	fmt.Println("volume tabung:", v1)

	fmt.Println("cara pemangilan 2")
	fmt.Println("volume tabung:", v2)

	fmt.Println("cara pemangilan 3")
	fmt.Println("volume tabung:", volumetabung(14, 100))
}
