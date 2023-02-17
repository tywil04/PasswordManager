# /api/v1/auth/signin
## Post
### Info
This endpoints starts the signin process. The `email` and `masterHash` are used to verify the user. The response contains a list of all the supported 2FA challenges the user has registered and an id for a challenge.

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
    "masterHash": base64 string
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


Status 403 (Signin Credentials Incorrect)
-----------------------------------------
{}
```

### Example
```javascript
// Todo
```