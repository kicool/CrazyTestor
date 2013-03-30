/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 2/2/13
 * Time: 2:07 AM
 * To change this template use File | Settings | File Templates.
 */
package controllers

import "github.com/robfig/revel"
import "net/http"
import "log"

func init(){
	revel.RegisterPlugin(MongodbPlugin{})
	RunContextMgr()
	go listen()
}

func listen(){
	http.HandleFunc("/", receiveHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
