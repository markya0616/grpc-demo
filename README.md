## 1. Build proto
protoc -I. --go_out=plugins=grpc:. api/pb/api.pro

## 3. Build client/server
go build -i -v -o bin/client github.com/markya0616/grpc-demo/cmd/client
go build -i -v -o bin/server github.com/markya0616/grpc-demo/cmd/server

## 4. Build Restful interface by grpc-gateway
protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/pb/api.proto

## 5. Build proxy
go build -i -v -o bin/proxy github.com/markya0616/grpc-demo/cmd/proxy
bin/proxy

curl -s -X GET -H "Content-Type: application/json" -H "Cache-Control: no-cache" "http://localhost:5454/v1/blocks/2"

curl -s -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -d '{
  "number": 1
}' "http://localhost:5454/v1/subscribe"
