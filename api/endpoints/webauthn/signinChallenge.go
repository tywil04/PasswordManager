package webauthn

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/smtp"
	internalWebauthn "PasswordManager/api/lib/webauthn"
)

type GetSigninChallengeInput struct {
	WebauthnChallengeId string `form:"webauthnChallengeId" json:"webauthnChallengeId" xml:"webauthnChallengeId"`
}

type PostSigninChallengeInput struct {
	WebauthnChallengeId string `form:"webauthnChallengeId" json:"webauthnChallengeId" xml:"webauthnChallengeId"`
	Credential          struct {
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

func GetSigninChallenge(c *gin.Context) {
	var input GetSigninChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.WebauthnChallengeId == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingWebauthnChallengeId", "message": "Required 'webauthnChallengeId' was not found."}})
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.WebauthnChallengeId)
	if dciErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingWebauthnChallengeId", "message": "Unable to parse 'webauthnChallengeId', expected uuid."}})
		return
	}

	foundWebauthnChallenge, fwcErr := db.GetWebauthnChallenge(decodedChallengeId)
	if fwcErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errWebauthnChallengeNotFound", "message": "Unable to find valid webauthn challenge using 'webauthnChallengeId'."}})
		return
	}

	foundUser, fuErr := db.GetWebauthnChallengeUser(foundWebauthnChallenge)
	if fuErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errWebauthnChallengeNotFound", "message": "Unable to find valid webauthn challenge using 'webauthnChallengeId'."}})
		return
	}

	options, sessionData, err := internalWebauthn.Web.BeginLogin(&internalWebauthn.User{User: foundUser})
	if err != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	challenge, challengeId := foundWebauthnChallenge.Update().
		SetChallenge(sessionData.Challenge).
		SetUserId(sessionData.UserID).
		SetAllowedCredentialIds(sessionData.AllowedCredentialIDs).
		SetUserVerification(string(sessionData.UserVerification)).
		SetExtensions(map[string]any(sessionData.Extensions)).
		SetUser(foundUser).
		Save(db.Context)

	if challengeId != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	c.JSON(200, gin.H{"webauthnChallengeId": challenge.ID.String(), "options": options})
}

func PostSigninChallenge(c *gin.Context) {
	var input PostSigninChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.WebauthnChallengeId == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingWebauthnChallengeId", "message": "Required 'webauthnChallengeId' was not found."}})
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.WebauthnChallengeId)
	if dciErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingWebauthnChallengeId", "message": "Unable to parse 'webauthnChallengeId', expected uuid."}})
		return
	}

	foundWebauthnChallenge, fwcErr := db.GetWebauthnChallenge(decodedChallengeId)
	if fwcErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errWebauthnChallengeNotFound", "message": "Unable to find valid webauthn challenge using 'webauthnChallengeId'."}})
		return
	}

	foundUser, foundUserErr := db.GetWebauthnChallengeUser(foundWebauthnChallenge)
	if foundUserErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errWebauthnChallengeNotFound", "message": "Unable to find valid webauthn challenge using 'webauthnChallengeId'."}})
		return
	}

	sessionData := webauthn.SessionData{
		Challenge:            foundWebauthnChallenge.Challenge,
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
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	_, credentialErr := internalWebauthn.Web.ValidateLogin(&internalWebauthn.User{User: foundUser}, sessionData, credentialData)
	if credentialErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errIncorrectWebauthn", "message": "Invalid webauthn attempt."}})
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})
		return
	}

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

	db.DeleteWebauthnChallenge(foundWebauthnChallenge)

	encodedProtectedDatabaseKey := base64.StdEncoding.EncodeToString(foundUser.ProtectedDatabaseKey)
	encodedProtectedDatabaseKeyIv := base64.StdEncoding.EncodeToString(foundUser.ProtectedDatabaseKeyIv)

	c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
}
