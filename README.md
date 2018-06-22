## 1. Build proto
protoc -I. --go_out=plugins=grpc:. api/pb/api.proto
