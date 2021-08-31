生成客户端代码和服务端代码命令：
```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --proto_path=$(go env GOPATH)/src/github.com/zbwang163/idl\
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ad_lego_overpass.proto
```