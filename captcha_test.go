package captcha

import (
	"github.com/billcoding/flygo"
	c "github.com/billcoding/flygo/context"
	"github.com/billcoding/flygo/middleware"
	"github.com/billcoding/flygo/session"
	"github.com/billcoding/flygo/session/memory"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	app := flygo.GetApp()
	cc := New()
	app.Use(cc)
	app.UseSession(memory.Provider(), &session.Config{Timeout: time.Hour * 12}, nil)
	app.GET("/get", func(c *c.Context) {
		c.Text(middleware.GetSession(c).Get(cc.sessionKey).(string))
	})
	app.GET("/clear", func(c *c.Context) {
		if cc.Match(c, true) {
			cc.Clear(c)
		}
	})
	app.GET("/invalidate", func(c *c.Context) {
		middleware.GetSession(c).Invalidate()
	})
	app.UseRecovery()
	app.Run()
}
