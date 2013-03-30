/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 3/30/13
 * Time: 11:41 AM
 * To change this template use File | Settings | File Templates.
 */
package models

import "CrazyTestor/app/models"

type Question struct {
	Id int
	TestId int
	Title string
	Answers []models.Answer
}

