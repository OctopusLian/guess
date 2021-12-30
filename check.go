/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-12-30 21:24:19
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 21:24:20
 */
package main

import (
	"database/sql"

	"github.com/astaxie/beego/toolbox"
)

type DatabaseCheck struct {
}

func (*DatabaseCheck) Check() error {
	_, err := sql.Open("mysql", "root:@tcp/guess?charset=utf8")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
}
