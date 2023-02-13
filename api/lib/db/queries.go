package db

import (
	"PasswordManager/ent"
	"PasswordManager/ent/password"
	"PasswordManager/ent/session"
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthnchallenge"
	"PasswordManager/ent/webauthncredential"

	"github.com/google/uuid"
)

// Password

func GetPassword(passwordId uuid.UUID) (*ent.Password, error) {
	return Client.Password.Get(Context, passwordId)
}

func DeletePasswordViaId(passwordId uuid.UUID) error {
	return Client.Password.DeleteOneID(passwordId).Exec(Context)
}

func DeletePassword(password *ent.Password) error {
	return Client.Password.DeleteOne(password).Exec(Context)
}

func GetUserPasswords(user *ent.User) ([]*ent.Password, error) {
	return user.QueryPasswords().All(Context)
}

func GetUserPassword(user *ent.User, passwordId uuid.UUID) (*ent.Password, error) {
	return user.QueryPasswords().Where(password.IDEQ(passwordId)).Unique(true).First(Context)
}

func DeleteUserPasswordViaId(user *ent.User, passwordId uuid.UUID) error {
	password, pErr := GetUserPassword(user, passwordId)
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

// Webauthn Challenge

func GetWebauthnChallenge(webauthnChallengeId uuid.UUID) (*ent.WebAuthnChallenge, error) {
	return Client.WebAuthnChallenge.Get(Context, webauthnChallengeId)
}

func DeleteWebauthnChallengeViaId(webauthnChallengeId uuid.UUID) error {
	return Client.WebAuthnChallenge.DeleteOneID(webauthnChallengeId).Exec(Context)
}

func DeleteWebauthnChallenge(webauthnChallenge *ent.WebAuthnChallenge) error {
	return Client.WebAuthnChallenge.DeleteOne(webauthnChallenge).Exec(Context)
}

func GetUserWebauthnChallenges(user *ent.User) ([]*ent.WebAuthnChallenge, error) {
	return user.QueryWebauthnChallenges().All(Context)
}

func GetUserWebauthnChallenge(user *ent.User, webauthnChallengeId uuid.UUID) (*ent.WebAuthnChallenge, error) {
	return user.QueryWebauthnChallenges().Where(webauthnchallenge.IDEQ(webauthnChallengeId)).Unique(true).First(Context)
}

func GetWebauthnChallengeUser(webauthnChallenge *ent.WebAuthnChallenge) (*ent.User, error) {
	return webauthnChallenge.QueryUser().Unique(true).First(Context)
}

// Email Challenges

func GetEmailChallenge(emailChallengeId uuid.UUID) (*ent.EmailChallenge, error) {
	return Client.EmailChallenge.Get(Context, emailChallengeId)
}

func DeleteEmailChallengeViaId(emailChallengeId uuid.UUID) error {
	return Client.EmailChallenge.DeleteOneID(emailChallengeId).Exec(Context)
}

func DeleteEmailChallenge(emailChallenge *ent.EmailChallenge) error {
	return Client.EmailChallenge.DeleteOne(emailChallenge).Exec(Context)
}

func GetEmailChallengeUser(emailChallenge *ent.EmailChallenge) (*ent.User, error) {
	return emailChallenge.QueryUser().Unique(true).First(Context)
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
