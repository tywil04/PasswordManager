package password

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/ent"
)

type GetInput struct{}

type PostInput struct {
	Name             string `form:"name" json:"name" xml:"name"`
	NameIv           string `form:"nameIv" json:"nameIv" xml:"nameIv"`
	Username         string `form:"username" json:"username" xml:"username"`
	UsernameIv       string `form:"usernameIv" json:"usernameIv" xml:"usernameIv"`
	Password         string `form:"password" json:"password" xml:"password"`
	PasswordIv       string `form:"passwordIv" json:"passwordIv" xml:"passwordIv"`
	Colour           string `form:"colour" json:"colour" xml:"colour"`
	AdditionalFields []struct {
		Key     string `form:"key" json:"key" xml:"key"`
		KeyIv   string `form:"keyIv" json:"keyIv" xml:"keyIv"`
		Value   string `form:"value" json:"value" xml:"value"`
		ValueIv string `form:"valueIv" json:"valueIv" xml:"valueIv"`
	} `form:"additionalFields" json:"additionalFields" xml:"additionalFields"`
	Urls []struct {
		Url   string `form:"url" json:"url" xml:"url"`
		UrlIv string `form:"urlIv" json:"urlIv" xml:"urlIv"`
	} `form:"urls" json:"urls" xml:"urls"`
}

type DeleteInput struct {
	PasswordId string `form:"passwordId" json:"passwordId" xml:"passwordId"`
}

