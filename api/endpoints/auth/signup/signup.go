package signup

import (
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
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
		c.JSON(400, exceptions.Builder("body", exceptions.Invalid, exceptions.JsonOrXml))
		return
	}

	if input.Email == "" {
		c.JSON(400, exceptions.Builder("email", exceptions.MissingParam))
		return
	}

	if input.MasterHash == "" {
		c.JSON(400, exceptions.Builder("masterHash", exceptions.MissingParam))
		return
	}

	if input.ProtectedDatabaseKey == "" {
		c.JSON(400, exceptions.Builder("protectedDatabaseKey", exceptions.MissingParam))
		return
	}

	if input.ProtectedDatabaseKeyIv == "" {
		c.JSON(400, exceptions.Builder("protectedDatabaseKeyIv", exceptions.MissingParam))
		return
	}

	if !validations.IsEmailValid(input.Email) {
		c.JSON(400, exceptions.Builder("email", exceptions.ParsingParam, exceptions.Email))
		return
	}

	decodedMasterHash, dmhErr := base64.StdEncoding.DecodeString(input.MasterHash)
	if dmhErr != nil {
		c.JSON(400, exceptions.Builder("masterHash", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	decodedProtectedDatabaseKey, dpdkErr := base64.StdEncoding.DecodeString(input.ProtectedDatabaseKey)
	if dpdkErr != nil {
		c.JSON(400, exceptions.Builder("protectedDatabaseKey", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	decodedProtectedDatabaseKeyIv, dpdkiErr := base64.StdEncoding.DecodeString(input.ProtectedDatabaseKeyIv)
	if dpdkiErr != nil {
		c.JSON(400, exceptions.Builder("protectedDatabaseKeyIv", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	testUser, _ := db.GetUserViaEmail(input.Email)
	if testUser != nil {
		c.JSON(400, exceptions.Builder("email", exceptions.InvalidParam, exceptions.InUse))
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
