export CGO_ENABLED=0 
export GOOS=windows
export GOARCH=amd64
go build -o Bilimcp.exe cmd/stdio/main.go

export CGO_ENABLED=0 
export GOOS=linux
export GOARCH=amd64
go build -o Bilimcp_linux cmd/stdio/main.go