func Get(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	passwords, passwordsErr := db.GetUserPasswords(authedUser)
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

func Post(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input PostInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, exceptions.Builder("body", exceptions.Invalid, exceptions.JsonOrXml))
		return
	}

	if input.Name == "" {
		c.JSON(400, exceptions.Builder("name", exceptions.MissingParam))
		return
	}

	if input.NameIv == "" {
		c.JSON(400, exceptions.Builder("nameIv", exceptions.MissingParam))
		return
	}

	if input.Username == "" {
		c.JSON(400, exceptions.Builder("username", exceptions.MissingParam))
		return
	}

	if input.UsernameIv == "" {
		c.JSON(400, exceptions.Builder("usernameIv", exceptions.MissingParam))
		return
	}

	if input.Password == "" {
		c.JSON(400, exceptions.Builder("password", exceptions.MissingParam))
		return
	}

	if input.PasswordIv == "" {
		c.JSON(400, exceptions.Builder("passwordIv", exceptions.MissingParam))
		return
	}

	if input.Colour == "" {
		c.JSON(400, exceptions.Builder("colour", exceptions.MissingParam))
		return
	}

	decodedName, dnErr := base64.StdEncoding.DecodeString(input.Name)
	if dnErr != nil {
		c.JSON(400, exceptions.Builder("name", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	decodedNameIv, dniErr := base64.StdEncoding.DecodeString(input.NameIv)
	if dniErr != nil {
		c.JSON(400, exceptions.Builder("nameIv", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	decodedUsername, duErr := base64.StdEncoding.DecodeString(input.Username)
	if duErr != nil {
		c.JSON(400, exceptions.Builder("username", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	decodedUsernameIv, duiErr := base64.StdEncoding.DecodeString(input.UsernameIv)
	if duiErr != nil {
		c.JSON(400, exceptions.Builder("usernameIv", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	decodedPassword, dpErr := base64.StdEncoding.DecodeString(input.Password)
	if dpErr != nil {
		c.JSON(400, exceptions.Builder("password", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	decodedPasswordIv, dpiErr := base64.StdEncoding.DecodeString(input.PasswordIv)
	if dpiErr != nil {
		c.JSON(400, exceptions.Builder("passwordIv", exceptions.ParsingParam, exceptions.Base64))
		return
	}

	_, hexErr := strconv.ParseInt(input.Colour, 16, 64)
	if hexErr != nil {
		c.JSON(400, exceptions.Builder("colour", exceptions.ParsingParam, exceptions.HexColour))
		return
	}

	additionalFields := make([]*ent.AdditionalField, len(input.AdditionalFields))
	for index, additionalField := range input.AdditionalFields {
		decodedKey, dkErr := base64.StdEncoding.DecodeString(additionalField.Key)
		if dkErr != nil {
			exceptions.Builder(fmt.Sprintf("additionalFields[%d].key", index), exceptions.InvalidParam, exceptions.Base64)
			c.JSON(400, exceptions.Builder(fmt.Sprintf("additionalFields[%d].key", index), exceptions.InvalidParam, exceptions.Base64))
			return
		}

		decodedKeyIv, dkiErr := base64.StdEncoding.DecodeString(additionalField.KeyIv)
		if dkiErr != nil {
			c.JSON(400, exceptions.Builder(fmt.Sprintf("additionalFields[%d].keyIv", index), exceptions.InvalidParam, exceptions.Base64))
			return
		}

		decodedValue, dvErr := base64.StdEncoding.DecodeString(additionalField.Value)
		if dvErr != nil {
			c.JSON(400, exceptions.Builder(fmt.Sprintf("additionalFields[%d].value", index), exceptions.InvalidParam, exceptions.Base64))
			return
		}

		decodedValueIv, dviErr := base64.StdEncoding.DecodeString(additionalField.ValueIv)
		if dviErr != nil {
			c.JSON(400, exceptions.Builder(fmt.Sprintf("additionalFields[%d].valueIv", index), exceptions.InvalidParam, exceptions.Base64))
			return
		}

		newAdditionalField, nafErr := db.Client.AdditionalField.Create().
			SetKey(decodedKey).
			SetKeyIv(decodedKeyIv).
			SetValue(decodedValue).
			SetValueIv(decodedValueIv).
			Save(db.Context)

		if nafErr != nil {
			c.JSON(500, exceptions.Builder(fmt.Sprintf("Additional Field [%d]", index), exceptions.Creating, exceptions.TryAgain))
			return
		}

		additionalFields[index] = newAdditionalField
	}

	urls := make([]*ent.Url, len(input.Urls))
	for index, url := range input.Urls {
		decodedUrl, duErr := base64.StdEncoding.DecodeString(url.Url)
		if duErr != nil {
			c.JSON(400, exceptions.Builder(fmt.Sprintf("urls[%d].url", index), exceptions.InvalidParam, exceptions.Base64))
			return
		}

		decodedUrlIv, duiErr := base64.StdEncoding.DecodeString(url.UrlIv)
		if duiErr != nil {
			c.JSON(400, exceptions.Builder(fmt.Sprintf("urls[%d].urlIv", index), exceptions.InvalidParam, exceptions.Base64))
			return
		}

		newUrl, nuErr := db.Client.Url.Create().
			SetURL(decodedUrl).
			SetUrlIv(decodedUrlIv).
			Save(db.Context)

		if nuErr != nil {
			c.JSON(500, exceptions.Builder(fmt.Sprintf("Url [%d]", index), exceptions.Creating, exceptions.TryAgain))
			return
		}

		urls[index] = newUrl
	}

	password, passwordErr := db.Client.Password.Create().
		SetUser(authedUser).
		SetName(decodedName).
		SetNameIv(decodedNameIv).
		SetUsername(decodedUsername).
		SetUsernameIv(decodedUsernameIv).
		SetPassword(decodedPassword).
		SetPasswordIv(decodedPasswordIv).
		SetColour(input.Colour).
		AddAdditionalFields(additionalFields...).
		AddUrls(urls...).
		Save(db.Context)

	if passwordErr != nil {
		c.JSON(500, exceptions.Builder("password", exceptions.Creating, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"passwordId": password.ID.String()})
}

func Delete(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input DeleteInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, exceptions.Builder("body", exceptions.Invalid, exceptions.JsonOrXml))
		return
	}

	if input.PasswordId == "" {
		c.JSON(400, exceptions.Builder("passwordId", exceptions.MissingParam))
		return
	}

	decodedPasswordId, dpiErr := uuid.Parse(input.PasswordId)
	if dpiErr != nil {
		c.JSON(400, exceptions.Builder("passwordId", exceptions.ParsingParam, exceptions.Uuid))
		return
	}

	dpErr := db.DeleteUserPasswordViaId(authedUser, decodedPasswordId)
	if dpErr != nil {
		c.JSON(400, exceptions.Builder("password", exceptions.Deleting))
		return
	}

	c.JSON(200, gin.H{})
}
