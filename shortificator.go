package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type url struct {
	global_map map[string]string
}

func (u url) Shorten(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	shorten_url := "http:\\\\iq.xi\\" + hex.EncodeToString(h.Sum(nil)[:5])
	u.global_map[shorten_url] = s
	return shorten_url
}

func (u url) Resolve(s string) string {
	if original_url, ok := u.global_map[s]; ok == true {
		return string(original_url)
	} else {
		return ""
	}
}

func main() {
	link := "http:\\\\google.com"
	var u Shortener = url{global_map: make(map[string]string)}
	var short_u = u.Shorten(link)
	var long_u = u.Resolve(short_u)
	fmt.Println(short_u)
	fmt.Println(long_u)
}
