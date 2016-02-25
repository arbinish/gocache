package main

import (
	"golinks"
	"log"
)

func main() {
	keys := []string{"key", "yubikey", "foo", "bar"}
	for _, k := range keys {
		v, _ := golinks.Get(k)
		log.Println(k, "->", v)
	}
}
