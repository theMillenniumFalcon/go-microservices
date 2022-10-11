A golang based microservice

To compile proto files
```
protoc --go_out=. --go-grpc_out=. ./messages/*.proto
```