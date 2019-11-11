package main

import "github.com/astaxie/beego"

//声明控制器
type MainController struct {
	beego.Controller  //内嵌beego.Controller
}

//服务端get请求重载
func (controller *MainController) Get() { 
	name := controller.Ctx.Input.Param(":name")//路由获取数据
	controller.Ctx.WriteString("Hello, " + name)//定义网页输出
}

func main() {
	beego.Router("/mycloudgo/:name", &MainController{})//设置路由信息
	beego.Run(":8080")//默认在8080端口上运行
}
