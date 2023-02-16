package signin

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/validations"
)

type PostInput struct {
	Email      string `form:"email" json:"email" xml:"email"`
	MasterHash string `form:"masterHash" json:"masterHash" xml:"masterHash"`
}

func Post(c *gin.Context) {
	var input PostInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.Email == "" {
		c.JSON(400, helpers.ErrorMissing("email"))
		return
	}

	if input.MasterHash == "" {
		c.JSON(400, helpers.ErrorMissing("masterHash"))
		return
	}

	if !validations.IsEmailValid(input.Email) {
		c.JSON(400, helpers.ErrorInvalid("email"))
		return
	}

	decodedMasterHash, dmhErr := base64.StdEncoding.DecodeString(input.MasterHash)
	if dmhErr != nil {
		c.JSON(400, helpers.ErrorInvalid("masterHash"))
		return
	}

	foundUser, _ := db.GetUserViaEmail(input.Email)
	if foundUser == nil {
		c.JSON(400, helpers.ErrorNotInUse("email"))
		return
	}

	strengthenedMasterHash := cryptography.StrengthenMasterHash(decodedMasterHash, foundUser.StrengthenedMasterHashSalt)
	sameMasterHash := cryptography.ConstantTimeCompare(strengthenedMasterHash, foundUser.StrengthenedMasterHash)
	if !sameMasterHash {
		c.JSON(400, helpers.ErrorInvalidCredentials())
		return
	}

	challenge, challengeErr := helpers.GenerateChallenge(foundUser)
	if challengeErr != nil {
		c.JSON(500, helpers.ErrorIssuing("challenge"))
		return
	}

	availableChallenges := helpers.GetAvailableChallenges(foundUser)

	c.JSON(200, gin.H{"challengeId": challenge.ID.String(), "availableChallenges": availableChallenges})
}
