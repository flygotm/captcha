package captcha

import (
	"github.com/billcoding/flygo"
	"strings"
)

//return current session captcha code
func (cc *captcha) Current(c *flygo.Context) string {
	if c.Session == nil {
		return ""
	}
	captchaCode := c.Session.Get(cc.sessionKey)
	if captchaCode == nil {
		return ""
	}
	return captchaCode.(string)
}

//return captcha equals
func (cc *captcha) Equals(c *flygo.Context, val string, ignoreCase bool) bool {
	current := cc.Current(c)
	if current == "" {
		return false
	}
	if val == "" {
		return false
	}
	if ignoreCase {
		return strings.ToUpper(current) == strings.ToUpper(val)
	}
	return current == val
}

//return captcha match
func (cc *captcha) Match(c *flygo.Context, ignoreCase bool) bool {
	return cc.Equals(c, c.Param(cc.sessionKey), ignoreCase)
}

//clear captcha
func (cc *captcha) Clear(c *flygo.Context) {
	c.Session.Del(cc.sessionKey)
}
