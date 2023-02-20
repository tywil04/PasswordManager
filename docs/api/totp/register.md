# /api/v1/totp/register
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
    "totpCredentialId": "uuid string",
    "totpSecret": "base64 string",
    "totpSecretQr": "base64 string"
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
    "totpCredentialId": "string",
    "code": "string"
}
```

### Response Format
#### 200
```json
{
    "totpCredentialId": "uuid string"
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

### Example
```javascript
// JAVASCRIPT EXAMPLE HERE
```