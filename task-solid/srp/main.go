package main

import "fmt"

type Angka struct{}

func (a Angka) Penjumlahan(x, y int64) int64 {
	return x + y
}

func (a Angka) Pengurangan(x, y int64) int64 {
	return x - y
}

// pada penerapan SRP Angka memiliki tanggung jawab 
// untuk melakukan operasi matematika penjumlahan dan pengurangan

func main() {
	//SRP (Single Responsibility Principle)
	angka := Angka{}
	fmt.Println("Penjumlahan: ", angka.Penjumlahan(8, 5))
	fmt.Println("Pengurangan: ", angka.Pengurangan(8, 5))
}