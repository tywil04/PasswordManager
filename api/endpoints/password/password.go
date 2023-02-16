package password

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/ent"
)

type GetInput struct {
	PasswordId string `form:"passwordId" json:"passwordId" xml:"passwordId"`
}

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
	}
	Urls []struct {
		Url   string `form:"url" json:"url" xml:"url"`
		UrlIv string `form:"urlIv" json:"urlIv" xml:"urlIv"`
	}
}

type DeleteInput struct {
	PasswordId string `form:"passwordId" json:"passwordId" xml:"passwordId"`
}

func Get(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input GetInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.PasswordId == "" {
		passwords, passwordsErr := db.GetUserPasswords(authedUser)
		if passwordsErr != nil {
			c.JSON(500, helpers.ErrorUnknown())
			return
		}

		jsonPasswords := make([]gin.H, len(passwords))
		for index, password := range passwords {
			additionalFields, afErr := db.GetPasswordAdditionalFields(password)
			if afErr != nil {
				c.JSON(500, helpers.ErrorUnknown())
				return
			}

			urls, uErr := db.GetPasswordUrls(password)
			if uErr != nil {
				c.JSON(500, helpers.ErrorUnknown())
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
	} else {
		decodedPasswordId, dpiErr := uuid.Parse(input.PasswordId)
		if dpiErr != nil {
			c.JSON(400, helpers.ErrorInvalid("passwordId"))
			return
		}

		password, passwordErr := db.GetUserPassword(authedUser, decodedPasswordId)
		if passwordErr != nil {
			c.JSON(400, helpers.ErrorInvalid("password"))
			return
		}

		additionalFields, afErr := db.GetPasswordAdditionalFields(password)
		if afErr != nil {
			c.JSON(500, helpers.ErrorUnknown())
			return
		}

		urls, uErr := db.GetPasswordUrls(password)
		if uErr != nil {
			c.JSON(500, helpers.ErrorUnknown())
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

		jsonPassword := gin.H{
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

		c.JSON(200, gin.H{"password": jsonPassword})
	}
}

func Post(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input PostInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.Name == "" {
		c.JSON(400, helpers.ErrorMissing("name"))
		return
	}

	if input.NameIv == "" {
		c.JSON(400, helpers.ErrorMissing("nameIv"))
		return
	}

	if input.Username == "" {
		c.JSON(400, helpers.ErrorMissing("username"))
		return
	}

	if input.UsernameIv == "" {
		c.JSON(400, helpers.ErrorMissing("usernameIv"))
		return
	}

	if input.Password == "" {
		c.JSON(400, helpers.ErrorMissing("password"))
		return
	}

	if input.PasswordIv == "" {
		c.JSON(400, helpers.ErrorMissing("passwordIv"))
		return
	}

	if input.Colour == "" {
		c.JSON(400, helpers.ErrorMissing("colour"))
		return
	}

	decodedName, dnErr := base64.StdEncoding.DecodeString(input.Name)
	if dnErr != nil {
		c.JSON(400, helpers.ErrorInvalid("name"))
		return
	}

	decodedNameIv, dniErr := base64.StdEncoding.DecodeString(input.NameIv)
	if dniErr != nil {
		c.JSON(400, helpers.ErrorInvalid("nameIv"))
		return
	}

	decodedUsername, duErr := base64.StdEncoding.DecodeString(input.Username)
	if duErr != nil {
		c.JSON(400, helpers.ErrorInvalid("username"))
		return
	}

	decodedUsernameIv, duiErr := base64.StdEncoding.DecodeString(input.UsernameIv)
	if duiErr != nil {
		c.JSON(400, helpers.ErrorInvalid("usernameIv"))
		return
	}

	decodedPassword, dpErr := base64.StdEncoding.DecodeString(input.Password)
	if dpErr != nil {
		c.JSON(400, helpers.ErrorInvalid("password"))
		return
	}

	decodedPasswordIv, dpiErr := base64.StdEncoding.DecodeString(input.PasswordIv)
	if dpiErr != nil {
		c.JSON(400, helpers.ErrorInvalid("passwordIv"))
		return
	}

	_, hexErr := strconv.ParseInt(input.Colour, 16, 64)
	if hexErr != nil {
		c.JSON(400, helpers.ErrorInvalid("colour"))
		return
	}

	additionalFields := make([]*ent.AdditionalField, len(input.AdditionalFields))
	for index, additionalField := range input.AdditionalFields {
		decodedKey, dkErr := base64.StdEncoding.DecodeString(additionalField.Key)
		if dkErr != nil {
			c.JSON(400, helpers.ErrorInvalid("additionalFields["+fmt.Sprint(index)+"].key"))
			return
		}

		decodedKeyIv, dkiErr := base64.StdEncoding.DecodeString(additionalField.KeyIv)
		if dkiErr != nil {
			c.JSON(400, helpers.ErrorInvalid("additionalFields["+fmt.Sprint(index)+"].keyIv"))
			return
		}

		decodedValue, dvErr := base64.StdEncoding.DecodeString(additionalField.Value)
		if dvErr != nil {
			c.JSON(400, helpers.ErrorInvalid("additionalFields["+fmt.Sprint(index)+"].value"))
			return
		}

		decodedValueIv, dviErr := base64.StdEncoding.DecodeString(additionalField.ValueIv)
		if dviErr != nil {
			c.JSON(400, helpers.ErrorInvalid("additionalFields["+fmt.Sprint(index)+"].valueIv"))
			return
		}

		newAdditionalField, nafErr := db.Client.AdditionalField.Create().
			SetKey(decodedKey).
			SetKeyIv(decodedKeyIv).
			SetValue(decodedValue).
			SetValueIv(decodedValueIv).
			Save(db.Context)

		if nafErr != nil {
			c.JSON(500, helpers.ErrorCreating("additionalFields['"+fmt.Sprint(index)+"]"))
			return
		}

		additionalFields[index] = newAdditionalField
	}

	urls := make([]*ent.Url, len(input.Urls))
	for index, url := range input.Urls {
		decodedUrl, duErr := base64.StdEncoding.DecodeString(url.Url)
		if duErr != nil {
			c.JSON(400, helpers.ErrorInvalid("urls['"+fmt.Sprint(index)+"].url"))
			return
		}

		decodedUrlIv, duiErr := base64.StdEncoding.DecodeString(url.UrlIv)
		if duiErr != nil {
			c.JSON(400, helpers.ErrorInvalid("urls['"+fmt.Sprint(index)+"].urlIv"))
			return
		}

		newUrl, nuErr := db.Client.Url.Create().
			SetURL(decodedUrl).
			SetUrlIv(decodedUrlIv).
			Save(db.Context)

		if nuErr != nil {
			c.JSON(500, helpers.ErrorCreating("urls['"+fmt.Sprint(index)+"]"))
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
		c.JSON(500, helpers.ErrorCreating("password"))
		return
	}

	c.JSON(200, gin.H{"passwordId": password.ID.String()})
}

func Delete(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input DeleteInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, helpers.ErrorInvalid("body"))
		return
	}

	if input.PasswordId == "" {
		c.JSON(400, helpers.ErrorMissing("passwordId"))
		return
	}

	decodedPasswordId, dpiErr := uuid.Parse(input.PasswordId)
	if dpiErr != nil {
		c.JSON(400, helpers.ErrorInvalid("passwordId"))
		return
	}

	dpErr := db.DeleteUserPasswordViaId(authedUser, decodedPasswordId)
	if dpErr != nil {
		c.JSON(400, helpers.ErrorDeleting("password"))
		return
	}

	c.JSON(200, gin.H{})
}
