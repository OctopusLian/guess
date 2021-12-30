/*
 * @Description:beego框架初始化文件
 * @Author: neozhang
 * @Date: 2021-12-30 18:29:51
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 21:21:22
 */
package controllers

import (
	//"guess/models"

	"encoding/json"
	"errors"
	"guess/models"

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
	var subject models.Subject
	err := func() error {
		id, err := c.GetInt("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		subject, err = models.GetSubject(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err != nil {
		c.Ctx.WriteString("wrong params")
	}

	var option map[string]string
	if err = json.Unmarshal([]byte(subject.Option), &option); err != nil {
		c.Ctx.WriteString("wrong params, json decode")
	}
	c.Data["ID"] = subject.Id
	c.Data["Option"] = option
	c.Data["Img"] = "/static" + subject.Img
	c.TplName = "guess.tpl"
}

func (c *MainController) Post() {
	var subject models.Subject
	err := func() error {
		id, err := c.GetInt("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		subject, err = models.GetSubject(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err != nil {
		c.Ctx.WriteString("wrong params")
	}
	answer := c.GetString("key")
	right := models.Answer(subject.Id, answer)

	c.Data["Right"] = right
	c.Data["Next"] = subject.Id + 1
	c.Data["ID"] = subject.Id
	c.TplName = "guess.tpl"
}
