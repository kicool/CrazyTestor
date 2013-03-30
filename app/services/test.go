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
	_testServiceInstance *TestService
)

const (
	TEST_COLLECTION = "tests"
)

type TestService struct {
	MongodbService
}

func (us *TestService) Get(id int64) string {
	return "test:" + strconv.Itoa(int(id))
}

func (us *TestService) Count() int {
	ret := 0
	ret, _ = us.c.Count()
	return ret
}

func (us *TestService) GetOne() *models.Test {
	result := &models.Test{}
	us.c.Find(nil).One(result)
	return result
}

func (us *TestService) Add(test *models.Test) {
	us.mutex.Lock()
	test.Id = GetIdsService().GetNext(us.coll)
	us.c.Insert(test)
	us.mutex.Unlock()
}

func (us *TestService) Find() []models.Test{
	result :=[]models.Test{}
	us.c.Find(nil).All(&result)
	return result
}

func InitTestService(session *mgo.Session, db *mgo.Database) {
	if _testServiceInstance == nil {
		_testServiceInstance = &TestService{}
		_testServiceInstance.Init(session, db, TEST_COLLECTION)
	}
}

func GetTestService() *TestService {
	return _testServiceInstance
}



