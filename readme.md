# CS1632_Deliverable4

This repo contains the system under test for CS1632's deliverable 4 at the University of Pittsburgh. This application is a web service written in Go. Service oriented and efficiency oriented performance tests should be executed against this web service to determine its key performance indicators.

This service will respond to the following requests:

`curl -X GET \
  http://localhost:8080 \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: localhost:8080' \
  -H 'Postman-Token: 5d1d285f-bb64-4a33-9e3e-074dd07d167d,23e3defe-ef7d-4f31-aec5-83f930098198' \
  -H 'User-Agent: PostmanRuntime/7.15.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache'`

`curl -X POST \
  http://localhost:8080 \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:8080' \
  -H 'Postman-Token: 2c3f6567-7024-4c0f-9fcd-02dc9101630e,5becb1db-ee87-477f-b854-8e439f9431d8' \
  -H 'User-Agent: PostmanRuntime/7.15.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 54' \
  -d '{
	"sender": "TestSender",
	"message": "TestMessage"
}'`

`curl -X DELETE \
  http://localhost:8080 \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:8080' \
  -H 'Postman-Token: e51f0bc1-0063-40ea-bb5f-73994e35d2ed,b315ef2f-e96c-48a4-a28f-53b84aa2b33c' \
  -H 'User-Agent: PostmanRuntime/7.15.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 12' \
  -d '{
	"id": 1
}'`

