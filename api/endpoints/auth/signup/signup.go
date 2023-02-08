package signup

import (
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/smtp"
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/user"
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
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.Email == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingEmail", "message": "Required 'email' was not found."}})
		return
	}

	if input.MasterHash == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingMasterHash", "message": "Required 'masterHash' was not found."}})
		return
	}

	if input.ProtectedDatabaseKey == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingProtectedDatabaseKey", "message": "Required 'protectedDatabaseKey' was not found."}})
		return
	}

	if input.ProtectedDatabaseKeyIv == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingProtectedDatabaseKeyIv", "message": "Required 'protectedDatabaseKeyIv' was not found."}})
		return
	}

	decodedMasterHash, dmhErr := base64.StdEncoding.DecodeString(input.MasterHash)
	if dmhErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingMasterHash", "message": "Unable to parse 'masterHash', expected base64 encoding."}})
		return
	}

	decodedProtectedDatabaseKey, dpdkErr := base64.StdEncoding.DecodeString(input.ProtectedDatabaseKey)
	if dpdkErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingProtectedDatabaseKey", "message": "Unable to parse 'protectedDatabaseKey', expected base64 encoding."}})
		return
	}

	decodedProtectedDatabaseKeyIv, dpdkiErr := base64.StdEncoding.DecodeString(input.ProtectedDatabaseKeyIv)
	if dpdkiErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingProtectedDatabaseKeyIv", "message": "Unable to parse 'protectedDatabaseKeyIv', expected base64 encoding."}})
		return
	}

	testUser, _ := db.Client.User.Query().Where(user.EmailEQ(input.Email)).First(db.Context)
	if testUser != nil {
		if !testUser.Verified {
			db.Client.User.DeleteOne(testUser).Exec(db.Context)
		} else if testUser.Verified {
			c.JSON(400, gin.H{"error": gin.H{"code": "errEmailInUse", "message": "Email is in use."}})
			return
		}
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
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	randomCode := cryptography.RandomString(8)
	challenge, challengeErr := db.Client.EmailChallenge.Create().
		SetCode(randomCode).
		SetUser(user).
		SetExpiry(time.Now().Add(time.Hour)).
		SetFor(emailchallenge.ForSignup).
		Save(db.Context)

	if challengeErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	go smtp.SendTemplate(input.Email, "PasswordManager5 Email Verification", smtp.ChallengeTemplate, smtp.ChallengeTemplateData{
		Code: randomCode,
	})

	c.JSON(200, gin.H{"emailChallengeId": challenge.ID.String()})
}
