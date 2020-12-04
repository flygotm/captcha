package captcha

import (
	"fmt"
	"github.com/billcoding/flygo"
)

//Define captcha struct
type captcha struct {
	sessionKey string
	width      int
	height     int
	length     int
	Generator  Generator
}

func New() *captcha {
	return &captcha{
		sessionKey: "FlygoCaptcha",
		width:      200,
		height:     80,
		length:     4,
		Generator:  &defaultGenerator{},
	}
}

func (cc *captcha) SessionKey(sessionKey string) *captcha {
	cc.sessionKey = sessionKey
	return cc
}

func (cc *captcha) Width(width int) *captcha {
	cc.width = width
	return cc
}

func (cc *captcha) Height(height int) *captcha {
	cc.height = height
	return cc
}

func (cc *captcha) Length(length int) *captcha {
	cc.length = length
	return cc
}

func (cc *captcha) Method() string {
	return "GET"
}

func (cc *captcha) Fields() []*flygo.Field {
	return nil
}

func (cc *captcha) Name() string {
	return "FlygoCaptcha"
}

func (cc *captcha) Pattern() string {
	return "/captcha/rand"
}

func (cc *captcha) Process() flygo.Handler {
	return func(c *flygo.Context) {
		rands := cc.Generator.generate(cc.length)
		bytes := imgText(cc.width, cc.height, rands)
		c.PNG(bytes)
		session := c.Session
		if session == nil {
			fmt.Println("session is nil")
		} else {
			session.Set(cc.sessionKey, rands)
		}
	}
}
