package helpers

import (
	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/ent"
	entChallenge "PasswordManager/ent/challenge"
	"encoding/base64"
	"strings"
	"time"

	"github.com/google/uuid"
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

func ValidateSession(authToken string) (bool, *ent.User, *ent.Session) {
	parts := strings.Split(authToken, ";")
	if len(parts) != 3 {
		return false, nil, nil
	}

	sessionId := parts[0]
	encodedSalt := parts[1]
	signature := parts[2]

	decodedSessionId, dsiErr := base64.StdEncoding.DecodeString(sessionId)
	if dsiErr != nil {
		return false, nil, nil
	}

	decodedSignature, dsErr := base64.StdEncoding.DecodeString(signature)
	if dsErr != nil {
		return false, nil, nil
	}

	parsedDecodedSessionId, pdsiErr := uuid.Parse(string(decodedSessionId[:]))
	if pdsiErr != nil {
		return false, nil, nil
	}

	session, sessionErr := db.Client.Session.Get(db.Context, parsedDecodedSessionId)
	if sessionErr != nil {
		return false, nil, nil
	}

	if session.Expiry.Before(time.Now()) {
		db.DeleteSession(session)
		return false, nil, nil
	}

	user, userErr := session.QueryUser().First(db.Context)
	if userErr != nil {
		return false, nil, nil
	}

	publicKey := cryptography.ImportPublicKey(session.N, session.E)
	valid := cryptography.VerifySignature(publicKey, decodedSignature, user.Email+base64.StdEncoding.EncodeToString(user.StrengthenedMasterHash)+encodedSalt)

	if valid {
		return true, user, session
	}

	return false, nil, nil
}
