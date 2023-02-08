package webauthn

import (
	"PasswordManager/api/lib/db"
	"PasswordManager/ent"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	User *ent.User
}

func (u User) WebAuthnID() []byte {
	return u.User.ID[:]
}

func (u User) WebAuthnName() string {
	return u.User.Email
}

func (u User) WebAuthnDisplayName() string {
	return u.User.Email
}

func (u User) WebAuthnIcon() string {
	return ""
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	credentials, _ := u.User.QueryWebauthnCredentials().All(db.Context)
	webauthnCredentials := make([]webauthn.Credential, len(credentials))

	for index, credential := range credentials {
		transport := make([]protocol.AuthenticatorTransport, len(credential.Transport))
		for tIndex, t := range credential.Transport {
			transport[tIndex] = (protocol.AuthenticatorTransport)(t)
		}

		webauthnCredentials[index] = webauthn.Credential{
			ID:              credential.CredentialId,
			PublicKey:       credential.PublicKey,
			AttestationType: credential.AttestationType,
			Transport:       transport,
			Authenticator: webauthn.Authenticator{
				AAGUID:       credential.Aaguid,
				SignCount:    credential.SignCount,
				CloneWarning: credential.CloneWarning,
			},
		}
	}

	return webauthnCredentials
}
