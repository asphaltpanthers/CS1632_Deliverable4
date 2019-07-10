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
  -H 'Postman-Token: 5d1d285f-bb64-4a33-9e3e-074dd07d167d,23e3defe-ef7d-4f31-aec5-83f930098198,53af1720-603d-4a71-8e07-14a31abf59c1' \
  -H 'User-Agent: PostmanRuntime/7.15.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache,no-cache' \
  -d '{
	"input": 100
}'`