# Todo API Spec

## Create Todo API

Endpoint: POST /api/todos

Headers :
- Authorization : Bearer token

Request Body : 
```json
{
    "title": "Berenang di ujung genteng",
    "description": "nemu ikan di laut",
    "categories": [     // optional
            {
                "id": "category-11111", 
            }
    ]
}
```

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "todo-11111",
        "title": "Berenang di ujung genteng",
        "description": "nemu ikan di laut",
        "categories": [
            {
                "id": "category-11111",
                "name": "Bermain 2"
            }
        ],
        "created_at": "2024-04-01T13:58:23Z",
        "updated_at": "2024-04-01T13:58:23Z"
    }
}
```

Response Body Error : 
```json
{
    "code": 400,
    "status": "BAD_REQUEST",
    "data": {
        "title": "title is required",
        "description": "description is required",
    }
}
```

## Update Todo API

Endpoint: PUT /api/todos/:todoId

Headers :
- Authorization : Bearer token

Request Body : 
```json
{
    "title": "Berenang di ujung pulau",
    "description": "nemu ikan di laut",
}
```

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "todo-11111",
        "title": "Berenang di ujung pulau",
        "description": "nemu ikan di laut",
        "categories": null,
        "created_at": "2024-04-01T13:58:23Z",
        "updated_at": "2024-04-01T14:02:08Z"
    }
}
```

Response Body Error : 
```json
{
    "code": 400,
    "status": "BAD_REQUEST",
    "data": {
        "description": "description is required",
        "title": "title is required"
    }
}
```

## Get Todo API

Endpoint: GET /api/todos/:todoId

Headers :
- Authorization : Bearer token

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "todo-daa764e9-985c-4fa6-bea0-6d82a1708926",
        "title": "Berenang di ujung pulau",
        "description": "nemu ikan di laut",
        "categories": null,
        "created_at": "2024-04-01T13:58:23Z",
        "updated_at": "2024-04-01T14:02:08Z"
    }
}
```

Response Body Error : 
```json
{
    "code": 404,
    "status": "NOT_FOUND",
    "data": "todo is not found"
}
```

## List Todo API

Endpoint: GET /api/todos

Headers :
- Authorization : Bearer Token

Query Params :
- search : Search by title or description, OPTIONAL.
- category : Search by category, OPTIONAL.

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "id": "todo-daa764e9-985c-4fa6-bea0-6d82a1708926",
            "title": "Berenang di ujung pulau",
            "description": "nemu ikan di laut",
            "categories": null,
            "created_at": "2024-04-01T13:58:23Z",
            "updated_at": "2024-04-01T14:02:08Z"
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

## Delete Todo API

Endpoint: DELETE /api/todos/:todoId

Headers :
- Authorization : Bearer token

Response Body Success : 
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": "todo-11111",
        "message": "todo has been deleted"
    }
}
```

Response Body Error: 
```json
{
    "code": 404,
    "status": "NOT_FOUND",
    "data": "todo is not found"
}
```





