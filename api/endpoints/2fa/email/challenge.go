package email

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/smtp"
)

type GetChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId" pmParseType:"uuid"`
}

type PostChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId" pmParseType:"uuid"`
	Code        string `form:"code" json:"code" xml:"code" pmParseType:"emailCode"`
}

func GetChallenge(c *gin.Context) {
	params := c.GetStringMap("params")

	challenge, challengeErr := db.GetChallenge(params["challengeId"].(uuid.UUID))
	if challengeErr != nil {
		c.JSON(400, exceptions.Builder("challengeId", exceptions.InvalidParam, exceptions.Uuid))
		return
	}

	if challenge.Expiry.Before(time.Now()) {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid, exceptions.Expired))
		db.DeleteChallenge(challenge)
		return
	}

	foundChallenge, foundChallengeErr := db.GetChallengeEmailChallenge(challenge)
	if foundChallengeErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	foundUser, foundUserErr := db.GetChallengeUser(challenge)
	if foundUserErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	randomCode := cryptography.RandomString(4) + "-" + cryptography.RandomString(4)

	_, err := foundChallenge.Update().SetCode(cryptography.HashString(randomCode)).Save(db.Context)
	if err != nil {
		c.JSON(500, exceptions.Builder("challenge", exceptions.Updating, exceptions.TryAgain))
		return
	}

	go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Email Verification", smtp.ChallengeTemplate, smtp.ChallengeTemplateData{Code: randomCode})

	c.JSON(200, gin.H{})
}

func PostChallenge(c *gin.Context) {
	params := c.GetStringMap("params")

	challenge, challengeErr := db.GetChallenge(params["challengeId"].(uuid.UUID))
	if challengeErr != nil {
		c.JSON(400, exceptions.Builder("challengeId", exceptions.InvalidParam, exceptions.Uuid))
		return
	}

	if challenge.Expiry.Before(time.Now()) {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid, exceptions.Expired))
		db.DeleteChallenge(challenge)
		return
	}

	foundChallenge, foundChallengeErr := db.GetChallengeEmailChallenge(challenge)
	if foundChallengeErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	if foundChallenge.Code == nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid, exceptions.DidntStart2FA))
		return
	}

	foundUser, foundUserErr := db.GetChallengeUser(challenge)
	if foundUserErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	sameCode := cryptography.ConstantTimeCompare(foundChallenge.Code, cryptography.HashString(params["code"].(string)))
	if !sameCode {
		c.JSON(400, exceptions.Builder("code", exceptions.IncorrectChallenge))
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})
		return
	}

	go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Successful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: true})

	if !foundUser.Verified {
		foundUser.Update().SetVerified(true).Exec(db.Context)
	}

	db.DeleteChallenge(challenge)

	token, encodedProtectedDatabaseKey, encodedProtectedDatabaseKeyIv, err := helpers.GenerateSession(foundUser)
	if err != nil {
		c.JSON(500, exceptions.Builder("session", exceptions.Issuing, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
}
