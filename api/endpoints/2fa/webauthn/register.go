package webauthn

import (
	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	internalWebauthn "PasswordManager/api/lib/webauthn"
	"PasswordManager/ent"
)

type GetRegisterInput struct{}

type PostRegisterInput struct {
	WebauthnRegisterChallengeId string `form:"webauthnRegisterChallengeId" json:"webauthnRegisterChallengeId" xml:"webauthnRegisterChallengeId" pmParseType:"uuid"`
	Name                        string `form:"name" json:"name" xml:"json"`
	Credential                  struct {
		AuthenticatorAttachment string `form:"authenticatorAttachment" json:"authenticatorAttachment" xml:"authenticatorAttachment" pmOptional:"true"`
		Id                      string `form:"id" json:"id" xml:"id" pmOptional:"true"`
		RawId                   string `form:"rawId" json:"rawId" xml:"rawId" pmOptional:"true"`
		Response                struct {
			AttestationObject string `form:"attestationObject" json:"attestationObject" xml:"attestationObject" pmOptional:"true"`
			ClientDataJSON    string `form:"clientDataJSON" json:"clientDataJSON" xml:"clientDataJSON" pmOptional:"true"`
		} `form:"response" json:"response" xml:"response" pmOptional:"true"`
		Type string `form:"type" json:"type" xml:"type" pmOptional:"true"`
	} `form:"credential" json:"credential" xml:"credential"`
}

func GetRegister(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	options, sessionData, err := internalWebauthn.Web.BeginRegistration(&internalWebauthn.User{User: authedUser})
	if err != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
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
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"webauthnRegisterChallengeId": challenge.ID.String(), "options": options})
}

func PostRegister(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	foundWebauthnChallenge, fwcErr := db.GetUserWebauthnRegisterChallengeViaId(authedUser, params["webauthnRegisterChallengeId"].(uuid.UUID))
	if fwcErr != nil {
		c.JSON(400, exceptions.Builder("webauthnRegisterChallenge", exceptions.Invalid))
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

	credentialData, cdErr := protocol.ParseCredentialCreationResponseBody(dataReader)
	if cdErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	credential, credentialErr := internalWebauthn.Web.CreateCredential(&internalWebauthn.User{User: authedUser}, sessionData, credentialData)
	if credentialErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
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
		SetName(params["name"].(string)).
		SetUser(authedUser).
		Save(db.Context)

	if wcErr != nil {
		c.JSON(500, exceptions.Builder("webauthnCredential", exceptions.Creating, exceptions.TryAgain))
		return
	}

	authedUser.Update().SetWebauthnEnabled(true).Exec(db.Context)
	db.Client.WebAuthnRegisterChallenge.DeleteOne(foundWebauthnChallenge).Exec(db.Context)

	c.JSON(200, gin.H{"webauthnCredentialId": webauthnCredential.ID.String()})
}
