# /api/v1/webauthn/credential
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
```
Headers:
--------
{
    "Authorization": valid authToken
}


Params:
-------
{}
```

### Response Format
```
Status 200:
-----------
{
    "webauthnCredentials": [
        {
            "id": uuid string,
            "name": string,
            "createdAt": "time"
        },
        ...
    ]
}


Status 500 (Server Error):
--------------------------
{
    "error": {
        "code": string,
        "message": string
    }
}
```

### Example
// EXAMPLE HERE //

## DELETE
### Description
// DESCRIPTION HERE //

### Request Format
```
Headers:
--------
{
    "Authorization": valid authToken
}


Params:
-------
{
    "webauthnCredentialId": string
}
```

### Response Format
```
Status 200:
-----------
{}


Status 400 (Client Error):
--------------------------
{
    "error": {
        "code": string,
        "message": string
    }
}


Status 500 (Server Error):
--------------------------
{
    "error": {
        "code": string,
        "message": string
    }
}
```

### Example
// EXAMPLE HERE //