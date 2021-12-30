/*
 * @Description:beego框架初始化文件
 * @Author: neozhang
 * @Date: 2021-12-30 18:29:51
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 19:21:16
 */
package controllers

import (
	//"guess/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }

func (c *MainController) Get() {
	// var subject models.Subject
	// err := func() {
	// 	id, err := c.GetInt("id")
	// 	beego.Info(id)
	// 	if err != nil {
	// 		id = 1
	// 	}
	// 	subject, err = models.Ge
	// }
}
