# REST API EXAMPLE REQUEST

## How to read examples

The first line of each examples show the verb GET, POST and etc. The portion of the URI that describes the resource and the REST API version number. For example, for signing in the example URI shows this:

`POST http://localhost:8080/api/v1/auth/login`

This indicates that you should make a POST request, using version v1 of the REST API.

<br>

## Register Account

**Request:**
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

**Success Response:**
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
**Possible errors:**

| Error Code | Description |
| -----------| ----------- |
| 400 Bad Request | Required fields were invalid, not specified. |

## Login

**Request:**
```
POST http://localhost:8080/api/v1/auth/login
Content-Type: application/json

{
  "identity": "your-email@gmail.com",
  "password": "your-secret-password"
}
```

**Success Response:**
```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImdhbGloZGV2ZWxvcG1lbnRAZ21haWwuY29tIiwiZXhwIjoxNjc1NDEzNDM4LCJpZCI6IjYzY2NmNTg0NDIxZWM2YmUwMjY3MzUwOSJ9.p1lmCG9336eH8EaPfc_Y4mz3CDb3QSab1JB5ZrfGsm0",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImdhbGloZGV2ZWxvcG1lbnRAZ21haWwuY29tIiwiZXhwIjoxNjc4MjA3MDM4LCJpZCI6IjYzY2NmNTg0NDIxZWM2YmUwMjY3MzUwOSJ9.CEFq7-emS2Ft1mZFuqgpDqiBRxSGFc1qfllUHWIjeDo"
  },
  "status_code": 200
}
```

**Error Response:**

```
HTTP/1.1 401 Unauthorized
Content-Type: application/json

{
  "message": "invalid password",
  "status_code": 401
}
```

**Possible errors:**

|Error Code|Description|
|----------|-----------|
|400 Bad Request|Required fields were invalid, not specified.|
|401 Unauthorized|Invalid user password|
|404 Not Found|The user account is not registered.|