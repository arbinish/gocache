all: linux mac

linux:
	env GOOS=linux GOARCH=amd64 GOPATH=${PWD} go build -o server.linux src/server.go 

mac:
	env GOOS=darwin GOARCH=amd64 GOPATH=${PWD} go build -o server.macosx src/server.go 
