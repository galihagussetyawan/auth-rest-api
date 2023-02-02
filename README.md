# REST API EXAMPLE REQUEST

## How to read examples

The first line of each examples show the verb GET, POST and etc. The portion of the URI that describes the resource and the REST API version number. For example, for signing in the example URI shows this:

`POST http://localhost:8080/api/v1/auth/login`

This indicates that you should make a POST request, using version v1 of the REST API.

<br>

## Register Account

Request:
```
POST http://localhost:8080/api/v1/user
Content-Type: application/json

{
  "firstname": "your-firstname",
  "lastname": "your-lastname",
  "email": "your-email@gmail.com",
  "password": "your-secret-password"
}
```

Response:
```
HTTP/1.1 201 Created
Content-Type: application/json

{
  "data": {
    "email": "your-email@gmail.com",
    "firstname": "your-firstname",
    "id": "63db59f4bad9c834019a498a",
    "lastname": "your-lastname"
  },
  "message": "success to create an account",
  "status_code": 201
}
```