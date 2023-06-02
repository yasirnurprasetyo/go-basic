package main

import "fmt"

type Bentuk interface {
	Luas() float64
}

type Segitiga struct {
	Lebar   float64
	Panjang float64
}

func (r Segitiga) Luas() float64 {
	return r.Lebar * r.Panjang
}

type Lingkaran struct {
	Radius float64
}

func (c Lingkaran) Luas() float64 {
	return 3.14 * c.Radius * c.Radius
}

func TotalArea(bentuks []Bentuk) float64 {
	total := 0.0
	for _, bentuk := range bentuks {
		total += bentuk.Luas()
	}
	return total
}

// 'Bentuk' mengimplementasikan interface yang dapat diterapkan pada Segitiga atau Lingkaran
// dimana tanpa mengubah fungsi 'TotalArea' 

func main() {
	// OCP (Open-Closed Principle)
	segitiga := Segitiga{Panjang: 8, Lebar: 2}
	lingkaran := Lingkaran{Radius: 5}
	bentuk := []Bentuk{segitiga, lingkaran}
	fmt.Println("Total area : ", TotalArea(bentuk))
}
