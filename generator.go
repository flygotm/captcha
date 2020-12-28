package captcha

import (
	"math/rand"
	"strings"
	"time"
)

//Generate interface
type Generator interface {
	//generate
	generate(length int) string
}

type defaultGenerator struct {
}

func (d *defaultGenerator) generate(length int) string {
	randStr := func(length int) string {
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		strs := make([]string, length)
		for i := 0; i < length; i++ {
			rand.Seed(time.Now().UnixNano())
			strs[i] = string(chars[rand.Intn(len(chars))])
		}
		return strings.Join(strs, "")
	}
	return randStr(length)
}
