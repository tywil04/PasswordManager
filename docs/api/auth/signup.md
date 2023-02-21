# /api/v1/auth/signup
If a GET request requires parameters, the parameters can only be passed using query parameters. (GET requests does not support body).

All requests need to have a `Content-type` header that is either `application/json` or `application/xml`, requests wont work otherwise. 

Required/Optional Request/Response parameters are denoted like so:
- `required param` or `<required param>`
- `[optional param]` 

This API is public, however, it is strongly recommended that you use an official client.

## POST
### Description
This endpoints starts the signup process. The response contains a list of all the supported 2FA challenges the user has registered and an id for a challenge (by default the user will have only an email challenge available). `masterHash` should be generated using the flow shown in `signupFlow.png` but technically any base64 string is valid because the server cannot validate what it doesn't know. `protectedDatabaseKey` is an encrypted randomly generated key and `protectedDatabaseKeyIv` is the iv used for encryption, again this technically can be any base64 string because the server cannot validate what it doesn't know. 

- `email` must be a valid email address, if its not an error will be returned.
- `masterHash` must be a base64 encoded string, if its not an error will be returned.
- `protectedDatabaseKey` must be a base64 encoded string, if its not an error will be returned.
- `protectedDatabaseKeyIv` must be a base64 encoded string, if its not an error will be returned.

### Request Format
#### Headers
```json
{}
```

#### Params
```json
{
    "email": "email string",
    "masterHash": "base64 string",
    "protectedDatabaseKey": "base64 string",
    "protectedDatabaseKeyIv": "base64 string"
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
This example uses JavaScript and my WebCrypto wrapper (`ui/lib/cryptography.js`) alongside `base64-arraybuffer` from npm. The `masterHash` is generated exactly as show in the `signupFlow.png` diagram.
```javascript
const email = "example@example.org"
const password = "examplePassword"

let masterKey = await cryptography.generateMasterKey(password, email)
let masterHash = await cryptography.generateMasterHash(password, masterKey) 
let databaseKey = await cryptography.generateDatabaseKey()
let protectedDatabaseKey = await cryptography.protectDatabaseKey(masterKey, databaseKey)

const response = await fetch("/api/v1/auth/signup", {
    method: "POST",
    headers: {
        "Content-type": "application/json", // json or xml is allowed
    },
    body: JSON.stringify({
        email: email,
        masterHash: base64.encode(masterHash),
        protectedDatabaseKey:  base64.encode(protectedDatabaseKey.key),
        protectedDatabaseKeyIv: base64.encode(protectedDatabaseKey.iv),
    })
})
const json = await response.json()

console.log(json) // print response to console
```