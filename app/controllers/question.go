/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 3/30/13
 * Time: 1:05 PM
 * To change this template use File | Settings | File Templates.
 */
package controllers

import (
	"github.com/robfig/revel"
	"CrazyTestor/app/models"
)

type QuestionController struct {
	 *revel.Controller
}

func (qc QuestionController) Index () revel.Result {
	return qc.Render()
}

func (qc QuestionController) AddAnswer(questionId,nextQuestionId int64,content,result string) revel.Result{
	ret:=make(map[string]interface {})
	ret["success"] = true

	answer := &models.Answer{}
	answer.Content = content
	answer.NextQuestionId = nextQuestionId
	answer.QuestionId = questionId
	answer.Result = result
	answerService.Add(answer)

	return qc.RenderJson(ret)

}

/**

 */
func (qc QuestionController) Add(testId int64,title string) revel.Result{
	ret:=make(map[string]interface {})
	ret["success"] = true

	question:=&models.Question{}
	question.Title = title
	question.TestId = testId

	questionService.Add(question)

	return qc.RenderJson(ret)
}
