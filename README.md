A golang based microservice

To compile proto files
```
protoc --go_out=. --go-grpc_out=. ./messages/*.proto
```

To generate the JWT_SECRET
```
require('crypto').randomBytes(35).toString("hex")
```