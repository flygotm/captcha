package captcha

import (
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/log"
	"github.com/billcoding/flygo/middleware"
)

//Define captcha struct
type captcha struct {
	logger     log.Logger
	sessionKey string
	width      int
	height     int
	length     int
	generator  Generator
}

//New
func New() *captcha {
	return &captcha{
		logger:     log.New("[Captcha]"),
		sessionKey: "Captcha",
		width:      200,
		height:     80,
		length:     4,
		generator:  &defaultGenerator{},
	}
}

//Name
func (cc *captcha) Name() string {
	return "Captcha"
}

//Type
func (cc *captcha) Type() *middleware.Type {
	return middleware.TypeHandler
}

//Method
func (cc *captcha) Method() middleware.Method {
	return middleware.MethodGet
}

//Pattern
func (cc *captcha) Pattern() middleware.Pattern {
	return "/captcha/rand"
}

//Handler
func (cc *captcha) Handler() func(c *c.Context) {
	return func(c *c.Context) {
		rands := cc.generator.generate(cc.length)
		bytes := imgText(cc.width, cc.height, rands)
		c.PNG(bytes)
		session := middleware.GetSession(c)
		if session != nil {
			session.Set(cc.sessionKey, rands)
		} else {
			cc.logger.Warn("session is nil")
		}
	}
}

//SessionKey
func (cc *captcha) SessionKey(sessionKey string) *captcha {
	cc.sessionKey = sessionKey
	return cc
}

//Width
func (cc *captcha) Width(width int) *captcha {
	cc.width = width
	return cc
}

//Height
func (cc *captcha) Height(height int) *captcha {
	cc.height = height
	return cc
}

//Length
func (cc *captcha) Length(length int) *captcha {
	cc.length = length
	return cc
}

//Generator
func (cc *captcha) Generator(generator Generator) *captcha {
	cc.generator = generator
	return cc
}
