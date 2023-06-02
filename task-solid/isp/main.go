package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type Faxer interface {
	Fax()
}

type MultifunctionDevice interface {
	Printer
	Scanner
	Faxer
}

type SimplePrinter struct{}

func (sp SimplePrinter) Print() {
	fmt.Println("Printing Function...")
}

type SimpleScanner struct{}

func (ss SimpleScanner) Scan() {
	fmt.Println("Scanning Function...")
}

type SimpleFaxer struct{}

func (sf SimpleFaxer) Fax() {
	fmt.Println("Faxing Function...")
}

// interface 'Printer' 'Scanner' 'Faxer' memiliki peran masing masing
// kemudian function 'MultifunctionDevice' berfungsi untuk mengkombinasikan ketiga interface

func main() {
	// ISP (Interface Segregation Principle)
	printer := SimplePrinter{}
	scanner := SimpleScanner{}
	faxer := SimpleFaxer{}
	printer.Print()
	scanner.Scan()
	faxer.Fax()
}
