export GOPATH=$(CURDIR)

default: main.go
	go build -v -o service-health
	GOOS=windows GOARCH=386 go build -o service-health.exe

run: default
	./service-health.exe
