/*
 * @Description:路由
 * @Author: neozhang
 * @Date: 2021-12-30 18:29:51
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 21:21:58
 */
package routers

import (
	"guess/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Include(&controllers.UserController{})
	//beego.Router("/", &controllers.UserController{})
	beego.Router("/", &controllers.MainController{})
}
