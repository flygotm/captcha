package captcha

import (
	"github.com/billcoding/flygo"
	"testing"
)

func TestNew(t *testing.T) {
	cc := New()
	app := flygo.GetApp().Use(cc)
	//app.Config.Flygo.Dev.Debug = true
	//app.SessionConfig.Enable = true
	//app.SessionConfig.Created = func(s flygo.Session) {
	//	fmt.Println(s.Id())
	//}
	//app.SessionConfig.SessionProvider = redisprovider.New(&redis.Options{
	//	Addr:     "139.196.40.100:6379",
	//	Password: "OQYG22dfd45gfgfgfB84V",
	//	DB:       14,
	//})
	app.Get("/", func(c *flygo.Context) {
		if cc.Match(c, true) {
			cc.Clear(c)
		}
	})
	app.Get("/invalidate", func(c *flygo.Context) {
		c.Session.Invalidate()
	})
	app.Run()
}
