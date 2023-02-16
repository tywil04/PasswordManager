package webauthn

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/ent"
)

type GetCredentialInput struct {
	WebauthnCredentialId string `form:"webauthnCredentialId" json:"webauthnCredentialId" xml:"webauthnCredentialId"`
}

type DeleteCredentialInput struct {
	WebauthnCredentialId string `form:"webauthnCredentialId" json:"webauthnCredentialId" xml:"webauthnCredentialId"`
}

func GetCredential(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input GetCredentialInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.WebauthnCredentialId == "" {
		credentials, credentialErr := db.GetUserWebauthnCredentials(authedUser)
		if credentialErr != nil {
			c.JSON(500, helpers.ErrorUnknown())
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
	} else {
		decodedCredentialId, dciErr := uuid.Parse(input.WebauthnCredentialId)
		if dciErr != nil {
			c.JSON(400, helpers.ErrorInvalid("webauthnCredentialId"))
			return
		}

		credential, credentialErr := db.GetUserWebauthnCredential(authedUser, decodedCredentialId)
		if credentialErr != nil {
			c.JSON(400, helpers.ErrorInvalid("webauthnCredential"))
			return
		}

		jsonCredential := gin.H{
			"id":        credential.ID.String(),
			"name":      credential.Name,
			"createdAt": credential.CreatedAt,
		}

		c.JSON(200, gin.H{"webauthnCredential": jsonCredential})
	}
}

func DeleteCredential(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input DeleteCredentialInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.WebauthnCredentialId == "" {
		c.JSON(400, helpers.ErrorMissing("webauthnCredentialId"))
		return
	}

	decodedCredentialId, dciErr := uuid.Parse(input.WebauthnCredentialId)
	if dciErr != nil {
		c.JSON(400, helpers.ErrorInvalid("webauthnCredentialId"))
		return
	}

	dcErr := db.DeleteUserWebauthnCredentialViaId(authedUser, decodedCredentialId)
	if dcErr != nil {
		c.JSON(400, helpers.ErrorDeleting("webauthnCredential"))
		return
	}

	numberOfCredentials, nocErr := db.CountUserWebauthnCredentials(authedUser)
	if nocErr != nil {
		c.JSON(500, helpers.ErrorUnknown())
		return
	}

	if numberOfCredentials == 0 {
		db.Client.User.UpdateOne(authedUser).SetWebauthnEnabled(false).Exec(db.Context)
	}

	c.JSON(200, gin.H{})
}
