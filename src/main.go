package main

import (
	"gocache"
	"log"
)

func main() {
	keys := []string{"key", "yubikey", "foo", "bar"}
	for _, k := range keys {
		v, _ := gocache.Get(k)
		log.Println(k, "->", v)
	}
}
