/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-12-30 18:29:51
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 19:14:52
 */
package main

import (
	_ "guess/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
