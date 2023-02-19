# /api/v1/webauthn/register
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
    "webauthnRegisterChallengeId": uuid string,
    "options": {
        "publicKey": {
            "challenge": url base64 string,
            "rp": {
                "name": string,
                "[icon]": string,
                "id": string
            },
            "user": {
                "name": string,
                "[icon]": string,
                "[displayName]": string,
                "id": url base64 string
            },
            "[pubKeyCredParams]": [
                {
                    "type": string,
                    "alg": int
                },
                ...
            ],
            "[authenticatorSelection]": {
                "[authenticatorAttachment]": string,
                "[requireResidentKey]": bool,
                "[residentKey]": string,
                "[userVerification]": string
            },
            "[timeout]": int,
            "[excludeCredentials]": [
                {
                    "type": string,
                    "id": url base64 string,
                    "[transports]": [
                        string,
                        ...
                    ]
                },
                ...
            ],
            "[attestation]": string
        }
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
    "webauthnRegisterChallengeId": string,
    "name": string,
    "credential": {
        "authenticatorAttachment": string,
        "id": string,
        "rawId": string,
        "response": {
            "attestationObject": string,
            "clientDataJSON": string
        },
        "type": string
    }
}
```

### Response Format
```
Status 200:
-----------
{
    "webauthnCredentialId": uuid string
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