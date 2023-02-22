package totp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/smtp"
)

const (
	ChallengeDescription string = ""
)

type PostChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId" pmParseType:"uuid"`
	Code        string `form:"code" json:"code" xml:"code"`
}

func PostChallenge(c *gin.Context) {
	params := c.GetStringMap("params")

	challenge, challengeErr := db.GetUnexpiredChallengeViaId(params["challengeId"].(uuid.UUID))
	if challengeErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid, exceptions.Expired))
		return
	}

	foundCredential, foundCredentialErr := db.GetChallengeTotpCredential(challenge)
	if foundCredentialErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	foundUser, foundUserErr := db.GetChallengeUser(challenge)
	if foundUserErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	valid := totp.Validate(params["code"].(string), foundCredential.Secret)

	if valid {
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Successful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: true})

		db.DeleteChallenge(challenge)

		token, encodedProtectedDatabaseKey, encodedProtectedDatabaseKeyIv, err := helpers.GenerateSession(foundUser)
		if err != nil {
			c.JSON(500, exceptions.Builder("session", exceptions.Issuing, exceptions.TryAgain))
			return
		}

		c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
	} else if !valid {
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})
		c.JSON(400, exceptions.Builder("code", exceptions.IncorrectChallenge))
	}
}
