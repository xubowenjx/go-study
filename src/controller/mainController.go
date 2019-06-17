package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

//开启一个服务
func main() {
	var main *MainController
	main = &MainController{}
	beego.Router("/", main)
	beego.Run()
}
