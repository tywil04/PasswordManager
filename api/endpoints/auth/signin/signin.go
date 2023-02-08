package signin

import (
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/smtp"
	"PasswordManager/api/lib/validations"
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/user"
)

type PostInput struct {
	Email      string `form:"email" json:"email" xml:"email"`
	MasterHash string `form:"masterHash" json:"masterHash" xml:"masterHash"`
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

	if !validations.IsEmailValid(input.Email) {
		c.JSON(400, gin.H{"error": gin.H{"code": "errInvalidEmail", "message": "Required 'email' is malformed."}})
		return
	}

	decodedMasterHash, dmhErr := base64.StdEncoding.DecodeString(input.MasterHash)
	if dmhErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingMasterHash", "message": "Unable to parse 'masterHash', expected base64 encoding."}})
		return
	}

	foundUser, _ := db.Client.User.Query().Where(user.EmailEQ(input.Email)).First(db.Context)
	if foundUser == nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errEmailNotInUse", "message": "Email is not in use."}})
		return
	}

	if !foundUser.Verified {
		c.JSON(400, gin.H{"error": gin.H{"code": "errUserNotVerified", "message": "User is not verified."}})
		return
	}

	strengthenedMasterHash := cryptography.StrengthenMasterHash(decodedMasterHash, foundUser.StrengthenedMasterHashSalt)
	sameMasterHash := cryptography.ConstantTimeCompare(strengthenedMasterHash, foundUser.StrengthenedMasterHash)
	if !sameMasterHash {
		c.JSON(400, gin.H{"error": gin.H{"code": "errInvalidCredentials", "message": "Invalid credentials."}})
		return
	}

	if foundUser.Default2FA == user.Default2FAEmail {
		randomCode := cryptography.RandomString(8)
		challenge, challengeErr := db.Client.EmailChallenge.Create().
			SetCode(randomCode).
			SetUser(foundUser).
			SetExpiry(time.Now().Add(time.Hour)).
			SetFor(emailchallenge.ForSignin).
			Save(db.Context)

		if challengeErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
			return
		}

		go smtp.SendTemplate(input.Email, "PasswordManager5 Email Verification", smtp.ChallengeTemplate, smtp.ChallengeTemplateData{
			Code: randomCode,
		})

		c.JSON(200, gin.H{"challengeType": "emailChallenge", "emailChallengeId": challenge.ID.String()})
	} else if foundUser.Default2FA == user.Default2FAWebauthn {
		challenge, challengeErr := db.Client.WebAuthnChallenge.Create().
			SetUser(foundUser).
			Save(db.Context)

		if challengeErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
			return
		}

		c.JSON(200, gin.H{"challengeType": "webauthnChallenge", "webauthnChallengeId": challenge.ID.String()})
	}
}
