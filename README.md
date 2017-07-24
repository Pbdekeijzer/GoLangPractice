# Json-RestAPI
This is a json rest api which implements a variety of example requests.
The api is fully functional when using postman.

# Running the solution
from ../src/
run the command:
- go get github.com/Pbdekeijzer/Json-RestAPI

After the packages are finished cloning, in
..src/github.com/Pbdekeijzer/Json-RestAPI
run these commands:
- go build
- Json-RestAPI.exe

# Running tests
from ..src/github.com/Pbdekeijzer/Json-RestAPI
run the command:
- go test -v ./...


Postman tests can be found in Tests.postman_collection.json

# To do:
- Implement forms for posts
- Implement authentication

# Implemented external packages:
- https://github.com/gorilla/mux
- https://github.com/urfave/negroni
- https://github.com/stretchr/testify

