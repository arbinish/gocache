package golinks

import (
	"log"
	"encoding/gob"
	"fmt"
	"bytes"
	"os"
	"errors"
)

const FNAME = "GOLINKS.gob"
var ErrNotFound = errors.New("NotFound")
var state = make(map[string]string)

func init() {
    fd, err := os.Open(FNAME)
    if err != nil {
        return
    }
    defer fd.Close()
    dec := gob.NewDecoder(fd)
    err = dec.Decode(&state)
    if err != nil {
        log.Println("Unable to init state")
    } else {
        log.Println("initialization completed")
    }
}

func Dump() string {
    var buf bytes.Buffer
    for k,v := range state {
        buf.WriteString(fmt.Sprintf("%s -> %s\n", k, v))
    }
    return buf.String()
}

func Get(key string) (response string, err error) {
    val, ok := state[key]
    if ok {
        response = val
    } else {
        err = ErrNotFound
    }
    return
}


func Set(key, value string) (err error) {
    state[key] = value
    defer func() {
        if r := recover(); r != nil {
            var ok bool
//            Note that err should not be reinitialized here
            err, ok = r.(error)
            if !ok {
                log.Println("Set failed for key", key, ":", r) 
            }
        }
    }()
    return
}


func Save() (err error) {
    fd, err := os.Open(FNAME)
    if err != nil {
        err = os.Remove(FNAME)
    }
    fd, err = os.Create(FNAME)
    if err != nil {
        log.Printf("Cannot Save state: %s\n", err)
        return
    } 
    defer fd.Close()
    enc := gob.NewEncoder(fd)
    enc.Encode(state)
    return
}


