// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/huyan/beego for the canonical source repository
// @license     http://github.com/huyan/beego/blob/master/LICENSE
// @authors     Unknwon

package controllers

import (
	"github.com/huyan/beego"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	m.Data["host"] = m.Ctx.Request.Host
	m.TplNames = "index.tpl"
}
