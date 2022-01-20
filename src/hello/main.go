package main

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello Web 3.0")
}

func main() {
	web.Router("/", &MainController{})
	web.Run()
}
