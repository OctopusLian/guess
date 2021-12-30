/*
 * @Description:路由
 * @Author: neozhang
 * @Date: 2021-12-30 18:29:51
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 19:40:58
 */
package routers

import (
	"guess/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.MainController{})
}
