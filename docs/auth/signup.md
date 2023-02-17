# /api/v1/auth/signup
## Post
### Info
This endpoints starts the signup process. The `email` and `masterHash` must be valid. The masterHash is used to verify the user later. The response contains a list of all the supported 2FA challenges the user has registered and an id for a challenge (by default the user will have only an email challenge available).

Challenge endpoints:
- Email: `/api/v1/email/challenge`
- Totp: `/api/v1/totp/challenge`
- WebAuthn: `/api/v1/webauthn/challenge`

While this API is public it is strongly recommended that you use the official client.

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
Status 200 (Success)
--------------------
{
    "availableChallenges": [
        ...string
    ],
    "challengeId": uuid string
}


Status 400 (Erroneous)
----------------------
{
    "error": {
        "code": string,
        "causee": string
    }
}
```

### Example
This example uses JavaScript and my WebCrypto wrapper (`ui/lib/cryptography`) alongside `base64-arraybuffer` from npm
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