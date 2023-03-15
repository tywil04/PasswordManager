package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
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
