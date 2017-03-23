
package main

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}
func (this *MainController) Put() {
	//                  //获取上传的文件
	//path := "d:/upload/" + h.Filename
	//f.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	_, h, _ := this.GetFile("myfile")
	fmt.Println(h.Filename)
	this.SaveToFile("myfile", "d:/upload/")                    //存文件
	this.Ctx.WriteString("hello world put")
}

func (this *MainController) Get() {

	this.Ctx.WriteString("hello world")
	//fmt.Print("it is :",this.Ctx.Input.Params())
}
func (this *MainController) Post(){
	this.Ctx.WriteString("hello post")
	fmt.Println(this.Ctx.Request.PostForm)
	fmt.Println(this.Ctx.Request.Body)
	f, h, _ := this.GetFile("upload")
	if f!=nil {
		fmt.Println("it can open",h.Filename)
		f.Close()
	}else {
		fmt.Println("it can`t open")
	}
	this.SaveToFile("upload", "d:/upload/ttt.txt")
}
func main() {
	beego.Router("/*", &MainController{})
	beego.Run()
}
