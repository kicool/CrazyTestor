/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 3/30/13
 * Time: 12:34 PM
 * To change this template use File | Settings | File Templates.
 */
package services

import (
	"labix.org/v2/mgo"
	"CrazyTestor/app/models"
	"labix.org/v2/mgo/bson"
)

var (
	_questionServiceInstance *QuestionService
)

const (
	QUESTION_COLLECTION = "questions"
)

type QuestionService struct {
	MongodbService
}

func (us *QuestionService) Get(id int64) ( *models.Question,[]models.Answer) {
	question:=&models.Question{}
	answers:=[]models.Answer{}
	if(id == 0){
	   question = us.GetOne()
	} else {
	   us.c.Find(bson.M{"id":id}).One(question)
	}
	answers = GetAnswerService().Find(question.Id)
	return question,answers
}

func (us *QuestionService) GetAnswer(questionId int64, content string) *models.Answer {
	question := &models.Question{}
	answers := []models.Answer{}
	us.c.Find(bson.M{"id": questionId}).One(question)

	answers = GetAnswerService().Find(question.Id)
	for _, answer:= range answers {
		 if(answer.Content == content) {
			 return &answer
		 }
	}

	return nil


}


func (us *QuestionService) Count() int {
	ret := 0
	ret, _ = us.c.Count()
	return ret
}

func (us *QuestionService) GetOne() *models.Question {
	result := &models.Question{}
	us.c.Find(nil).One(result)
	return result
}
func (us *QuestionService) Find(testId int64) []models.Question{
	result :=[]models.Question{}
	us.c.Find(bson.M{"testid": testId}).All(&result)
	return result
}

func (us *QuestionService) Add(q *models.Question) {
	us.mutex.Lock()
	q.Id = GetIdsService().GetNext(us.coll)
	us.c.Insert(q)
	us.mutex.Unlock()
}

func InitQuestionService(session *mgo.Session, db *mgo.Database) {
	if _questionServiceInstance == nil {
		_questionServiceInstance = &QuestionService{}
		_questionServiceInstance.Init(session, db, QUESTION_COLLECTION)
	}
}

func GetQuestionService() *QuestionService {
	return _questionServiceInstance
}



