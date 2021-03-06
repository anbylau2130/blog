package routers

import (
	"blog/controllers"
	"blog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "*:Home")
	beego.Router("/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/Logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/main", &controllers.MainController{}, "*:Home")
	beego.Router("/main/Home", &controllers.MainController{}, "*:Home")
	beego.Router("/main/CheckSSO", &controllers.MainController{}, "*:CheckSSO")
	beego.Router("/main/GetPWD", &controllers.MainController{}, "*:GetPWD")
	beego.Router("/main/Login", &controllers.MainController{}, "*:Login")
	beego.Router("/main/GetMenuHorizontal", &controllers.MainController{}, "*:GetMenuHorizontal")
	beego.Router("/main/GetMenusVertical", &controllers.MainController{}, "*:GetMenusVertical")
	beego.AutoPrefix("/", &controllers.ServiceController{})
	beego.AutoPrefix("/admin", &admin.MenusController{})
	beego.AutoPrefix("/admin", &admin.CorpController{})
	beego.AutoPrefix("/admin", &admin.UserController{})
	beego.AutoPrefix("/admin", &admin.ChatRoomController{})
	beego.AutoPrefix("/admin", &admin.RoleController{})
	beego.AutoPrefix("/admin", &admin.PrivilegeController{})
	//beego.Include(&controllers.MainController{})
	//beego.AutoPrefix("/Admin",&controllers.AdminController{})

}
