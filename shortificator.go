package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type url struct {
	global_map map[string]string
}

func (u url) Shorten(s string) string {
	shorten_url := ""
	temp_s := s
	for {
		h := md5.New()
		h.Write([]byte(temp_s))
		shorten_url = "http://iq.xi/" + hex.EncodeToString(h.Sum(nil)[:5])
		if _, ok := u.global_map[shorten_url]; ok {
			temp_s += string(rand.Int())
		} else {
			break
		}
	}

	u.global_map[shorten_url] = s

	return shorten_url
}

func (u url) Resolve(s string) string {
	return string(u.global_map[s])
}

func main() {

	var u Shortener = url{global_map: make(map[string]string)}
	link := "http://google.com"

	for i := 0; i < 10; i++ {
		var short_u = u.Shorten(link)
		var original_u = u.Resolve(short_u)
		fmt.Println("Short URL:", short_u, "Original URL:", original_u, "Count: ", i)
	}

}
