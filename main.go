package main

import (
	"demo-pdf/data"
	"demo-pdf/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)
type RowOptions struct {
	RowHeight float64
	Field     string
	FieldData string
}

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)

	// _, heightMain := m.GetPageSize()
	// fmt.Println(heightMain)

	// total number of pages
	m.SetAliasNbPages("{nbs}")
	m.SetFirstPageNb(1)

	// readWrite()


	buildHeading(m)

	// spaceAfterHeader := m.GetCurrentOffset()
	// spaceAfterHeaderStr := fmt.Sprintf("space after building header %f", spaceAfterHeader)
	// fmt.Println(spaceAfterHeaderStr)

	buildNurseListRowsCols(m)
	// spaceAfterBuilding := m.GetCurrentOffset()
	// spaceAfterBuildingStr := fmt.Sprintf("space after building content %f", spaceAfterBuilding)
	// fmt.Println(spaceAfterBuildingStr)
	// requestedUserRow := fmt.Sprintf("%s %s", requestedUser.FirstName, requestedUser.LastName)


	buildFooter(m)

	// spaceAfterFooter := m.GetCurrentOffset()
	// spaceAfterFooterStr := fmt.Sprintf("space after footer %f", spaceAfterFooter)
	// fmt.Println(spaceAfterFooterStr)

	// buildNurseListRowsCols(m)



	err := m.OutputFileAndClose("pdfs/moroto-demo.pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}
	// fmt.Println("PDF saved successfully")
}

// fn for trying to get an image from url
// func readWrite() {
// 	r, err := http.Get(logoURL)
// 	if err != nil {
// 		log.Println("Cannot get from URL", err)
// 	}
// 	defer r.Body.Close()

// 	data, _ := ioutil.ReadAll(r.Body)
// 	ioutil.WriteFile("rw.data", data, 0755)
// }

func buildHeading(m pdf.Maroto) {

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(6, func() {
				err := m.FileImage("images/go-logo.jpg", props.Rect{
					Top:     0,
					Left:    3,
					Percent: 50,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}

			})
			m.Col(6, func() {
				m.Text("Clinician Profile", props.Text{
					Style: consts.Bold,
					Align: consts.Right,
					Color: color.Color{},
				})
			})
		})
	})
}

func buildContentContainer(m pdf.Maroto) {
	buildNurseListRowsCols(m)

}


func buildNurseListRowsCols(m pdf.Maroto) {
	requestedUser := data.GenerateNurse()
	requestedUserRow := fmt.Sprintf("%s %s", requestedUser.FirstName, requestedUser.LastName)
	m.Row(20, func() {
		m.Col(12, func() {
			m.Text("Clinician Profile Report", props.Text{
				Size: 20,
			})
		})
	})

	m.Row(20, func() {
		m.Col(12, func() {
			m.Text(requestedUserRow, props.Text{
				Align: consts.Left,
				Size:  16,
			})
		})
	})

	// Table of user data
	m.SetBorder(true)

	// Id Row
	IdRow := &RowOptions{10, "ID", requestedUser.ID}

	generateRow(m, IdRow)

	// utils.AreaLeft(m, "after id row")

	roleRow := &RowOptions{10, "Role", requestedUser.JobTitle}

	generateRow(m, roleRow)

	// utils.AreaLeft(m, "after role row")

	// Email
	emailRow := &RowOptions{10, "Email", requestedUser.Email}

	generateRow(m, emailRow)

	// utils.AreaLeft(m, "after email row")
	utils.CalcPageSpaceRemaining(m)


	// var formattedAddress = fmt.Sprintf("%s %s, %s %s", requestedUser.Address, requestedUser.City, requestedUser.State, requestedUser.Zip)


	// addressTextProps := props.Text{Top: 3, Left: 3}
	// addressRowHeight := utils.CalcRowHeight(m, formattedAddress, addressTextProps, 6, 6)
	// fmt.Println(addressRowHeight)
	m.Line(2, props.Line{
		Width: 0,
	})

	createFakeAddressRow(m, requestedUser)
	phoneRow := &RowOptions{10, "Phone Number", requestedUser.PhoneNumber}
	generateRow(m, phoneRow)
	secondaryPhoneRow := &RowOptions{10, "Secondary Phone Number", requestedUser.SecondaryPhoneNumber}
	generateRow(m, secondaryPhoneRow)

}

// will need to add 3 vertical div lines if using this
func createFakeAddressRow(m pdf.Maroto, requestedUser data.NurseInfo) {
	formattedSecondLine := fmt.Sprintf("%s, %s %s", requestedUser.City, requestedUser.State, requestedUser.Zip)
	m.SetBorder(false)

	m.Row(5, func() {
		m.Col(6, func() {
			m.Text("Address", props.Text{
				Style: consts.Bold,
				Left:  3,
				Top:   3,
			})
		})
		m.Col(6, func() {
			m.Text(requestedUser.Address, props.Text{
				Left: 3,
			})
			// attempting to use the underlying package to create multicell
			// gofpdf.Pdf.MultiCell(0, 5, requestedUser.Address, "", "", false)

		})
	})
	m.Row(5, func() {
		m.Col(6, func() {

		})
		m.Col(6, func() {
			m.Text(formattedSecondLine, props.Text{
				Align: consts.Left,
				Left:  3,
			})
		})
	})
	m.SetBorder(true)
}

func generateRow(m pdf.Maroto, r *RowOptions) {
	m.Row(r.RowHeight, func() {
		m.Col(6, func() {
			m.Text(r.Field, props.Text{
				Style: consts.Bold,
				Left:  3,
				Top:   3,
			})
		})
		m.Col(6, func() {
			m.Text(r.FieldData, props.Text{
				Left: 3,
				Top:  3,
			})
		})
	})
}

// will automatically overflow
//func buildNurseList(m pdf.Maroto) {
//	tableHeadings := []string{"ID", "Job Title", "Email", "Phone Number", "Secondary Phone Number"}
//	contents := data.GenerateNurse()
//	m.Row(10, func() {
//		m.Col(12, func() {
//			m.Text("Nurse Info", props.Text{
//				Top:    2,
//				Size:   13,
//				Color:  color.NewWhite(),
//				Family: consts.Courier,
//				Style:  consts.Bold,
//				Align:  consts.Center,
//			})
//		})
//	})
//
//	m.SetBackgroundColor(color.NewWhite())
//
//	m.TableList(tableHeadings, contents, props.TableList{
//		HeaderProp: props.TableListContent{
//			Size:      9,
//			GridSizes: []uint{3, 1, 3, 3, 2},
//		},
//		ContentProp: props.TableListContent{
//			Size:      8,
//			GridSizes: []uint{3, 1, 3, 3, 2},
//		},
//		Align:              consts.Top,
//		HeaderContentSpace: 1,
//		Line:               false,
//	})
//
//}

func buildFooter(m pdf.Maroto) {
	var supervisor = data.GenerateSupervisor()
	var exportDate = data.GenerateExportDate()
	placeholderText := fmt.Sprintf("Exported on %s %d, %d at %d:%d AM by: %s %s %s %s", exportDate.Month, exportDate.Day, exportDate.Year, exportDate.Hour, exportDate.Min, supervisor.FirstName, supervisor.MiddleName, supervisor.LastName, supervisor.Email)
	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text(placeholderText, props.Text{
					Align: consts.Left,
					Size:  8,
				})
				m.Text(strconv.Itoa(m.GetCurrentPage())+"/{nbs}", props.Text{
					Align: consts.Right,
					Size:  8,
				})
			})
		})
	})
}