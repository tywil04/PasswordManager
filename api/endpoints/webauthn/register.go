package webauthn

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	internalWebauthn "PasswordManager/api/lib/webauthn"
	"PasswordManager/ent"
	"PasswordManager/ent/user"
)

type GetRegisterInput struct{}

type PostRegisterInput struct {
	WebauthnChallengeId string `form:"webauthnChallengeId" json:"webauthnChallengeId" xml:"webauthnChallengeId"`
	Name                string `form:"name" json:"name" xml:"json"`
	Credential          struct {
		AuthenticatorAttachment string `form:"authenticatorAttachment" json:"authenticatorAttachment" xml:"authenticatorAttachment"`
		Id                      string `form:"id" json:"id" xml:"id"`
		RawId                   string `form:"rawId" json:"rawId" xml:"rawId"`
		Response                struct {
			AttestationObject string `form:"attestationObject" json:"attestationObject" xml:"attestationObject"`
			ClientDataJSON    string `form:"clientDataJSON" json:"clientDataJSON" xml:"clientDataJSON"`
		} `form:"response" json:"response" xml:"response"`
		Type string `form:"type" json:"type" xml:"type"`
	} `form:"credential" json:"credential" xml:"credential"`
}

func GetRegister(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	options, sessionData, err := internalWebauthn.Web.BeginRegistration(&internalWebauthn.User{User: authedUser})
	if err != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	challenge, challengeId := db.Client.WebAuthnChallenge.Create().
		SetChallenge(sessionData.Challenge).
		SetUserId(sessionData.UserID).
		SetAllowedCredentialIds(sessionData.AllowedCredentialIDs).
		SetUserVerification(string(sessionData.UserVerification)).
		SetExtensions(map[string]any(sessionData.Extensions)).
		SetUser(authedUser).
		Save(db.Context)

	if challengeId != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	c.JSON(200, gin.H{"webauthnChallengeId": challenge.ID.String(), "options": options})
}

func PostRegister(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input PostRegisterInput

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

	foundWebauthnChallenge, fwcErr := db.GetUserWebauthnChallenge(authedUser, decodedChallengeId)
	if fwcErr != nil {
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

	credentialData, cdErr := protocol.ParseCredentialCreationResponseBody(dataReader)
	if cdErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	credential, credentialErr := internalWebauthn.Web.CreateCredential(&internalWebauthn.User{User: authedUser}, sessionData, credentialData)
	if credentialErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	transport := make([]string, len(credential.Transport))
	for tIndex, t := range credential.Transport {
		transport[tIndex] = string(t)
	}

	webauthnCredential, wcErr := db.Client.WebAuthnCredential.Create().
		SetCredentialId(credential.ID).
		SetPublicKey(credential.PublicKey).
		SetAttestationType(credential.AttestationType).
		SetTransport(transport).
		SetAaguid(credential.Authenticator.AAGUID).
		SetSignCount(credential.Authenticator.SignCount).
		SetCloneWarning(credential.Authenticator.CloneWarning).
		SetName(input.Name).
		SetUser(authedUser).
		Save(db.Context)

	if wcErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	authedUser.Update().SetDefault2FA(user.Default2FAWebauthn).Exec(db.Context)
	db.Client.WebAuthnChallenge.DeleteOne(foundWebauthnChallenge).Exec(db.Context)

	c.JSON(200, gin.H{"webauthnCredentialId": webauthnCredential.ID.String()})
}
