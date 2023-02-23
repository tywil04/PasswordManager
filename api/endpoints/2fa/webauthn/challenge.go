package webauthn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/api/lib/smtp"
	internalWebauthn "PasswordManager/api/lib/webauthn"
)

const (
	ChallengeDescription string = "This is the endpoint for completing a challenge via a webauthn device. A user must explicitly register a webauthn device for it to be used as a 2FA device. Both a challengeId and list of available challenges will be issued to the client by some other request (e.g /api/v1/auth/signup). To start the 2FA process webauthn, you need to contact the server via a GET request containing a challengeId. The response will be a valid webauthn credential, this should be used on the client to start the verification process. Once the user has verified there device, the result should be sent to the server alongside the challengeId using a POST request. If the 2FA credential from the device matches a previously registered credential then the user is authenticated. Which means a valid authToken (they expire after an hour) is generated and returned (this allows you to make requests for authenticated endpoints). The users protectedDatabaseKey and protectedDatabaseKeyIv are also returned so they can be decoded and stored within sessionStorage."
)

type GetChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId" pmParseType:"uuid"`
}

type PostChallengeInput struct {
	ChallengeId string `form:"challengeId" json:"challengeId" xml:"challengeId" pmParseType:"uuid"`
	Credential  struct {
		AuthenticatorAttachment string `form:"authenticatorAttachment" json:"authenticatorAttachment" xml:"authenticatorAttachment" pmOptional:"true"`
		Id                      string `form:"id" json:"id" xml:"id" pmOptional:"true"`
		RawId                   string `form:"rawId" json:"rawId" xml:"rawId" pmOptional:"true"`
		Response                struct {
			AuthenticatorData string `form:"authenticatorData" json:"authenticatorData" xml:"authenticatorData" pmOptional:"true"`
			ClientDataJSON    string `form:"clientDataJSON" json:"clientDataJSON" xml:"clientDataJSON" pmOptional:"true"`
			Signature         string `form:"signature" json:"signature" xml:"signature" pmOptional:"true"`
		} `form:"response" json:"response" xml:"response" pmOptional:"true"`
		Type string `form:"type" json:"type" xml:"type" pmOptional:"true"`
	} `form:"credential" json:"credential" xml:"credential"`
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

	foundWebauthnChallenge, fwcErr := db.GetChallengeWebauthnChallenge(challenge)
	if fwcErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	foundUser, fuErr := db.GetChallengeUser(challenge)
	if fuErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	options, sessionData, err := internalWebauthn.Web.BeginLogin(&internalWebauthn.User{User: foundUser})
	if err != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
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
		c.JSON(500, exceptions.Builder("challenge", exceptions.Updating, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"options": options})
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

	foundWebauthnChallenge, fwcErr := db.GetChallengeWebauthnChallenge(challenge)
	if fwcErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	foundUser, foundUserErr := db.GetChallengeUser(challenge)
	if foundUserErr != nil {
		c.JSON(400, exceptions.Builder("challenge", exceptions.Invalid))
		return
	}

	sessionData := webauthn.SessionData{
		Challenge:            foundWebauthnChallenge.SdChallenge,
		UserID:               foundWebauthnChallenge.UserId,
		AllowedCredentialIDs: foundWebauthnChallenge.AllowedCredentialIds,
		UserVerification:     protocol.UserVerificationRequirement(foundWebauthnChallenge.UserVerification),
		Extensions:           protocol.AuthenticationExtensions(foundWebauthnChallenge.Extensions),
	}

	data, _ := json.Marshal(params["credential"].(map[string]any))
	dataReader := bytes.NewReader(data)

	credentialData, cdErr := protocol.ParseCredentialRequestResponseBody(dataReader)
	fmt.Println(cdErr)
	if cdErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	_, credentialErr := internalWebauthn.Web.ValidateLogin(&internalWebauthn.User{User: foundUser}, sessionData, credentialData)
	if credentialErr != nil {
		c.JSON(400, exceptions.Builder("webauthn", exceptions.IncorrectChallenge))
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})
		return
	}

	go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Successful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: true})

	db.DeleteChallenge(challenge)

	token, encodedProtectedDatabaseKey, encodedProtectedDatabaseKeyIv, err := helpers.GenerateSession(foundUser)
	if err != nil {
		c.JSON(500, exceptions.Builder("session", exceptions.Issuing, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
}
