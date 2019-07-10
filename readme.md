# CS1632_Deliverable4

This repo contains the system under test for CS1632's deliverable 4 at the University of Pittsburgh. This application is a web service written in Go. Service oriented and efficiency oriented performance tests should be executed against this web service to determine its key performance indicators.

## Installation instructions

1. Install Go. https://golang.org/dl/
2. Clone this repo.
3. Open a command prompt in the repo's directory and execute the command `go build main.go`.
4. From the command prompt run the generated executable. i.e. `main.exe` `./main`.

This service will respond to the following request:

`curl -X GET \
  http://localhost:8080 \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:8080' \
  -H 'Postman-Token: e31254d1-e02a-4986-b482-b1bbe6d35157,59c7c5cc-5a0e-4461-9b57-8c83705096ac' \
  -H 'User-Agent: PostmanRuntime/7.15.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 16' \
  -d '{ "input": 100 }'`
