/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 2/2/13
 * Time: 1:18 AM
 * To change this template use File | Settings | File Templates.
 */
package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"log"
	"CrazyTestor/app/services"
)

var (
	session *mgo.Session
	db      *mgo.Database
	dbHost  = "localhost"
	dbName  = "revel_test"

	idsService *services.IdsService
	testService *services.TestService
	questionService *services.QuestionService
	answerService *services.AnswerService

)

type MongodbPlugin struct {
	revel.EmptyPlugin
}

func (p MongodbPlugin) OnAppStart() {
	var err error
	var confFound bool
	var readDbHost,readDbName string

	readDbHost,confFound =revel.Config.String("dbHost")
	if confFound {
		dbHost = readDbHost
	}

	readDbName,confFound = revel.Config.String("dbName");
	log.Println(readDbName)
	if confFound{
		dbName = readDbName
	}

	log.Println("dbHost:" + dbHost)
	log.Println("dbName:" + dbName)


	session, err = mgo.Dial(dbHost)
	db = session.DB(dbName)
	if err != nil {
		log.Fatal(err)
	}

	// init services
	//services.InitUserService(session,db)
	services.InitIdsService(session,db)
	services.InitTestService(session,db)
	services.InitQuestionService(session,db)
	services.InitAnswerService(session,db)
	//userService = services.GetUserService()
	idsService = services.GetIdsService()
	testService = services.GetTestService()
	questionService = services.GetQuestionService()
	answerService = services.GetAnswerService()
}

