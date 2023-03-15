# /api/v1/2fa/email/challenge
If a GET request requires parameters, the parameters can only be passed using query parameters. (GET requests do not support having a body).

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
{}
```

#### Params
```json
{
    "challengeId": "uuid string",
    "code": "string"
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
{}
```

#### Params
```json
{
    "challengeId": "uuid string",
    "code": "string"
}
```

### Response Format
#### 200
```json
{
    "authToken": "authToken string",
    "protectedDatabaseKey": "base64 string",
    "protectedDatabaseKeyIv": "base64 string"
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