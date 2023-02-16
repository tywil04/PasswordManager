package signup

import (
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/validations"
)

type PostInput struct {
	Email                  string `form:"email" json:"email" xml:"email"`
	MasterHash             string `form:"masterHash" json:"masterHash" xml:"masterHash"`
	ProtectedDatabaseKey   string `form:"protectedDatabaseKey" json:"protectedDatabaseKey" xml:"protectedDatabaseKey"`
	ProtectedDatabaseKeyIv string `form:"protectedDatabaseKeyIv" json:"protectedDatabaseKeyIv" xml:"protectedDatabaseKeyIv"`
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

	if input.ProtectedDatabaseKey == "" {
		c.JSON(400, helpers.ErrorMissing("protectedDatabaseKey"))
		return
	}

	if input.ProtectedDatabaseKeyIv == "" {
		c.JSON(400, helpers.ErrorMissing("protectedDatabaseKeyIv"))
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

	decodedProtectedDatabaseKey, dpdkErr := base64.StdEncoding.DecodeString(input.ProtectedDatabaseKey)
	if dpdkErr != nil {
		c.JSON(400, helpers.ErrorInvalid("protectedDatabaseKey"))
		return
	}

	decodedProtectedDatabaseKeyIv, dpdkiErr := base64.StdEncoding.DecodeString(input.ProtectedDatabaseKeyIv)
	if dpdkiErr != nil {
		c.JSON(400, helpers.ErrorInvalid("protectedDatabaseKeyIv"))
		return
	}

	testUser, _ := db.GetUserViaEmail(input.Email)
	if testUser != nil {
		c.JSON(400, helpers.ErrorInUse("email"))
		return
	}

	salt := cryptography.RandomBytes(16)
	strongMasterHash := cryptography.StrengthenMasterHash(decodedMasterHash, salt)

	user, userErr := db.Client.User.Create().
		SetEmail(input.Email).
		SetStrengthenedMasterHash(strongMasterHash).
		SetStrengthenedMasterHashSalt(salt).
		SetProtectedDatabaseKey(decodedProtectedDatabaseKey).
		SetProtectedDatabaseKeyIv(decodedProtectedDatabaseKeyIv).
		Save(db.Context)

	if userErr != nil {
		c.JSON(500, helpers.ErrorCreating("user"))
		return
	}

	challenge, challengeErr := helpers.GenerateChallenge(user)
	fmt.Println(challengeErr)
	if challengeErr != nil {
		c.JSON(500, helpers.ErrorIssuing("challenge"))
		return
	}

	availableChallenges := helpers.GetAvailableChallenges(user)

	c.JSON(200, gin.H{"challengeId": challenge.ID.String(), "availableChallenges": availableChallenges})
}
