# /api/v1/password
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
    "passwords": [
        {
            "id": uuid string,
            "name": base64 string,
            "nameIv": base64 string,
            "username": base64 string,
            "usernameIv": base64 string,
            "password": base64 string,
            "passwordIv": base64 string,
            "colour": "hex colour string",
            "additionalFields": [
                {
                    "key": base64 string,
                    "keyIv": base64 string,
                    "value": base64 string,
                    "valueIv": base64 string
                }
            ],
            "urls": [
                {
                    "url": base64 string,
                    "urlIv": base64 string
                }
            ]
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

## POST
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
    "name": string,
    "nameIv": string,
    "username": string,
    "usernameIv": string,
    "password": string,
    "passwordIv": string,
    "colour": string,
    "additionalFields": [
        {
            "key": string,
            "keyIv": string,
            "value": string,
            "valueIv": string
        },
        ...
    ],
    "urls": [
        {
            "url": string,
            "urlIv": string
        },
        ...
    ],
    "passwordId": string
}
```

### Response Format
```
Status 200:
-----------
{
    "passwordId": uuid string
}


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
    "name": string,
    "nameIv": string,
    "username": string,
    "usernameIv": string,
    "password": string,
    "passwordIv": string,
    "colour": string,
    "additionalFields": [
        {
            "key": string,
            "keyIv": string,
            "value": string,
            "valueIv": string
        },
        ...
    ],
    "urls": [
        {
            "url": string,
            "urlIv": string
        },
        ...
    ],
    "passwordId": string
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
```

### Example
// EXAMPLE HERE //