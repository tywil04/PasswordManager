package webauthn

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/ent"
	"PasswordManager/ent/user"
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
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.WebauthnCredentialId == "" {
		credentials, credentialErr := db.GetUserWebauthnCredentials(authedUser)
		if credentialErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
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
			c.JSON(400, gin.H{"error": gin.H{"code": "errParsingWebauthnCredentialId", "message": "Unable to parse 'webauthnCredentialId', expected uuid."}})
			return
		}

		credential, credentialErr := db.GetUserWebauthnCredential(authedUser, decodedCredentialId)
		if credentialErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errWebauthnCredentialNotFound", "message": "Unable to find valid webauthn credential using 'webauthnCredentialId'."}})
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
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.WebauthnCredentialId == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingWebauthnCredentialId", "message": "Required 'webauthnCredentialId' was not found."}})
		return
	}

	decodedCredentialId, dciErr := uuid.Parse(input.WebauthnCredentialId)
	if dciErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingWebauthnCredentialId", "message": "Unable to parse 'webauthnCredentialId', expected uuid."}})
		return
	}

	dcErr := db.DeleteUserWebauthnCredentialViaId(authedUser, decodedCredentialId)
	if dcErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errDeletingWebauthnCredential", "message": "Unable to delete webauthn credential using 'webauthnCredentialId', is the webauthn credential yours?"}})
		return
	}

	numberOfCredentials, nocErr := db.CountUserWebauthnCredentials(authedUser)
	if nocErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	if numberOfCredentials == 0 {
		db.Client.User.UpdateOne(authedUser).SetDefault2FA(user.Default2FAEmail).Exec(db.Context)
	}

	c.JSON(200, gin.H{})
}
