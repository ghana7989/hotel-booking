# Hotel reservation backend

# Project environment variables
```
HTTP_LISTEN_ADDRESS=:3000
JWT_SECRET=somethingsupersecretthatNOBODYKNOWS
MONGO_DB_NAME=hotel-booking
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

## Project outline
- users -> book room from an hotel 
- admins -> going to check reservation/bookings 
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
### Mongodb driver 
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber 
Documentation
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```

### Required GO global packages

To run the tests in watch mode you need https://github.com/gotestyourself/gotestsum this package installed globally and then run `make test-watch` to run the tests in watch mode.
To install the gotestsum package run `go install gotest.tools/gotestsum@latest` and make sure the `$GOPATH/bin` is in your $PATH.

### Swagger UI

Once the server is running go to /swagger/index.html to see the swagger UI.
