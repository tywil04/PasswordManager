package db

import (
	"PasswordManager/ent"
	"PasswordManager/ent/note"
	"PasswordManager/ent/password"
	"PasswordManager/ent/session"
	"PasswordManager/ent/user"
	"PasswordManager/ent/vault"
	"PasswordManager/ent/webauthncredential"
	"PasswordManager/ent/webauthnregisterchallenge"

	"github.com/google/uuid"
)

// Vault
func GetUserVaults(user *ent.User) ([]*ent.Vault, error) {
	return user.QueryVaults().All(Context)
}

func GetUserVault(user *ent.User, vaultId uuid.UUID) (*ent.Vault, error) {
	return user.QueryVaults().Where(vault.IDEQ(vaultId)).Unique(true).First(Context)
}

func DeleteVault(vault *ent.Vault) error {
	return Client.Vault.DeleteOne(vault).Exec(Context)
}

func DeleteUserVaultViaId(user *ent.User, vaultId uuid.UUID) error {
	vault, vErr := GetUserVault(user, vaultId)
	if vErr != nil || vault == nil {
		return vErr
	}
	return DeleteVault(vault)
}

// note
func DeleteNote(nodeId uuid.UUID) error {
	return Client.Note.DeleteOneID(nodeId).Exec(Context)
}

func GetUserNotes(user *ent.User) ([]*ent.Note, error) {
	return user.QueryVaults().QueryNotes().All(Context)
}

func GetVaultNotes(vault *ent.Vault) ([]*ent.Note, error) {
	return vault.QueryNotes().All(Context)
}

func GetVaultNote(vault *ent.Vault, noteId uuid.UUID) (*ent.Note, error) {
	return vault.QueryNotes().Where(note.IDEQ(noteId)).Unique(true).First(Context)
}

func DeleteVaultNoteViaId(vault *ent.Vault, noteId uuid.UUID) error {
	password, pErr := GetVaultNote(vault, noteId)
	if pErr != nil || password == nil {
		return pErr
	}
	return DeleteNote(noteId)
}

func GetNoteVault(note *ent.Note) (*ent.Vault, error) {
	return note.QueryVault().Unique(true).First(Context)
}

// Password
func GetPassword(passwordId uuid.UUID) (*ent.Password, error) {
	return Client.Password.Get(Context, passwordId)
}

func GetUserPasswords(user *ent.User) ([]*ent.Password, error) {
	return user.QueryVaults().QueryPasswords().All(Context)
}

func DeletePasswordViaId(passwordId uuid.UUID) error {
	return Client.Password.DeleteOneID(passwordId).Exec(Context)
}

func DeletePassword(password *ent.Password) error {
	return Client.Password.DeleteOne(password).Exec(Context)
}

func GetVaultPasswords(vault *ent.Vault) ([]*ent.Password, error) {
	return vault.QueryPasswords().All(Context)
}

func GetVaultPassword(vault *ent.Vault, passwordId uuid.UUID) (*ent.Password, error) {
	return vault.QueryPasswords().Where(password.IDEQ(passwordId)).Unique(true).First(Context)
}

func DeleteVaultPasswordViaId(vault *ent.Vault, passwordId uuid.UUID) error {
	password, pErr := GetVaultPassword(vault, passwordId)
	if pErr != nil || password == nil {
		return pErr
	}
	return DeletePassword(password)
}

func GetPasswordAdditionalFields(password *ent.Password) ([]*ent.AdditionalField, error) {
	return password.QueryAdditionalFields().All(Context)
}

func GetPasswordUrls(password *ent.Password) ([]*ent.Url, error) {
	return password.QueryUrls().All(Context)
}

func GetPasswordVault(password *ent.Password) (*ent.Vault, error) {
	return password.QueryVault().Unique(true).First(Context)
}

// Webauthn Credential

// func GetWebauthnCredential(webauthnCredentialId uuid.UUID) (*ent.WebAuthnCredential, error) {
// 	return Client.WebAuthnCredential.Get(Context, webauthnCredentialId)
// }

// func DeleteWebauthnCredentialViaId(webauthnCredentialId uuid.UUID) error {
// 	return Client.WebAuthnCredential.DeleteOneID(webauthnCredentialId).Exec(Context)
// }

func DeleteWebauthnCredential(webauthnCredential *ent.WebAuthnCredential) error {
	return Client.WebAuthnCredential.DeleteOne(webauthnCredential).Exec(Context)
}

func GetUserWebauthnCredentials(user *ent.User) ([]*ent.WebAuthnCredential, error) {
	return user.QueryWebauthnCredentials().All(Context)
}

