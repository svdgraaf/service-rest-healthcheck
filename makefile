VERSION=0.0.2
GOPATH=$(CURDIR)

default: main.go
	# darwin
	GOOS=darwin GOARCH=386 go build -v -o bin/service-health-darwin-$(VERSION)
	# windows
	GOOS=windows GOARCH=386 go build -o bin/service-health-winx86-$(VERSION).exe

dep:
	go get -d
