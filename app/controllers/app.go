package controllers

import (
	"github.com/robfig/revel"
	"CrazyTestor/app/models"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}

func (c Application) Test() revel.Result {
	list:=testService.Find()
	module := "test"
	return c.Render(list,module)
}

func (c Application) Question() revel.Result {
	tests:=testService.Find();
	module := "question"

	return c.Render(module,tests)
}

func (c Application) Login() revel.Result{
	return c.Render()
}



func (c Application) Api() revel.Result {
	 testBean := &models.Test{};
	testBean.Title="123"
	testBean.Id= 1
	//c.RenderXml()
	return c.RenderXml(testBean)
}
