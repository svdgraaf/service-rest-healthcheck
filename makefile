export GOPATH=$(CURDIR)

default: main.go
	go get -d
	go build -v -o build/service-health
	GOOS=windows GOARCH=386 go build -o build/service-health.exe

dep:
	#go get golang.org/x/sys/windows/svc/mgr
	go get github.com/golang/...
