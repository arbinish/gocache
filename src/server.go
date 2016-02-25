package main 

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "golinks"
    "io/ioutil"
    "time"
    "log"
)


var sync = make(chan bool)

func GetKey(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    name := urlParams["key"]
    val, err := golinks.Get(name)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, err.Error())
    } else {
        fmt.Fprintf(w, val)
    }
}

func SetKey(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    name := urlParams["key"]
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
	    w.WriteHeader(http.StatusBadRequest)
	    fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, string(payload))
		golinks.Set(name, string(payload))    
		go func() {
		    sync <- true
		}()
	}	  
}

func InitSaveCycle(saveChan <-chan bool) {
    // Throttle disk writes
    ticker := time.Tick(5 * time.Second)
    log.Println("Initialized background sync thread ...")
    for _ = range saveChan {
        <-ticker
        log.Println("Saving to disk ... ")
        golinks.Save()
    }
    log.Println("Sync thread completed.")    
}


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/api/get/{key}", GetKey)
    router.HandleFunc("/api/set/{key}", SetKey).Methods("POST")
    http.Handle("/", router)

	
    // 	persist save
    go InitSaveCycle(sync)
    
    defer func() {
        close(sync)
    }()
    
    http.ListenAndServe(":8080", nil)
}

