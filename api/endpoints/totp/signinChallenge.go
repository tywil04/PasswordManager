package totp

import (
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/smtp"
)

type PostSigninTotpChallengeInput struct {
	TotpCredentialId string `form:"totpCredentialId" json:"totpCredentialId" xml:"totpCredentialId"`
	Code             string `form:"code" json:"code" xml:"code"`
}

func PostSigninTotpChallenge(c *gin.Context) {
	var input PostSigninTotpChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.TotpCredentialId == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingTotpCredentialId", "message": "Required 'totpCredentialId' was not found."}})
		return
	}

	if input.Code == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingCode", "message": "Required 'code' was not found."}})
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.TotpCredentialId)
	if dciErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingTotpCredentialId", "message": "Unable to parse 'totpCredentialId', expected uuid."}})
		return
	}

	foundCredential, foundCredentialErr := db.GetTotpCredentialViaId(decodedChallengeId)
	if foundCredentialErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errTotpCredentialNotFound", "message": "Unable to find valid challenge using 'totpCredentialId'."}})
		return
	}

	foundUser, foundUserErr := db.GetTotpCredentialUser(foundCredential)
	if foundUserErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	valid := totp.Validate(input.Code, foundCredential.Secret)

	if valid {
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Successful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: true})

		salt := cryptography.RandomBytes(32)
		encodedSalt := base64.StdEncoding.EncodeToString(salt)
		publicKey, signature := cryptography.GenerateSignature(foundUser.Email + base64.StdEncoding.EncodeToString(foundUser.StrengthenedMasterHash) + encodedSalt)

		session, sessionErr := db.Client.Session.Create().
			SetUser(foundUser).
			SetN(publicKey.N.Bytes()).
			SetE(publicKey.E).
			SetExpiry(time.Now().Add(time.Hour)).
			Save(db.Context)

		if sessionErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
			return
		}

		token := base64.StdEncoding.EncodeToString([]byte(session.ID.String())) + ";" + encodedSalt + ";" + base64.StdEncoding.EncodeToString(signature)

		encodedProtectedDatabaseKey := base64.StdEncoding.EncodeToString(foundUser.ProtectedDatabaseKey)
		encodedProtectedDatabaseKeyIv := base64.StdEncoding.EncodeToString(foundUser.ProtectedDatabaseKeyIv)

		c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
	} else if !valid {
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})

		c.JSON(400, gin.H{"error": gin.H{"code": "errIncorrectTotpChallengeCode", "message": "Incorrect code for totpChallenge."}})
	}
}
