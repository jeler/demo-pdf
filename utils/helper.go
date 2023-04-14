package utils

import (
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
