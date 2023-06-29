package utils

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func CalcRowHeight(m pdf.Maroto, text string, textProp props.Text, gridWidth, colWidth uint) float64 {
	pdfM := m.(*pdf.PdfMaroto)
	percent := float64(colWidth) / float64(gridWidth)
	pageWidth, _ := pdfM.Pdf.GetPageSize()
	left, _, right, _ := pdfM.Pdf.GetMargins()
	width := (pageWidth - right - left) * percent
	lines := 1.0
	if !textProp.Extrapolate {
		lines = float64(pdfM.TextHelper.GetLinesQuantity(text, textProp, width))
	}
	fontHeight := textProp.Size/pdfM.Font.GetScaleFactor() + textProp.VerticalPadding
	return lines*fontHeight + textProp.Top
}


func AreaLeft(m pdf.Maroto, area string) {
	height := m.GetCurrentOffset()
	heightStr := fmt.Sprintf("space after building %s %f", area, height)
	fmt.Println(heightStr)
}

func CalcPageSpaceRemaining(m pdf.Maroto) float64 {
	_, totalPageSize := m.GetPageSize()
	currentSpace := m.GetCurrentOffset()
	remaining := totalPageSize - currentSpace;
	fmt.Println(remaining)
	return remaining
}
