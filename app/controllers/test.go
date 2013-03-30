/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 3/30/13
 * Time: 12:40 PM
 * To change this template use File | Settings | File Templates.
 */
package controllers

import (
	"github.com/robfig/revel"
)

type TestController struct {
	revel.Controller
}

func (tc TestController) Index() revel.Result {
	ret:=make(map[string]interface {})
	ret["success"] = true;
	return tc.RenderJson(ret)
}

func (tc TestController) Add(title string) revel.Result {
	ret:=make(map[string]interface {})

	ret["title"]=title
	return tc.RenderJson(ret)
}

