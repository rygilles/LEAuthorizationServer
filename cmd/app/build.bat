set GOARCH=amd64
set GOOS=linux
#go tool dist install -v pkg/runtime
#go install -v -a std

set GOARCH=amd64
set GOOS=linux
go build -o ../../builds/lin64_server main.go

set GOARCH=386
set GOOS=linux
go build -o ../../builds/lin386_server main.go

set GOARCH=amd64
set GOOS=windows
go build -o ../../builds/win64_server.exe main.go

set GOARCH=386
set GOOS=windows
go build -o ../../builds/win386_server.exe  main.go

