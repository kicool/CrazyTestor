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
	"CrazyTestor/app/models"
)

type TestController struct {
	revel.Controller
}

func (tc TestController) Index() revel.Result {
	ret:=make(map[string]interface {})
	ret["success"] = true;
	return tc.RenderJson(ret)
}

func (tc TestController) Add(title,desc string) revel.Result {
	ret:=make(map[string]interface {})

	ret["success"]=true
	testBean := &models.Test{}
	testBean.Title = title
	testBean.Desc = desc
	testService.Add(testBean)
	ret["value"] = testBean
	return tc.RenderJson(ret)
}

func (tc TestController) Find() revel.Result {
	ret:=make(map[string]interface {})
	list:=testService.Find()
	ret["list"] = list
	return tc.RenderJson(ret)
}

