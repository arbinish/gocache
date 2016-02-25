# golinks

A simple key value store with Rest End points.

## Supported operations:

- Get
    curl http://localhost:8080/api/get/key

- Set
    curl -d value -X POST http://localhost:8080/api/set/key


### How to run

- Install Go1.5+ SDK
- Clone repo
- env GOPATH=$PWD go build src/server.go

A pre-built mac [server.macosx] and linux [server.linux] binaries are already 
available in the repo, for your testing.

Enjoy!!
