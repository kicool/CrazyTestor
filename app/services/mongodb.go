/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 2/2/13
 * Time: 3:49 AM
 * To change this template use File | Settings | File Templates.
 */
package services

import (
	"labix.org/v2/mgo"
	"sync"
	"labix.org/v2/mgo/bson"
	"log"
)

type MongodbService struct {
	session *mgo.Session
	db *mgo.Database
	c  *mgo.Collection
	coll string

	//sync
	mutex *sync.Mutex
}

func (ms *MongodbService) Init(session *mgo.Session,db *mgo.Database,coll string){
	ms.session = session
	ms.db = db;
	ms.c = db.C(coll)
	ms.coll = coll
	ms.mutex =new(sync.Mutex)
}

var (
	_idsService * IdsService
)
const (
	IDS_COLLECTION = "ids"
)

func InitIdsService(session *mgo.Session,db *mgo.Database){
	if _idsService == nil {
		_idsService = &IdsService{}
		_idsService.Init(session,db,IDS_COLLECTION)
	}
}

func GetIdsService() *IdsService {
	return _idsService
}

type IdsService struct {
	MongodbService
}

type Ids struct {

	Coll string
	Id	int64
}

func (is IdsService) GetNext(coll string) int64 {
	var id int64
   	is.mutex.Lock()

	_,err:=is.c.Upsert(&bson.M{"coll":coll},&bson.M{"$inc":&bson.M{"id":1}})
	if err!=nil{
		log.Fatal(err)
	}
	ids := Ids{}

	err = is.c.Find(&bson.M{"coll":coll}).One(&ids)
	if err!= nil {
		log.Fatal(err)
	}
	id = ids.Id
	is.mutex.Unlock()
	return id
}

