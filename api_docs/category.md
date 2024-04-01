# Category API Spec

## Create Category API

Endpoint: POST /api/categories

Headers :
- Authorization : Bearer Token

Request Body : 
```json
{
    "name": "Bermain"
}
```

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "category-11111",
        "name": "Bermain"
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

## Update Category API

Endpoint: PUT /api/categories/:categoryId

Headers :
- Authorization : Bearer Token

Request Body : 
```json
{
    "categoryName": "Bermain 2"
}
```

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "category-11111",
        "name": "Bermain 2"
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

## Get Category API

Endpoint: GET /api/categories/:categoryId

Headers :
- Authorization : Bearer Token

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "category-11111",
        "name": "Bermain 2"
    }
}
```

Response Body Error : 
```json
{
    "code": 404,
    "status": "NOT_FOUND",
    "data": "category is not found"
}
```

## List Category API

Endpoint: GET /api/categories

Headers :
- Authorization : Bearer Token

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "id": "category-11111",
            "name": "Bermain 2"
        }
    ]
}
```

Response Body Error : 
```json
{
    "code": 401,
    "status": "UNAUTHORIZED",
    "data": {
        "message": "invalid or expired token"
    }
}
```

## Remove Category API

Endpoint: DELETE /api/categories/:categoryId

Headers :
- Authorization : Bearer Token

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "category-22222",
        "name": "Bermain 1"
    }
}
```

Response Body Error : 
```json
{
    "code": 404,
    "status": "NOT_FOUND",
    "data": "category is not found"
}
```






