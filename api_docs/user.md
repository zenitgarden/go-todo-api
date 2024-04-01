# User API Spec

## Register User API

Endpoint: POST /api/users

Request Body : 
```json
{
    "username": "john",
    "password": "secret",
    "name": "John"
}
```

Response Body Success : 
```json
{   
    "code": 200,
    "status": "OK",
    "data": {
        "username": "john",
        "name": "John"
    }
}
```

Response Body Error : 
```json
{
    "code": 400,
    "status": "BAD_REQUEST",
    "data": {
        "name": "name is required",
        "password": "password is required",
        "username": "username is required"
    }
}
```

## Login User API

Endpoint: POST /api/users/login

Request Body : 
```json
{
    "username": "john",
    "password": "secret",
}
```

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
       "token": "token"
    }
}
```

Response Body Error : 
```json
{
    "code": 400,
    "status": "BAD_REQUEST",
    "data": "username or password wrong"
}
```

## Update User API

Endpoint : PUT /api/users

Headers :
- Authorization : Bearer Token

Request Body :
```json
{
    "name": "john wick"
}
```

Response Body Success :
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "username": "john",
        "name": "john wick"
    }
}
```

Response Body Error :
```json
{
    "code": 400,
    "status": "BAD_REQUEST",
    "data": {
        "name": "name is required"
    }
}
```