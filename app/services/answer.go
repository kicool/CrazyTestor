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
	"strconv"
	"CrazyTestor/app/models"
)

var (
	_answerServiceInstance *AnswerService
)

const (
	ANSWER_COLLECTION = "answers"
)

type AnswerService struct {
	MongodbService
}

func (us *AnswerService) Get(id int64) string {
	return "test:" + strconv.Itoa(int(id))
}

func (us *AnswerService) Count() int {
	ret := 0
	ret, _ = us.c.Count()
	return ret
}

func (us *AnswerService) GetOne() *models.Answer {
	result := &models.Answer{}
	us.c.Find(nil).One(result)
	return result
}

func (us *AnswerService) Add(answer *models.Answer) {
	us.mutex.Lock()
	answer.Id = GetIdsService().GetNext(us.coll)
	us.c.Insert(answer)
	us.mutex.Unlock()
}

func InitAnswerService(session *mgo.Session, db *mgo.Database) {
	if _answerServiceInstance == nil {
		_answerServiceInstance = &AnswerService{}
		_answerServiceInstance.Init(session, db, TEST_COLLECTION)
	}
}

func GetAnswerService() *AnswerService {
	return _answerServiceInstance
}



