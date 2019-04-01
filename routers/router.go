package routers

import (
	"github.com/astaxie/beego"
	"mite_beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
