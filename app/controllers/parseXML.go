package controllers

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/xml"
)

type SessionData struct {
	FromUserName string `xml:FromUserName`
	ToUserName string `xml:ToUserName`
	CreateTime string `xml:CreateTime`
	MsgType string `xml:MsgType`
	Content string `xml:Content`
	MsgId string `xml:MsgID`
}

func (data SessionData) ToMap() map[string]string {
	var result =  make(map[string] string)
	result["ToUserName"] = data.ToUserName
	result["FromUserName"] = data.FromUserName
	result["CreateTime"] = data.CreateTime
	result["MsgType"] = data.MsgType
	result["Content"] = data.Content
	result["MsgId"] = data.MsgId
	return result
}

func ToXML(rawMap *map[string]string) string{
	var result string = "<xml>";
	for  k,v:=range *rawMap{
		result += "<" + k + ">"
		result += v
		result += "</" + k + ">"
	}
	result += "</xml>"
	return result
}


func receiveHandler(w http.ResponseWriter, req *http.Request){
	log.Print("here")
	// deal input
	var data SessionData
	xmlValue, err :=ioutil.ReadAll(req.Body)
	if err != nil{
		fmt.Println(err)
	}
	defer req.Body.Close()

	xml.Unmarshal(xmlValue, &data)
	parsedMap := data.ToMap()
	output:= Invoke(&parsedMap)
//	fmt.Println("after parse:" , parsedMap)
	
	// deal output
	xmlStr:= ToXML(output)
	fmt.Fprint(w, xmlStr)
}

func main(){
	http.HandleFunc("/", receiveHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

