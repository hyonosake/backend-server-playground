# Simple web server
### This project provides a simple Web-server to play around with

## Install
#### Simply run this command in your terminal:
```
git clone https://github.com/hyonosake/backend-server-playground

cd backend-server-playground 
./web-server
```
It will automatically run your web-server. 

## Requisites
 There are none - you can simply run web-server using  
 
However, if you want to experiment and get more control over web-server, use `make build` command instead, followed by `make run` - it will do pretty much the same (`Go > v.1.20` required locally)


## How to play with
Web-server is available on host `127.0.0.1`, port `123321`
Currently only server-side is supported, i.e. server has some endpoints for requests, but he's not the one who initiates connection


# Availible Requests 

---

## GET /internal/healthcheck
Request is used to check server state
#### Request
* Content-type: application/json
* Accepted content: any
* Accepted methods: any
```json
{}
```
#### Response
```json
{
    "alive": true
}
```
* Accepted Content-type: application/json
* Accepted Headers: `statusOK(200)`
 
---

## POST /orders/simple_order
Request is used to create a simple order
#### Request
* Accepted Content-type: application/json. Server will return `BadRequest(400)` and Error message otherwise
* Accepted methods: `POST`
```json
{
  "productId": 12345
}
```
#### Response
* Content-type: `application/json`
* Possible Headers:
  * `statusOK(200)` - everything went smooth
  * `BadRequest(400)` - error with provided request (invalid content-type, or request method, or data provided)
  * `InternalServerError(500)` - something went wrong on server-side :(

Response will contain either result object or error object, depending on status of processed request
#### Successful response
```json
{
  "result": {
    "message": "some msg",
    "orderId": 1
  }
}
```
#### Error response
```json
{
  "error": {
    "message": "content type must be application/json, 'text/plain' provided"
  }
}
```







