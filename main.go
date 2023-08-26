package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {

	items := []SendItem{
		{PickupCode: "ABC123"},
		{PickupCode: "DEF456"},
		{PickupCode: "GHI781"},
		{PickupCode: "GHI782"},
		{PickupCode: "GHI783"},
		{PickupCode: "GHI784"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
		{PickupCode: "GHI785"},
	}

	GeneratePDF(items)
	lain(items)

}

type SendItem struct {
	PickupCode string
}

func lain(items []SendItem) {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)

	rowsPerPage := 7
	colsPerRow := 4

	for i := 0; i < len(items); i += rowsPerPage * colsPerRow {
		m.Row(50*float64(rowsPerPage), func() {
			for r := 0; r < rowsPerPage; r++ {
				m.Row(30, func() {
					for c := 0; c < colsPerRow; c++ {
						index := i + r*colsPerRow + c
						if index < len(items) {
							m.Col(12/uint(colsPerRow), func() {
								m.QrCode(items[index].PickupCode, props.Rect{
									Percent: 70,
									Center:  true,
								})
								m.Text(items[index].PickupCode, props.Text{
									Size:  10,
									Align: consts.Bottom,
								})
							})
						}
					}
				})
			}
		})
	}

	m.SetBorder(false)

	naming := time.Now().Format("2006-01-02-15-04-05")
	pdfPath := fmt.Sprintf("qrcodes-%s.pdf", naming)
	err := m.OutputFileAndClose(pdfPath)
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))

}

// func lain(items []SendItem) {
// 	begin := time.Now()
// 	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
// 	m.SetBorder(true)

// 	rowsPerPage := 7
// 	colsPerRow := 4

// 	for i := 0; i < len(items); i += rowsPerPage * colsPerRow {
// 		numRows := min(rowsPerPage, (len(items)-i+colsPerRow-1)/colsPerRow)

// 		if i > 0 { // Tambahkan halaman baru sebelum mengisi data QR code pada halaman selanjutnya
// 			m.AddPage()
// 		}

// 		m.Row(50*float64(numRows), func() {
// 			for r := 0; r < numRows; r++ {
// 				m.Row(50, func() {
// 					for c := 0; c < colsPerRow; c++ {
// 						index := i + r*colsPerRow + c
// 						if index < len(items) {
// 							m.Col(12/uint(colsPerRow), func() {
// 								qrSize := 33.33 // Persentase untuk QR code 4x4 cm
// 								m.QrCode(items[index].PickupCode, props.Rect{
// 									Percent: qrSize,
// 									Center:  true,
// 								})
// 								m.Text(items[index].PickupCode, props.Text{
// 									Size:  10,
// 									Align: consts.Bottom,
// 								})
// 							})
// 						}
// 					}
// 				})
// 			}
// 		})
// 	}

// 	m.SetBorder(false)

// 	naming := time.Now().Format("2006-01-02-15-04-05")
// 	pdfPath := fmt.Sprintf("qrcodes-%s.pdf", naming)
// 	err := m.OutputFileAndClose(pdfPath)
// 	if err != nil {
// 		fmt.Println("Could not save PDF:", err)
// 		os.Exit(1)
// 	}

// 	end := time.Now()
// 	fmt.Println(end.Sub(begin))
// }

func GeneratePDF(items []SendItem) {

	fmt.Println(items)

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetBorder(true)

	rowsPerPage := 7
	colsPerRow := 4

	for i := 0; i < len(items); i += rowsPerPage * colsPerRow {
		m.Row(50*float64(rowsPerPage), func() {
			for r := 0; r < rowsPerPage; r++ {
				m.Row(30, func() {
					for c := 0; c < colsPerRow; c++ {
						index := i + r*colsPerRow + c
						if index < len(items) {
							m.Col(12/uint(colsPerRow), func() {
								m.QrCode(items[index].PickupCode, props.Rect{
									Percent: 70,
									Center:  true,
								})
								m.Text(items[index].PickupCode, props.Text{
									Size:  10,
									Align: consts.Bottom,
								})
							})
						}
					}
				})
			}
		})
	}

	m.SetBorder(false)

	pdf, err := m.Output()
	fmt.Println(pdf)
	if err != nil {
		fmt.Println("Could not save PDF:", err)
	}

	err = m.OutputFileAndClose("qrgrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
	}

}
