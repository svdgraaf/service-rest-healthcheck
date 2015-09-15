VERSION=0.0.3
GOPATH=$(CURDIR)

default: main.go
	# darwin
	# GOOS=darwin GOARCH=386 go build -v -o bin/service-health-darwin-$(VERSION)
	# windows
	GOOS=windows GOARCH=386 go build -o bin/service-rest-healthcheck-winx86-$(VERSION).exe
	GOOS=windows GOARCH=amd64 go build -o bin/service-rest-healthcheck-winamd64-$(VERSION).exe

dep:
	go get -d
