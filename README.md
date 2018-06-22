## 1. Build proto
protoc -I. --go_out=plugins=grpc:. api/pb/api.pro

## 3. Build client/server
go build -i -v -o bin/client github.com/markya0616/grpc-demo/cmd/client
go build -i -v -o bin/server github.com/markya0616/grpc-demo/cmd/server
