package webauthn

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/ent"
)

const (
	CredentialDescription string = ""
)

type GetCredentialInput struct{}

type DeleteCredentialInput struct {
	WebauthnCredentialId string `form:"webauthnCredentialId" json:"webauthnCredentialId" xml:"webauthnCredentialId" pmParseType:"uuid"`
}

func GetCredential(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	credentials, credentialErr := db.GetUserWebauthnCredentials(authedUser)
	if credentialErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	jsonCredentials := make([]gin.H, len(credentials))
	for index, credential := range credentials {
		jsonCredentials[index] = gin.H{
			"id":        credential.ID.String(),
			"name":      credential.Name,
			"createdAt": credential.CreatedAt,
		}
	}

	c.JSON(200, gin.H{"webauthnCredentials": jsonCredentials})
}

func DeleteCredential(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	dcErr := db.DeleteUserWebauthnCredentialViaId(authedUser, params["webauthnCredentialId"].(uuid.UUID))
	if dcErr != nil {
		c.JSON(500, exceptions.Builder("webauthnCredential", exceptions.Deleting, exceptions.TryAgain))
		return
	}

	numberOfCredentials, nocErr := db.CountUserWebauthnCredentials(authedUser)
	if nocErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	if numberOfCredentials == 0 {
		db.Client.User.UpdateOne(authedUser).SetWebauthnEnabled(false).Exec(db.Context)
	}

	c.JSON(200, gin.H{})
}
