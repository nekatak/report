package generator

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {
	Convey("test for specific data you can generate successfully XML and PDF", t, func() {
		si := SingleInventory{
			Name:  "printer",
			Price: "97",
		}
		var testData = Data{
			Organization: "Fake",
			Reported_at:  "2019",
			Created_at:   "2019",
			Inventory:    []SingleInventory{si},
		}

		Convey("Check first XML....", func() {
			xmlBytes, err := testData.GenerateXML()
			s := string(xmlBytes[:])
			fmt.Println(s)

			So(err, ShouldBeNil)
			So(xmlBytes, ShouldHaveSameTypeAs, []byte{})
			check := strings.Index(s, testData.Organization)
			So(check, ShouldNotEqual, -1) // -1 is returned from strings if the string is not found
			check = strings.Index(s, testData.Inventory[0].Name)
			So(check, ShouldNotEqual, -1)
		})

		Convey("Check that PDF file is also present after calling GeneratePDF", func() {
			pdfBytes, err := testData.GeneratePDF()

			So(err, ShouldBeNil)
			So(pdfBytes, ShouldHaveSameTypeAs, []byte{})

			// check that file is found
			_, errFileNotFound := os.Stat("./report.pdf")
			So(errFileNotFound, ShouldBeNil)

			// remove report pdf before next tests
			err = os.Remove("./report.pdf")
			So(err, ShouldBeNil)
		})

	})
}
