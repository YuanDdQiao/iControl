package main

import (
	"html/template"
	"iControl/models"
	_ "iControl/routers"
	"net/http"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
