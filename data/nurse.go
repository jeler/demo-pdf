package data

import (
	"github.com/brianvoe/gofakeit/v6"
)

type NurseInfo struct {
	str                  string
	ID                   string `fake:"{loremipsumsentence:10}"`
	FirstName            string `fake:"{firstName}"`
	LastName             string `faker:"lastName"`
	JobTitle             string `fake:"{jobtitle}"`
	Email                string `fake:"{email}"`
	PhoneNumber          string `fake:"{phone}"`
	SecondaryPhoneNumber string `fake:"{phone}"`
	Address              string
	City                 string
	State                string
	Zip                  string
}

func GenerateNurse() NurseInfo {
	var nurse NurseInfo
	nurse.ID = gofakeit.UUID()
	nurse.FirstName = gofakeit.FirstName()
	nurse.LastName = gofakeit.LastName()
	nurse.JobTitle = gofakeit.JobTitle()
	nurse.Email = gofakeit.Email()
	nurse.PhoneNumber = gofakeit.Phone()
	nurse.SecondaryPhoneNumber = gofakeit.Phone()
	nurse.Address = gofakeit.Address().Street
	nurse.City = gofakeit.City()
	nurse.State = gofakeit.State()
	nurse.Zip = gofakeit.Zip()
	return nurse
}