func GetUserWebauthnCredential(user *ent.User, webauthnCredentialId uuid.UUID) (*ent.WebAuthnCredential, error) {
	return user.QueryWebauthnCredentials().Where(webauthncredential.IDEQ(webauthnCredentialId)).Unique(true).First(Context)
}

func CountUserWebauthnCredentials(user *ent.User) (int, error) {
	return user.QueryWebauthnCredentials().Count(Context)
}

func DeleteUserWebauthnCredentialViaId(user *ent.User, webauthnCredentialId uuid.UUID) error {
	webauthnCredential, wcErr := GetUserWebauthnCredential(user, webauthnCredentialId)
	if wcErr != nil || webauthnCredential == nil {
		return wcErr
	}
	return DeleteWebauthnCredential(webauthnCredential)
}

// Challenges
func GetChallenge(challengeId uuid.UUID) (*ent.Challenge, error) {
	return Client.Challenge.Get(Context, challengeId)
}

func GetChallengeUser(challenge *ent.Challenge) (*ent.User, error) {
	return challenge.QueryUser().First(Context)
}

func DeleteChallenge(challenge *ent.Challenge) error {
	emailChallenge, ecErr := GetChallengeEmailChallenge(challenge)
	if ecErr == nil {
		DeleteEmailChallenge(emailChallenge)
	}

	webauthnChallenge, wcErr := GetChallengeWebauthnChallenge(challenge)
	if wcErr == nil {
		DeleteWebauthnChallenge(webauthnChallenge)
	}

	return Client.Challenge.DeleteOne(challenge).Exec(Context)
}

// Webauthn Challenge
func GetChallengeWebauthnChallenge(challenge *ent.Challenge) (*ent.WebAuthnChallenge, error) {
	return challenge.QueryWebauthnChallenge().Unique(true).First(Context)
}

func DeleteWebauthnChallengeViaId(webauthnChallengeId uuid.UUID) error {
	return Client.WebAuthnChallenge.DeleteOneID(webauthnChallengeId).Exec(Context)
}

func DeleteWebauthnChallenge(webauthnChallenge *ent.WebAuthnChallenge) error {
	return Client.WebAuthnChallenge.DeleteOne(webauthnChallenge).Exec(Context)
}

// Webauthn Register Challenges
func GetUserWebauthnRegisterChallengeViaId(user *ent.User, webauthnRegisterChallengeId uuid.UUID) (*ent.WebAuthnRegisterChallenge, error) {
	return user.QueryWebauthnRegisterChallenges().Where(webauthnregisterchallenge.IDEQ(webauthnRegisterChallengeId)).Unique(true).First(Context)
}

// Totp Challenges
func GetChallengeTotpCredential(challenge *ent.Challenge) (*ent.TotpCredential, error) {
	return challenge.QueryTotpCredential().Unique(true).First(Context)
}

func GetUserTotpCredential(user *ent.User) (*ent.TotpCredential, error) {
	return user.QueryTotpCredential().Unique(true).First(Context)
}

// Email Challenges
func GetChallengeEmailChallenge(challenge *ent.Challenge) (*ent.EmailChallenge, error) {
	return challenge.QueryEmailChallenge().Unique(true).First(Context)
}

func DeleteEmailChallengeViaId(emailChallengeId uuid.UUID) error {
	return Client.EmailChallenge.DeleteOneID(emailChallengeId).Exec(Context)
}

func DeleteEmailChallenge(emailChallenge *ent.EmailChallenge) error {
	return Client.EmailChallenge.DeleteOne(emailChallenge).Exec(Context)
}

// Session

func GetUserSession(user *ent.User, sessionId uuid.UUID) (*ent.Session, error) {
	return user.QuerySessions().Where(session.IDEQ(sessionId)).Unique(true).First(Context)
}

func DeleteUserSession(user *ent.User, session *ent.Session) error {
	session, sErr := GetUserSession(user, session.ID)
	if sErr != nil {
		return sErr
	}
	return Client.Session.DeleteOne(session).Exec(Context)
}

// User

func GetUser(userId uuid.UUID) (*ent.User, error) {
	return Client.User.Get(Context, userId)
}

func GetUserViaEmail(email string) (*ent.User, error) {
	return Client.User.Query().Where(user.EmailEQ(email)).First(Context)
}

// func DeletetUserId(userId uuid.UUID) error {
// 	return Client.User.DeleteOneID(userId).Exec(Context)
// }

// func DeletetUser(user *ent.User) error {
// 	return Client.User.DeleteOne(user).Exec(Context)
// }
