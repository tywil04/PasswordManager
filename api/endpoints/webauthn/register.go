package webauthn

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	internalWebauthn "PasswordManager/api/lib/webauthn"
	"PasswordManager/ent"
)

type GetRegisterInput struct{}

type PostRegisterInput struct {
	WebauthnRegisterChallengeId string `form:"webauthnChallengeId" json:"webauthnChallengeId" xml:"webauthnChallengeId"`
	Name                        string `form:"name" json:"name" xml:"json"`
	Credential                  struct {
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
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	challenge, challengeId := db.Client.WebAuthnRegisterChallenge.Create().
		SetSdChallenge(sessionData.Challenge).
		SetUserId(sessionData.UserID).
		SetAllowedCredentialIds(sessionData.AllowedCredentialIDs).
		SetUserVerification(string(sessionData.UserVerification)).
		SetExtensions(map[string]any(sessionData.Extensions)).
		SetUser(authedUser).
		Save(db.Context)

	if challengeId != nil {
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	c.JSON(200, gin.H{"webauthnRegisterChallengeId": challenge.ID.String(), "options": options})
}

func PostRegister(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input PostRegisterInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.WebauthnRegisterChallengeId == "" {
		c.JSON(400, helpers.ErrorMissing("webauthnRegisterChallengeId"))
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.WebauthnRegisterChallengeId)
	if dciErr != nil {
		c.JSON(400, helpers.ErrorInvalid("webauthnRegisterChallengeId"))
		return
	}

	foundWebauthnChallenge, fwcErr := db.GetUserWebauthnRegisterChallengeViaId(authedUser, decodedChallengeId)
	if fwcErr != nil {
		c.JSON(400, helpers.ErrorInvalid("webauthnRegisterChallenge"))
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

	credentialData, cdErr := protocol.ParseCredentialCreationResponseBody(dataReader)
	if cdErr != nil {
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	credential, credentialErr := internalWebauthn.Web.CreateCredential(&internalWebauthn.User{User: authedUser}, sessionData, credentialData)
	if credentialErr != nil {
		c.JSON(500, helpers.ErrorUnknown())
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
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	authedUser.Update().SetWebauthnEnabled(true).Exec(db.Context)
	db.Client.WebAuthnRegisterChallenge.DeleteOne(foundWebauthnChallenge).Exec(db.Context)

	c.JSON(200, gin.H{"webauthnCredentialId": webauthnCredential.ID.String()})
}
