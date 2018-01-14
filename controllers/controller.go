// Package controllers as usual should accomodate the business logic
// here only desides if xml or pdf will be returned
package controllers

import (
	"report/db"
	"github.com/gin-gonic/gin"
)

func GetReport(ctx *gin.Context, client db.ClientInterface, file_type string)  {
	Data, err := client.GetData()
	if err != nil {
		ctx.Abort()
		return
	}

	var res []byte
	if file_type == "xml" {
		res, err = Data.GenerateXML()
		if err != nil {
			ctx.Abort()
			return
		}
		// set correct headers for browsers
		ctx.Writer.Header().Set("Content-Type", "application/xml")
	}else {
		// return by default xml
		res, err = Data.GeneratePDF()
		if err != nil {
			ctx.Abort()
			return
		}
		// set correct headers for browser
		ctx.Writer.Header().Set("Content-Type", "application/pdf")
	}

	ctx.Status(200)
	ctx.Writer.Write(res)

	return
}
