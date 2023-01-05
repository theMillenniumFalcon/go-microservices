A golang based microservice

To compile proto files
```
protoc --go_out=. --go-grpc_out=. ./messages/*.proto
```

To generate the JWT_SECRET
```
require('crypto').randomBytes(35).toString("hex")
```

To make a mongo DB database
```
docker run -it --rm --name mongodb_container -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -v mongodata:/data/db -d -p 27017:27017 mongo
```