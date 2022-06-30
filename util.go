package httprelay

import (
	"math/rand"
	"time"
)

// https://stackoverflow.com/a/22892986/625521 ////////////////////////////////////////
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStr(n int) string {
	//return "12345678" // TODO: !!!!! restore
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//////////////////////////////////////////////////////////////////////////////////////
