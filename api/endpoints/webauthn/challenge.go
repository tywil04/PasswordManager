package webauthn

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/smtp"
	internalWebauthn "PasswordManager/api/lib/webauthn"
)

type GetChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId"`
}

type PostChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId"`
	Credential  struct {
		AuthenticatorAttachment string `form:"authenticatorAttachment" json:"authenticatorAttachment" xml:"authenticatorAttachment"`
		Id                      string `form:"id" json:"id" xml:"id"`
		RawId                   string `form:"rawId" json:"rawId" xml:"rawId"`
		Response                struct {
			AuthenticatorData string `form:"authenticatorData" json:"authenticatorData" xml:"authenticatorData"`
			ClientDataJSON    string `form:"clientDataJSON" json:"clientDataJSON" xml:"clientDataJSON"`
			Signature         string `form:"signature" json:"signature" xml:"signature"`
		} `form:"response" json:"response" xml:"response"`
		Type string `form:"type" json:"type" xml:"type"`
	} `form:"credential" json:"credential" xml:"credential"`
}

func GetChallenge(c *gin.Context) {
	var input GetChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.ChallengeId == "" {
		c.JSON(400, helpers.ErrorMissing("challengeId"))
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

	foundWebauthnChallenge, fwcErr := db.GetChallengeWebauthnChallenge(challenge)
	if fwcErr != nil {
		c.JSON(400, helpers.ErrorInvalid("challenge"))
		return
	}

	foundUser, fuErr := db.GetChallengeUser(challenge)
	if fuErr != nil {
		c.JSON(400, helpers.ErrorUnknown())
		return
	}

	options, sessionData, err := internalWebauthn.Web.BeginLogin(&internalWebauthn.User{User: foundUser})
	if err != nil {
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	_, updateErr := foundWebauthnChallenge.Update().
		SetSdChallenge(sessionData.Challenge).
		SetUserId(sessionData.UserID).
		SetAllowedCredentialIds(sessionData.AllowedCredentialIDs).
		SetUserVerification(string(sessionData.UserVerification)).
		SetExtensions(map[string]any(sessionData.Extensions)).
		Save(db.Context)

	if updateErr != nil {
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	c.JSON(200, gin.H{"options": options})
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

	foundWebauthnChallenge, fwcErr := db.GetChallengeWebauthnChallenge(challenge)
	if fwcErr != nil {
		c.JSON(400, helpers.ErrorInvalid("challenge"))
		return
	}

	foundUser, foundUserErr := db.GetChallengeUser(challenge)
	if foundUserErr != nil {
		c.JSON(400, helpers.ErrorInvalid("challenge"))
		return
	}

	sessionData := webauthn.SessionData{
		Challenge:            foundWebauthnChallenge.SdChallenge,
		UserID:               foundWebauthnChallenge.UserId,
		AllowedCredentialIDs: foundWebauthnChallenge.AllowedCredentialIds,
		UserVerification:     protocol.UserVerificationRequirement(foundWebauthnChallenge.UserVerification),
		Extensions:           protocol.AuthenticationExtensions(foundWebauthnChallenge.Extensions),
	}

	data, _ := json.Marshal(input.Credential)
	dataReader := bytes.NewReader(data)

	credentialData, cdErr := protocol.ParseCredentialRequestResponseBody(dataReader)
	fmt.Println(cdErr)
	if cdErr != nil {
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	_, credentialErr := internalWebauthn.Web.ValidateLogin(&internalWebauthn.User{User: foundUser}, sessionData, credentialData)
	if credentialErr != nil {
		c.JSON(400, helpers.ErrorChallenge("webauthn"))
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})
		return
	}

	go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Successful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: true})

	db.DeleteChallenge(challenge)

	token, encodedProtectedDatabaseKey, encodedProtectedDatabaseKeyIv, err := helpers.GenerateSession(foundUser)
	if err != nil {
		c.JSON(500, helpers.ErrorIssuing("session"))
		return
	}

	c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
}
