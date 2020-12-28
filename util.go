package captcha

import (
	"github.com/billcoding/calls"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/middleware"
	"strings"
)

//return current session captcha code
func (cc *captcha) Current(c *c.Context) (str string) {
	calls.NNil(middleware.GetSession(c), func() {
		str = middleware.GetSession(c).Get(cc.sessionKey).(string)
	})
	return str
}

//return captcha equals
func (cc *captcha) Equals(c *c.Context, val string, ignoreCase bool) bool {
	current := cc.Current(c)
	if current == "" {
		return false
	}
	if val == "" {
		return false
	}
	return current == val || (ignoreCase && strings.EqualFold(current, val))
}

//return captcha match
func (cc *captcha) Match(c *c.Context, ignoreCase bool) bool {
	return cc.Equals(c, c.Param(cc.sessionKey), ignoreCase)
}

//clear captcha
func (cc *captcha) Clear(c *c.Context) {
	calls.NNil(middleware.GetSession(c), func() {
		middleware.GetSession(c).Del(cc.sessionKey)
	})
}
