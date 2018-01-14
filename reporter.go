
package main

import (

	"report/controllers"
	"report/db"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	// first connect to db
	conf := db.Conf{
		Host: "candidate.suade.org",
		Database: "suade",
		User: "interview",
		Password: "LetMeIn",
		SSLMode: "disable",
	}

	dbclient, err := db.Connect(conf)
	if err != nil {
		// if no connection to db simply quit
		os.Exit(1)
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/report/xml", func(c *gin.Context){
		controllers.GetReport(c, dbclient, "xml")
	})

	r.GET("/report/pdf", func(c *gin.Context) {
		controllers.GetReport(c, dbclient, "pdf")
	})

	r.Run()
}
