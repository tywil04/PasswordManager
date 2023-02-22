package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
)

const (
	SignupDescription string = `This endpoints starts the signup process. The response contains a list of all the supported 2FA challenges the user has registered and an id for a challenge (by default the user will have only an email challenge available, other challenges include totp and webauthn). 

- 'email' should be a valid email address
- 'masterHash' should be generated using a flow similar to 'docs/signupFlow.png' but technically any base64 string is valid because the server cannot validate what it doesn't know.
- 'protectedDatabaseKey' is an encrypted randomly generated key that is encrypted and 'protectedDatabaseKeyIv' is the iv used for encryption. These should be generated and encrypted using a flow similar to 'docs/signupFlow.png'. Both of these can technically be any base64 string because the server cannot validate encrypted data.`
)

type PostSignupInput struct {
	Email                  string `form:"email" json:"email" xml:"email" pmParseType:"email"`
	MasterHash             string `form:"masterHash" json:"masterHash" xml:"masterHash" pmParseType:"base64"`
	ProtectedDatabaseKey   string `form:"protectedDatabaseKey" json:"protectedDatabaseKey" xml:"protectedDatabaseKey" pmParseType:"base64"`
	ProtectedDatabaseKeyIv string `form:"protectedDatabaseKeyIv" json:"protectedDatabaseKeyIv" xml:"protectedDatabaseKeyIv" pmParseType:"base64"`
}

func PostSignup(c *gin.Context) {
	params := c.GetStringMap("params")

	email := params["email"].(string)
	testUser, _ := db.GetUserViaEmail(email)
	if testUser != nil {
		c.JSON(400, exceptions.Builder("email", exceptions.InvalidParam, exceptions.InUse))
		return
	}

	salt := cryptography.RandomBytes(16)
	strongMasterHash := cryptography.StrengthenMasterHash(params["masterHash"].([]byte), salt)

	user, userErr := db.Client.User.Create().
		SetEmail(email).
		SetStrengthenedMasterHash(strongMasterHash).
		SetStrengthenedMasterHashSalt(salt).
		SetProtectedDatabaseKey(params["protectedDatabaseKey"].([]byte)).
		SetProtectedDatabaseKeyIv(params["protectedDatabaseKeyIv"].([]byte)).
		Save(db.Context)

	if userErr != nil {
		c.JSON(500, exceptions.Builder("user", exceptions.Creating, exceptions.TryAgain))
		return
	}

	challenge, challengeErr := helpers.GenerateChallenge(user)
	fmt.Println(challengeErr)
	if challengeErr != nil {
		c.JSON(500, exceptions.Builder("challenge", exceptions.Issuing, exceptions.TryAgain))
		return
	}

	availableChallenges := helpers.GetAvailableChallenges(user)

	c.JSON(200, gin.H{"challengeId": challenge.ID.String(), "availableChallenges": availableChallenges})
}
