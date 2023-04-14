package data

import (
	"github.com/brianvoe/gofakeit/v6"
)

type Supervisor struct {
	FirstName  string `fake:"{firstname}"`
	LastName   string `fake:"{lastname}"`
	MiddleName string `fake:"{firstname}"`
	Email      string `fake:"{email}"`
}

type ExportDate struct {
	Month string `fake:"{monthstring}"`
	Day   int    `fake:"{day}"`
	Year  int    `fake:"{year}"`
	Hour  int    `fake:"{hour}"`
	Min   int    `fake:"{minute}"`
}

func GenerateSupervisor() Supervisor {
	var sup Supervisor
	gofakeit.Struct(&sup)
	return sup
}

func GenerateExportDate() ExportDate {
	var date ExportDate
	gofakeit.Struct(&date)
	return date
}
