# gocache

A simple key value store with Rest end points.

## Supported operations:

- Get
    curl http://localhost:8080/api/get/key

- Set
    curl -d value -X POST http://localhost:8080/api/set/key


### How to run

- Install Go1.5+ SDK
- Clone repo
- run make. either `make mac` or `make linux` depending on your platform

A pre-built mac [server.macosx] and linux [server.linux] binaries are already 
available in the repo, for your testing.

Enjoy!!
