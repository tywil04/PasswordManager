# /api/v1/vaults/notes
If a GET request requires parameters, the parameters can only be passed using query parameters. (GET requests does not support body).

All requests need to have a `Content-type` header that is either `application/json` or `application/xml`, requests wont work otherwise. 

Required/Optional Request/Response parameters are denoted like so:
- `required param` or `<required param>`
- `[optional param]` 

This API is public, however, it is strongly recommended that you use an official client.

## GET
### Description
// DESCRIPTION HERE //

### Request Format
#### Headers
```json
{
    "Authorization": "valid authToken"
}
```

#### Params
```json
{}
```

### Response Format
#### 200
```json
{
    "notes": [
        {
            "id": "uuid string",
            "vaultId": "uuid string",
            "name": "base64 string",
            "nameIv": "base64 string",
            "title": "base64 string",
            "titleIv": "base64 string",
            "content": "base64 string",
            "contentIv": "base64 string",
            "colour": "base64 string",
            "colourIv": "base64 string"
        }
    ]
}
```
#### 500 (Server Error)
```json
{
    "error": {
        "code": "string",
        "message": "string"
    }
}
```

### Example
```javascript
// JAVASCRIPT EXAMPLE HERE
```

## POST
### Description
// DESCRIPTION HERE //

### Request Format
#### Headers
```json
{
    "Authorization": "valid authToken"
}
```

#### Params
```json
{
    "vaultId": "uuid string",
    "name": "base64 string",
    "nameIv": "base64 string",
    "title": "base64 string",
    "titleIv": "base64 string",
    "content": "base64 string",
    "contentIv": "base64 string",
    "colour": "base64 string",
    "colourIv": "base64 string",
    "noteId": "uuid string"
}
```

### Response Format
#### 200
```json
{
    "noteId": "uuid string"
}
```
#### 400 (Client Error)
```json
{
    "error": {
        "code": "string",
        "message": "string"
    }
}
```
#### 500 (Server Error)
```json
{
    "error": {
        "code": "string",
        "message": "string"
    }
}
```

### Example
```javascript
// JAVASCRIPT EXAMPLE HERE
```

## PUT
### Description
// DESCRIPTION HERE //

### Request Format
#### Headers
```json
{
    "Authorization": "valid authToken"
}
```

#### Params
```json
{
    "vaultId": "uuid string",
    "name": "base64 string",
    "nameIv": "base64 string",
    "title": "base64 string",
    "titleIv": "base64 string",
    "content": "base64 string",
    "contentIv": "base64 string",
    "colour": "base64 string",
    "colourIv": "base64 string",
    "noteId": "uuid string"
}
```

### Response Format
#### 200
```json
{
    "noteId": "uuid string"
}
```
#### 400 (Client Error)
```json
{
    "error": {
        "code": "string",
        "message": "string"
    }
}
```
#### 500 (Server Error)
```json
{
    "error": {
        "code": "string",
        "message": "string"
    }
}
```

### Example
```javascript
// JAVASCRIPT EXAMPLE HERE
```

## DELETE
### Description
// DESCRIPTION HERE //

### Request Format
#### Headers
```json
{
    "Authorization": "valid authToken"
}
```

#### Params
```json
{
    "vaultId": "uuid string",
    "name": "base64 string",
    "nameIv": "base64 string",
    "title": "base64 string",
    "titleIv": "base64 string",
    "content": "base64 string",
    "contentIv": "base64 string",
    "colour": "base64 string",
    "colourIv": "base64 string",
    "noteId": "uuid string"
}
```

### Response Format
#### 200
```json
{}
```
#### 400 (Client Error)
```json
{
    "error": {
        "code": "string",
        "message": "string"
    }
}
```

### Example
```javascript
// JAVASCRIPT EXAMPLE HERE
```