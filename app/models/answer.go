/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 3/30/13
 * Time: 11:43 AM
 * To change this template use File | Settings | File Templates.
 */
package models

type Answer struct {
	Id int
	Content string
	Result string
	NextQuestionId int
}

