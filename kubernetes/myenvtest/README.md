export GOPATH=...
export GO111MODULE=on
go mod init example.com/main
go mod download
go mod vendor
go mod tidy
