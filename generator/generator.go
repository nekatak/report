// Package generator provides the functions to convert the data
// to xml format or pdf format
package generator

import (
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"

	"github.com/signintech/gopdf"
)

// SingleInventory is responsible to unpack each obj inside inventory
type SingleInventory struct {
	Name   string	`json:name`
	Price  string	`json:price`
}

// Data is the struct to unpack the data from postgres
type Data struct {
	Organization  string      		`json:organization`
	Reported_at   string      		`json:reported_at`
	Created_at    string      		`json:created_at`
	Inventory     [] SingleInventory 	`json:inventory`
}

// GenerateXML is responsible to convert Data retrieved from postgres
// to the proper xml format
func (d Data) GenerateXML() (output []byte, err error){
	output, err = xml.MarshalIndent(d, "report", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return
}

// GeneratePDF converts Data to pdf page
// setting font with format found in pdf folder
// starting the pdf writing from top of page
// there is no check for limits of page
// given that data are 5-6 lines each time
// TODO: add check for values that could overflow the file
// TODO: return bytes (after checking the format of file)
func (d Data) GeneratePDF() (file []byte, err error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) // A4 paper dimensions
	pdf.AddPage()
	err = pdf.AddTTFFont("loma", "./Loma.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("loma", "", 10)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "THE REPORT")

	pdf.SetX(0) //move current location
	pdf.SetY(30)
	formatted_string := fmt.Sprintf("Organization: %s", d.Organization)
	pdf.Cell(nil, formatted_string) //print text

	pdf.SetX(0) //move current location
	pdf.SetY(42)
	formatted_string = fmt.Sprintf("reported at:%s", d.Reported_at)
	pdf.Cell(nil, formatted_string) //print text

	pdf.SetX(0) // move one more line down
	pdf.SetY(54)
	formatted_string = fmt.Sprintf("created at: %s", d.Created_at)
	pdf.Cell(nil, formatted_string)

	pdf.SetX(150)
	pdf.SetY(300)

	// this loop will set each of the objects in inventory in a new line in the pdf
	y := 200.0
	for _, obj := range d.Inventory {
		formattedVal := fmt.Sprintf("%s: %s", obj.Name, obj.Price)
		pdf.Cell(nil, formattedVal)  // print name: value for each in inventory
		y += 15
		pdf.SetY(y)
	}

	pdf.WritePdf("report.pdf") // return bytes?

	file, err = ioutil.ReadFile("./report.pdf") // reading again the file to convert to buffer???

	return file, err
}