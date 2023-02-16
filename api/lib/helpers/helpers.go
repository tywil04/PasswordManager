package helpers

import (
	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/ent"
	entChallenge "PasswordManager/ent/challenge"
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateEmailChallenge(challenge *ent.Challenge) (*ent.EmailChallenge, error) {
	eChallenge, eChallengeErr := db.Client.EmailChallenge.Create().
		SetChallenge(challenge).
		Save(db.Context)

	return eChallenge, eChallengeErr
}

func GenerateWebauthnChallenge(challenge *ent.Challenge) (*ent.WebAuthnChallenge, error) {
	wChallenge, wChallengeErr := db.Client.WebAuthnChallenge.Create().
		SetChallenge(challenge).
		Save(db.Context)

	return wChallenge, wChallengeErr
}

func GenerateChallenge(user *ent.User) (*ent.Challenge, error) {
	oldChallenge, ocErr := user.QueryChallenges().Where(entChallenge.ExpiryGT(time.Now())).First(db.Context)
	if ocErr == nil {
		return oldChallenge, nil
	}

	challengeBuilder := db.Client.Challenge.Create().
		SetUser(user)

	if user.TotpEnabled {
		totp, totpErr := user.QueryTotpCredential().First(db.Context)
		if totpErr != nil {
			return nil, totpErr
		}
		challengeBuilder.SetTotpCredential(totp)
	}

	challenge, challengeErr := challengeBuilder.Save(db.Context)

	if challengeErr != nil {
		return nil, challengeErr
	}

	_, ecErr := GenerateEmailChallenge(challenge)
	if ecErr != nil {
		return nil, ecErr
	}

	if user.WebauthnEnabled {
		_, wcErr := GenerateWebauthnChallenge(challenge)
		if wcErr != nil {
			return nil, wcErr
		}
	}

	return challenge, nil
}

func GetAvailableChallenges(user *ent.User) []string {
	available := []string{"email"}

	if user.TotpEnabled {
		available = append(available, "totp")
	}

	if user.WebauthnEnabled {
		available = append(available, "webauthn")
	}

	return available
}

func GenerateSession(user *ent.User) (string, string, string, error) {
	salt := cryptography.RandomBytes(32)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	publicKey, signature := cryptography.GenerateSignature(user.Email + base64.StdEncoding.EncodeToString(user.StrengthenedMasterHash) + encodedSalt)

	session, sessionErr := db.Client.Session.Create().
		SetUser(user).
		SetN(publicKey.N.Bytes()).
		SetE(publicKey.E).
		SetExpiry(time.Now().Add(time.Hour)).
		Save(db.Context)

	if sessionErr != nil {
		return "", "", "", sessionErr
	}

	token := base64.StdEncoding.EncodeToString([]byte(session.ID.String())) + ";" + encodedSalt + ";" + base64.StdEncoding.EncodeToString(signature)
	encodedProtectedDatabaseKey := base64.StdEncoding.EncodeToString(user.ProtectedDatabaseKey)
	encodedProtectedDatabaseKeyIv := base64.StdEncoding.EncodeToString(user.ProtectedDatabaseKeyIv)

	return token, encodedProtectedDatabaseKey, encodedProtectedDatabaseKeyIv, nil
}

func ErrorMissing(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorMissing", "subject": subject}}
}

func ErrorInvalid(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorInvalid", "subject": subject}}
}

func ErrorExpired(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorExpired", "subject": subject}}
}

func ErrorInUse(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorInUse", "subject": subject}}
}

func ErrorNotInUse(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorNotInUse", "subject": subject}}
}

func ErrorIssuing(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorIssuing", "subject": subject}}
}

func ErrorDeleting(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorDeleting", "subject": subject}}
}

func ErrorCreating(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorCreating", "subject": subject}}
}

func ErrorInvalidCredentials() gin.H {
	return gin.H{"error": gin.H{"code": "errorInvalidCredentials"}}
}

func ErrorChallenge(subject string) gin.H {
	return gin.H{"error": gin.H{"code": "errorChallenge", "subject": subject, "message": "Failed challenge."}}
}

func ErrorUnknown() gin.H {
	return gin.H{"error": gin.H{"code": "errorUnknown", "message": "Please try again later."}}
}
