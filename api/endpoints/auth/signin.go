package auth

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
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

	sameMasterHash := cryptography.CompareMasterHash(foundUser.StrengthenedMasterHash, params["masterHash"].([]byte), foundUser.StrengthenedMasterHashSalt)
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
