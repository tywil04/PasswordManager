package exceptions

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Addition int64

const (
	Undefined Addition = iota
	Base64
	Uuid
	HexColour
	EmailCode
	TotpCode
	Email
	JsonOrXml
	Expired
	InUse
	NotInUse
	Owns
	Authtoken
	DidntStart2FA
	TryAgain
)

func (a Addition) String() string {
	switch a {
	case Base64:
		return "Are you sure it's base64 encoded?"
	case Uuid:
		return "Are you sure it's a well-formed uuid?"
	case HexColour:
		return "Are you sure it's a valid hex colour code?"
	case EmailCode:
		return "Are you sure it's a valid email code?"
	case TotpCode:
		return "Are you sure it's a valid totp code?"
	case Email:
		return "Are you sure it's a valid email address?"
	case JsonOrXml:
		return "Are you sure it's valid JSON or XML?"
	case Expired:
		return "Are you sure it hasn't expired?"
	case InUse:
		return "Are you sure it isn't in use?"
	case NotInUse:
		return "Are you sure it's in use?"
	case Owns:
		return "Are you sure you own it?"
	case Authtoken:
		return "Are you sure it's both a non-expired and valid authToken?"
	case DidntStart2FA:
		return "Are you sure you started the 2FA process? Some challenges require you to do so."
	case TryAgain:
		return "Please try again later."
	}
	return ""
}

type Error int64

const (
	MissingParam Error = iota
	InvalidParam
	InvalidHeader
	ParsingParam
	Invalid
	Creating
	Updating
	Issuing
	Deleting
	IncorrectCredentials
	IncorrectChallenge
	Unknown
)

func (e Error) String() (string, string) {
	switch e {
	case MissingParam:
		return "missingParam", "Missing parameter '%s'."
	case InvalidParam:
		return "invalidParam", "Invalid parameter '%s'."
	case InvalidHeader:
		return "invalidHeader", "Invalid header '%s'."
	case ParsingParam:
		return "parsingParam", "Failed to parse parameter '%s'."
	case Invalid:
		return "invalid", "Invalid %s."
	case Creating:
		return "creating", "Failed to create %s."
	case Updating:
		return "updating", "Failed to update %s."
	case Issuing:
		return "issuing", "Failed to issue %s."
	case Deleting:
		return "deleting", "Failed to delete %s."
	case IncorrectCredentials:
		return "credentials", "Incorrect credentials."
	case IncorrectChallenge:
		return "challenge", "Incorrect challenge response '%s'."
	case Unknown:
		return "unknown", "An unknown error has occured."
	}
	return "", ""
}

func Builder(agent string, e Error, additions ...Addition) gin.H {
	etName, etMessage := e.String()
	message := fmt.Sprintf(etMessage, agent)

	for _, addition := range additions {
		if addition != Undefined {
			message += " " + addition.String()
		}
	}

	return gin.H{"error": gin.H{"type": etName, "message": message}}
}
