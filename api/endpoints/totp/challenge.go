package totp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/smtp"
)

type PostChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId"`
	Code        string `form:"code" json:"code" xml:"code"`
}

func PostChallenge(c *gin.Context) {
	var input PostChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.ChallengeId == "" {
		c.JSON(400, helpers.ErrorMissing("challengeId"))
		return
	}

	if input.Code == "" {
		c.JSON(400, helpers.ErrorMissing("code"))
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.ChallengeId)
	if dciErr != nil {
		c.JSON(400, helpers.ErrorInvalid("challengeId"))
		return
	}

	challenge, challengeErr := db.GetUnexpiredChallengeViaId(decodedChallengeId)
	if challengeErr != nil {
		c.JSON(400, helpers.ErrorExpired("challenge"))
		return
	}

	foundCredential, foundCredentialErr := db.GetChallengeTotpCredential(challenge)
	if foundCredentialErr != nil {
		c.JSON(400, helpers.ErrorInvalid("challenge"))
		return
	}

	foundUser, foundUserErr := db.GetChallengeUser(challenge)
	if foundUserErr != nil {
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	valid := totp.Validate(input.Code, foundCredential.Secret)

	if valid {
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Successful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: true})

		db.DeleteChallenge(challenge)

		token, encodedProtectedDatabaseKey, encodedProtectedDatabaseKeyIv, err := helpers.GenerateSession(foundUser)
		if err != nil {
			c.JSON(500, helpers.ErrorIssuing("session"))
			return
		}

		c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
	} else if !valid {
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})
		c.JSON(403, gin.H{})
	}
}
