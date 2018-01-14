
package main

import (
	"github.com/astaxie/beego"

	"report/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})

	beego.Run()
}