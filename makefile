export GOPATH=$(pwd)

default: main.go
	# darwin
	GOOS=darwin GOARCH=386 go build -v -o build/service-health-darwin
	# windows
	GOOS=windows GOARCH=386 go build -o build/service-health-windows-x86.exe

dep:

	go get -d
