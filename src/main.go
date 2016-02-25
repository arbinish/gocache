package main 

import (
	"log"
	"golinks"
)

func main() {
    keys := []string{"key", "yubikey", "foo", "bar"}
    for _, k := range keys {
        v, _ := golinks.Get(k)
        log.Println(k, "->", v)
    }
}
