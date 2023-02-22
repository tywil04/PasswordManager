package auth

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
)

const (
	SigninDescription string = `This endpoints starts the signin process. The response contains a list of all the supported 2FA challenges the user has registered and an id for a challenge (by default the user will have only an email challenge available, other challenges include totp and webauthn). 

- 'email' should be a valid email address
- 'masterHash' should be generated using a flow similar to 'docs/signinFlow.png' but technically any base64 string is valid because the server cannot validate what it doesn't know.`
)

type PostSigninInput struct {
	Email      string `form:"email" json:"email" xml:"email" pmParseType:"email"`
	MasterHash string `form:"masterHash" json:"masterHash" xml:"masterHash" pmParseType:"base64"`
}

func PostSignin(c *gin.Context) {
	params := c.GetStringMap("params")

	foundUser, _ := db.GetUserViaEmail(params["email"].(string))
	if foundUser == nil {
		c.JSON(400, exceptions.Builder("email", exceptions.InvalidParam, exceptions.NotInUse))
		return
	}

	strengthenedMasterHash := cryptography.StrengthenMasterHash(params["masterHash"].([]byte), foundUser.StrengthenedMasterHashSalt)
	sameMasterHash := cryptography.ConstantTimeCompare(strengthenedMasterHash, foundUser.StrengthenedMasterHash)
	if !sameMasterHash {
		c.JSON(400, exceptions.Builder("", exceptions.IncorrectCredentials))
		return
	}

	challenge, challengeErr := helpers.GenerateChallenge(foundUser)
	if challengeErr != nil {
		c.JSON(500, exceptions.Builder("challenge", exceptions.Issuing, exceptions.TryAgain))
		return
	}

	availableChallenges := helpers.GetAvailableChallenges(foundUser)

	c.JSON(200, gin.H{"challengeId": challenge.ID.String(), "availableChallenges": availableChallenges})
}
