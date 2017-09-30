package main

import (
	"beeblog/controllers"
	"beeblog/models"
	_ "beeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	//注册数据库
	models.RegisterDB()
}

func main() {
	//开启ORM调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)

	//注册beego路由
	beego.Router("/", &controllers.HomeController{}) //首页的路由
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})       //自动路由
	beego.Router("/login", &controllers.LoginController{}) //登陆页的路由
	beego.Run()
}
