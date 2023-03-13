package vaults

import (
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/ent"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	Description string = ""
)

type GetInput struct{}

type PostInput struct {
	Name     string `form:"name" json:"name" xml:"name" pmParseType:"base64"`
	NameIv   string `form:"nameIv" json:"nameIv" xml:"nameIv" pmParseType:"base64"`
	Colour   string `form:"colour" json:"colour" xml:"colour" pmParseType:"base64"`
	ColourIv string `form:"colourIv" json:"colourIv" xml:"colourIv" pmParseType:"base64"`
}

type DeleteInput struct {
	VaultId string `form:"vaultId" json:"vaultId" xml:"vaultId" pmParseType:"uuid"`
}

func Get(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	vaults, vaultsErr := db.GetUserVaults(authedUser)
	if vaultsErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	processedVaults := make([]gin.H, len(vaults))
	for index, vault := range vaults {
		processedVaults[index] = gin.H{
			"id":       vault.ID.String(),
			"name":     vault.Name,
			"nameIv":   vault.NameIv,
			"colour":   vault.Colour,
			"colourIv": vault.ColourIv,
		}
	}

	c.JSON(200, gin.H{"vaults": processedVaults})
}

func Post(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	newVault, newVaultErr := db.Client.Vault.Create().
		SetUser(authedUser).
		SetName(params["name"].([]byte)).
		SetNameIv(params["nameIv"].([]byte)).
		SetColour(params["colour"].([]byte)).
		SetColourIv(params["colourIv"].([]byte)).
		Save(db.Context)

	if newVaultErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"vaultId": newVault.ID.String()})
}

func Delete(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	dvErr := db.DeleteUserVaultViaId(authedUser, params["vaultId"].(uuid.UUID))
	if dvErr != nil {
		c.JSON(400, exceptions.Builder("password", exceptions.Deleting))
		return
	}

	c.JSON(200, gin.H{})
}
