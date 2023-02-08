# API
`Content-type` must be set to `application/json` or `application/xml` for the API to work.


# /api/v1/auth/signup
## Post
### Info
This is the API endpoint that creates users. The `email` must be valid, the `masterHash` is used to verify the user after creation (it is strengthend server side before stored). The `protectedDatabaseKey` and `protectedDatabaseKeyIv` is a randomly generated key that gets encrypted via the `masterKey`. The response is an id for a challenge the user must complete for the user to be verified. Only verified users can signin. If an account is not verified, the next time somebody creates an account with that email address the account pending verification will be deleted. This API is public but its recommended to use the official client.

To verify use the endpoint `/api/v1/email/signupChallenge`.

### Request Format
```
Headers:
--------
None


Query Params:
-------------
None


Body Params:
------------
{
    "email": string,
    "masterHash": base64 string,
    "protectedDatabaseKey": base64 string,
    "protectedDatabaseKeyIv": base64 string,
}
```

### Response Format
```
Status 200
----------
{
	"emailChallengeId": uuid string
}


Status 400
----------
{
    "error": {
        "code": string,
        "message": string
    }
}
```

### Example
This example uses my WebCrypto wrapper (`ui/lib/cryptography`) alongside `base64-arraybuffer` from npm
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
        "Content-type": "application/json",
    },
    body: JSON.stringify({
        email: email,
        masterHash: base64.encode(masterHash),
        protectedDatabaseKey:  base64.encode(protectedDatabaseKey.key),
        protectedDatabaseKeyIv: base64.encode(protectedDatabaseKey.iv),
    })
})
const json = await response.json()

console.log(json)
```

# /api/v1/auth/signin
## Post
### Info
This is the API endpoint that authenticates users. The `email` must be valid, the `masterHash` is used to verify the user. The response is an id for a challenge the user must complete for the users `protectedDatabaseKey` and a valid `authToken` to be generated and returned. Only verified users can signin. This API is public but its recommended to use the official client.

Depending on the challenge returned, to verify use the endpoints `/api/v1/webauthn/signinChallenge || /api/v1/email/signinChallenge`. 

### Request Format
```
Headers:
--------
None


Query Params:
-------------
None


Body Params:
------------
{
    "email": string,
    "masterHash": base64 string,
}
```

### Response Format
```
Status 200
----------
{
    "challengeType": string,
    "webauthnChallengeId" uuid string || "emailChallengeId" uuid string
}


Status 400
----------
{
    "error": {
        "code": string,
        "message": string
    }
}
```

### Example
```javascript
const email = "example@example.org"
const password = "examplePassword"

let masterKey = await cryptography.generateMasterKey(password, email)
let masterHash = await cryptography.generateMasterHash(password, masterKey) 

const response = await fetch("/api/v1/auth/signin", {
    method: "POST",
    headers: {
        "Content-type": "application/json",
    },
    body: JSON.stringify({
        email: email,
        masterHash: base64.encode(masterHash),
    })
})
const json = await response.json()

console.log(json)
```