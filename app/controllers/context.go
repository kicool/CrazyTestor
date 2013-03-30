package controllers

import (
	"fmt"
	"strconv"
)

var (
	cmgr        *ContextMgr = nil
	DEFAULT_LEN             = 100
)

type Message map[string]string

type Context struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        int64
}

func NewContext(s *Message) *Context {
	var err error
	c := Context{ToUserName: (*s)["ToUserName"],
		FromUserName: (*s)["FromUserName"],
		CreateTime:   (*s)["CreateTime"],
		MsgType:      (*s)["CreateTime"],
		Content:      (*s)["Content"]}
	c.MsgId, err = strconv.ParseInt((*s)["MsgId"], 10, 64)
	if err != nil {
		panic("strconv")
	}

	return &c
}

func (c *Context) sameUser(s *Message) bool {
	if c.FromUserName == (*s)["FromUserName"] {
		return true
	}
	return false
}

func (c *Context) handle(s *Message) {
	switch (*s)["Content"] {
	case "?":
		fmt.Println("start test")

	default:
	}
}

type ContextMgr struct {
	ctxs []*Context
	in   chan *Message
	out  chan *Message
}

func Invoke(s *map[string]string) *map[string]string {
	(*cmgr).in <- (*Message)(s)
	omsg := <-(*cmgr).out
	return (*map[string]string)(omsg)
}

func RunContextMgr() {
	if cmgr == nil {
		cmgr = new(ContextMgr)
		cmgr.ctxs = make([]*Context, DEFAULT_LEN, DEFAULT_LEN)
		cmgr.in = make(chan *Message)
		cmgr.out = make(chan *Message)
		fmt.Println("ContextMgr created")
	}

	go func() {
		fmt.Println("ContextMgr Loop")
		for { //TODO for quite
			select {
			case s := <-cmgr.in:
				var i int
				var k *Context
				for i, k = range cmgr.ctxs {
					if k != nil && k.sameUser(s) {
						fmt.Println("Message Exit")
						go k.handle(s)
						goto NEXT
					}
				}
				fmt.Println("Message not Exit", i)
				for i, k = range cmgr.ctxs {
					if k == nil {
						cmgr.ctxs[i] = NewContext(s)
						go cmgr.ctxs[i].handle(s)
					} else {
						panic("Out of slice len")
					}
				}
			}
		NEXT:
		}
	}()
}
