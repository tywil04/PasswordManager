package email

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/smtp"
)

type GetChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId"`
}

type PostChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId"`
	Code        string `form:"code" json:"code" xml:"code"`
}

func GetChallenge(c *gin.Context) {
	var input GetChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, exceptions.Builder("body", exceptions.Invalid, exceptions.JsonOrXml))
		return
	}

	if input.ChallengeId == "" {
		c.JSON(400, exceptions.Builder("challengeId", exceptions.MissingParam))
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.ChallengeId)
	if dciErr != nil {
		c.JSON(400, exceptions.Builder("challengeId", exceptions.ParsingParam, exceptions.Uuid))
		return
	}

	challenge, challengeErr := db.GetUnexpiredChallengeViaId(decodedChallengeId)
	if challengeErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid, exceptions.Expired))
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

	randomCode := cryptography.RandomString(8)

	_, err := foundChallenge.Update().SetCode(randomCode).Save(db.Context)
	if err != nil {
		c.JSON(500, exceptions.Builder("challenge", exceptions.Updating, exceptions.TryAgain))
		return
	}

	go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Email Verification", smtp.ChallengeTemplate, smtp.ChallengeTemplateData{Code: randomCode})

	c.JSON(200, gin.H{})
}

func PostChallenge(c *gin.Context) {
	var input PostChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, exceptions.Builder("body", exceptions.Invalid, exceptions.JsonOrXml))
		return
	}

	if input.ChallengeId == "" {
		c.JSON(400, exceptions.Builder("challengeId", exceptions.MissingParam))
		return
	}

	if input.Code == "" {
		c.JSON(400, exceptions.Builder("code", exceptions.MissingParam))
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.ChallengeId)
	if dciErr != nil {
		c.JSON(400, exceptions.Builder("challengeId", exceptions.ParsingParam, exceptions.Uuid))
		return
	}

	challenge, challengeErr := db.GetUnexpiredChallengeViaId(decodedChallengeId)
	if challengeErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid, exceptions.Expired))
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

	sameCode := cryptography.ConstantTimeCompare([]byte(foundChallenge.Code), []byte(input.Code))
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
