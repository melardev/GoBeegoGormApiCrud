package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SendJson(dto interface{}) {
	this.Data["json"] = dto
	this.ServeJSON()
}
