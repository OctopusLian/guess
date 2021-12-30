/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-12-30 18:29:51
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 19:40:05
 */
package main

import (
	_ "guess/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:@/guess?charset=utf8")
	beego.Run()
}
