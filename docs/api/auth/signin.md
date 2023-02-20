# /api/v1/auth/signin
If a GET request requires parameters, the parameters can only be passed using query parameters. (GET requests does not support body).

All requests need to have a `Content-type` header that is either `application/json` or `application/xml`, requests wont work otherwise. 

Required/Optional Request/Response parameters are denoted like so:
- `required param` or `<required param>`
- `[optional param]` 

This API is public, however, it is strongly recommended that you use an official client.

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
    "email": "string",
    "masterHash": "string"
}
```

### Response Format
#### 200
```json
{
    "challengeId": "uuid string",
    "availableChallenges": [
        "string"
    ]
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