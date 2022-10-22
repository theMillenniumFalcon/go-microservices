A golang based microservice

To compile proto files
```
protoc --go_out=. --go-grpc_out=. ./messages/*.proto
```

To generate A JWT_SECRET_KEY
```
require('crypto').randomBytes(35).toString("hex")
```