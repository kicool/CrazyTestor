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
	_questionServiceInstance *QuestionService
)

const (
	QUESTION_COLLECTION = "tests"
)

type questionService struct {
	MongodbService
}

func (us *questionService) Get(id int64) string {
	return "test:" + strconv.Itoa(int(id))
}

func (us *questionService) Count() int {
	ret := 0
	ret, _ = us.c.Count()
	return ret
}

func (us *questionService) GetOne() *models.Test {
	result := &models.Test{}
	us.c.Find(nil).One(result)
	return result
}

func (us *questionService) Add(test *models.Test) {
	us.mutex.Lock()
	test.Id = GetIdsService().GetNext(us.coll)
	us.c.Insert(test)
	us.mutex.Unlock()
}

func InitQuestionService(session *mgo.Session, db *mgo.Database) {
	if _questionServiceInstance == nil {
		_questionServiceInstance = &questionService{}
		_questionServiceInstance.Init(session, db, TEST_COLLECTION)
	}
}

func GetQuestionService() *questionService {
	return _questionServiceInstance
}



