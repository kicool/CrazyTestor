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

func (c Application) Question(testId int64) revel.Result {
	tests:=testService.Find();
	module := "question"

    if(testId ==0){
		testBean := testService.GetOne();
		testId = testBean.Id
	}


	list:=questionService.Find(testId)

	return c.Render(module,tests,testId,list)
}

func (c Application) Answer(questionId int64,testId int64) revel.Result {
	questions := questionService.Find(testId)
   	answers:=  answerService.Find(questionId)
	module := "question"
    return c.Render(questions,questionId,testId,answers,module)
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
