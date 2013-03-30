/**
 * Created with IntelliJ IDEA.
 * User: liangdi
 * Date: 2/2/13
 * Time: 2:07 AM
 * To change this template use File | Settings | File Templates.
 */
package controllers

import "github.com/robfig/revel"


func init(){
	revel.RegisterPlugin(MongodbPlugin{})
}

