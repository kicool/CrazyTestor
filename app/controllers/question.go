/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 3/30/13
 * Time: 1:05 PM
 * To change this template use File | Settings | File Templates.
 */
package controllers

import "github.com/robfig/revel"

type QuestionController struct {
	 *revel.Controller
}

func (qc QuestionController) Index () revel.Result {
	return qc.Render()
}
/**

 */
func (qc QuestionController) Add(testId int64,title string) revel.Result{
	ret:=make(map[string]interface {})
	return qc.RenderJson(ret)
}

func (qc QuestionController) AddAnswer(questionId,nextId int64,content,result string)revel.Result{
	ret:=make(map[string]interface {})
	return qc.RenderJson(ret)
}
