package main

import "fmt"

type Database interface {
	Save(data string)
}

type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("Saving data to MySQL:", data)
}

type PostgreSQL struct{}

func (p PostgreSQL) Save(data string) {
	fmt.Println("Saving data to PostgreSQL:", data)
}

type DataManager struct {
	Database Database
}

func (dm DataManager) SaveData(data string) {
	dm.Database.Save(data)
}

// interface 'Database' digunakan oleh 'MySQL' dan 'PostgreSQL'
// dan 'DataManager' bergantung pada abstraksi 'Database' namun bukan implementasi spesifik

func main() {
	// DIP (Dependency Inversion Principle)
	mySQL := MySQL{}
	dataManager := DataManager{Database: mySQL}
	dataManager.SaveData("Hello, World!")
}
