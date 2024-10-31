package all
protobuf的底层实现
https://github.com/protocolbuffers/protobuf/releases
protoc -I . goods.proto --go_out=plugins=grpc:.
protoc -I . helloworld.proto --go_out=plugins=grpc:.
protoc -I . helloworld.proto --go_out=. --go-grpc_out=.
protoc -I . --go_out=plugins=grpc:. --validate_out="lang=go:." form.proto
"google.golang.org/grpc/metadata"
go-grpc-middleware
"google.golang.org/grpc/status"
google.golang.org/grpc/codes
