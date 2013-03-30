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
	"iia/app/services"
)

var (
	session *mgo.Session
	db      *mgo.Database
	dbHost  = "localhost"
	dbName  = "revel_test"

	idsService *services.IdsService
	userService *services.UserService

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
	services.InitUserService(session,db)
	services.InitIdsService(session,db)

	userService = services.GetUserService()
	idsService = services.GetIdsService()
}

