package main

import "fmt"

type Kendaraan interface {
	BahanBakar() string
}

type Pesawat struct{}

func (p Pesawat) BahanBakar() string {
	return "Avtur"
}

type Mobil struct{}

func (m Mobil) BahanBakar() string {
	return "Bensin"
}

// 'Kendaraan' adalah interface yang digunakan oleh 'Pesawat' dan 'Mobil'
// dan keduanya dapat secara bergantian memanggil fungsi 'BahanBakar()'

func main() {
	// LSP (Liskov Substitution Principle)
	kendaraans := []Kendaraan{Pesawat{}, Mobil{}}
	for _, kendaraan := range kendaraans {
		fmt.Println(kendaraan.BahanBakar())
	}
}
