package vault

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/ent"
)

const (
	Description string = ""
)

type GetPasswordInput struct {
	VaultId string `form:"vaultId" json:"vaultId" xml:"vaultId" pmParseType:"uuid"`
}

type PostPasswordInput struct {
	VaultId          string `form:"vaultId" json:"vaultId" xml:"vaultId" pmParseType:"uuid"`
	Name             string `form:"name" json:"name" xml:"name" pmParseType:"base64"`
	NameIv           string `form:"nameIv" json:"nameIv" xml:"nameIv" pmParseType:"base64"`
	Username         string `form:"username" json:"username" xml:"username" pmParseType:"base64"`
	UsernameIv       string `form:"usernameIv" json:"usernameIv" xml:"usernameIv" pmParseType:"base64"`
	Password         string `form:"password" json:"password" xml:"password" pmParseType:"base64"`
	PasswordIv       string `form:"passwordIv" json:"passwordIv" xml:"passwordIv" pmParseType:"base64"`
	Colour           string `form:"colour" json:"colour" xml:"colour" pmParseType:"hexcolour"`
	AdditionalFields []struct {
		Key     string `form:"key" json:"key" xml:"key" pmParseType:"base64"`
		KeyIv   string `form:"keyIv" json:"keyIv" xml:"keyIv" pmParseType:"base64"`
		Value   string `form:"value" json:"value" xml:"value" pmParseType:"base64"`
		ValueIv string `form:"valueIv" json:"valueIv" xml:"valueIv" pmParseType:"base64"`
	} `form:"additionalFields" json:"additionalFields" xml:"additionalFields"`
	Urls []struct {
		Url   string `form:"url" json:"url" xml:"url" pmParseType:"base64"`
		UrlIv string `form:"urlIv" json:"urlIv" xml:"urlIv" pmParseType:"base64"`
	} `form:"urls" json:"urls" xml:"urls"`
}

type DeletePasswordInput struct {
	VaultId    string `form:"vaultId" json:"vaultId" xml:"vaultId" pmParseType:"uuid"`
	PasswordId string `form:"passwordId" json:"passwordId" xml:"passwordId"`
}

func GetPassword(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	vault, vaultErr := db.GetUserVault(authedUser, params["vaultId"].(uuid.UUID))
	if vaultErr != nil {
		c.JSON(400, exceptions.Builder("vaultId", exceptions.InvalidParam, exceptions.Uuid, exceptions.Owns))
		return
	}

	passwords, passwordsErr := db.GetVaultPasswords(vault)
	if passwordsErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	jsonPasswords := make([]gin.H, len(passwords))
	for index, password := range passwords {
		additionalFields, afErr := db.GetPasswordAdditionalFields(password)
		if afErr != nil {
			c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
			return
		}

		urls, uErr := db.GetPasswordUrls(password)
		if uErr != nil {
			c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
			return
		}

		jsonAdditionalFields := make([]gin.H, len(additionalFields))
		for afIndex, additionalField := range additionalFields {
			jsonAdditionalFields[afIndex] = gin.H{
				"key":     additionalField.Key,
				"keyIv":   additionalField.KeyIv,
				"value":   additionalField.Value,
				"valueIv": additionalField.ValueIv,
			}
		}

		jsonUrls := make([]gin.H, len(urls))
		for uIndex, url := range urls {
			jsonUrls[uIndex] = gin.H{
				"url":   url.URL,
				"urlIv": url.UrlIv,
			}
		}

		jsonPasswords[index] = gin.H{
			"id":               password.ID.String(),
			"name":             password.Name,
			"nameIv":           password.NameIv,
			"username":         password.Username,
			"usernameIv":       password.UsernameIv,
			"password":         password.Password,
			"passwordIv":       password.PasswordIv,
			"colour":           password.Colour,
			"additionalFields": jsonAdditionalFields,
			"urls":             jsonUrls,
		}
	}

	c.JSON(200, gin.H{"passwords": jsonPasswords})
}

func PostPassword(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	vault, vaultErr := db.GetUserVault(authedUser, params["vaultId"].(uuid.UUID))
	if vaultErr != nil {
		c.JSON(400, exceptions.Builder("vaultId", exceptions.InvalidParam, exceptions.Uuid, exceptions.Owns))
		return
	}

	additionalFields := params["additionalFields"].([]map[string]any)
	entAdditionalFields := make([]*ent.AdditionalField, len(additionalFields))
	for index, additionalField := range additionalFields {
		newAdditionalField, nafErr := db.Client.AdditionalField.Create().
			SetKey(additionalField["key"].([]byte)).
			SetKeyIv(additionalField["keyIv"].([]byte)).
			SetValue(additionalField["value"].([]byte)).
			SetValueIv(additionalField["valueIv"].([]byte)).
			Save(db.Context)

		if nafErr != nil {
			c.JSON(500, exceptions.Builder(fmt.Sprintf("Additional Field [%d]", index), exceptions.Creating, exceptions.TryAgain))
			return
		}

		entAdditionalFields[index] = newAdditionalField
	}

	urls := params["urls"].([]map[string]any)
	entUrls := make([]*ent.Url, len(urls))
	for index, url := range urls {
		newUrl, nuErr := db.Client.Url.Create().
			SetURL(url["url"].([]byte)).
			SetUrlIv(url["urlIv"].([]byte)).
			Save(db.Context)

		if nuErr != nil {
			c.JSON(500, exceptions.Builder(fmt.Sprintf("Url [%d]", index), exceptions.Creating, exceptions.TryAgain))
			return
		}

		entUrls[index] = newUrl
	}

	newPassword, newPasswordErr := db.Client.Password.Create().
		SetVault(vault).
		SetName(params["name"].([]byte)).
		SetNameIv(params["nameIv"].([]byte)).
		SetUsername(params["username"].([]byte)).
		SetUsernameIv(params["usernameIv"].([]byte)).
		SetPassword(params["password"].([]byte)).
		SetPasswordIv(params["passwordIv"].([]byte)).
		SetColour(params["colour"].(string)).
		AddAdditionalFields(entAdditionalFields...).
		AddUrls(entUrls...).
		Save(db.Context)

	if newPasswordErr != nil {
		c.JSON(500, exceptions.Builder("password", exceptions.Creating, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"passwordId": newPassword.ID.String()})
}

func DeletePassword(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	vault, vaultErr := db.GetUserVault(authedUser, params["vaultId"].(uuid.UUID))
	if vaultErr != nil {
		c.JSON(400, exceptions.Builder("vaultId", exceptions.InvalidParam, exceptions.Uuid, exceptions.Owns))
		return
	}

	dpErr := db.DeleteVaultPasswordViaId(vault, params["passwordId"].(uuid.UUID))
	if dpErr != nil {
		c.JSON(400, exceptions.Builder("password", exceptions.Deleting))
		return
	}

	c.JSON(200, gin.H{})
}
