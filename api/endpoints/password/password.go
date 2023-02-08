package password

import (
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/emoji"
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
	Emoji            string `form:"emoji" json:"emoji" xml:"emoji"`
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
	var input DeleteInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.PasswordId == "" {
		passwords, passwordsErr := db.Client.Password.Query().All(db.Context)
		if passwordsErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
			return
		}

		jsonPasswords := make([]gin.H, len(passwords))
		for index, password := range passwords {
			additionalFields, afErr := password.QueryAdditionalFields().All(db.Context)
			if afErr != nil {
				c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
				return
			}

			urls, uErr := password.QueryUrls().All(db.Context)
			if uErr != nil {
				c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
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
				"emoji":            password.Emoji,
				"additionalFields": jsonAdditionalFields,
				"urls":             jsonUrls,
			}
		}

		c.JSON(200, gin.H{"passwords": jsonPasswords})
	} else {
		decodedPasswordId, dpiErr := uuid.Parse(input.PasswordId)
		if dpiErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errParsingPasswordId", "message": "Unable to parse 'passwordId', expected uuid."}})
			return
		}

		password, passwordErr := db.Client.Password.Get(db.Context, decodedPasswordId)
		if passwordErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errPasswordNotFound", "message": "Unable to find valid password using 'passwordId'."}})
			return
		}

		additionalFields, afErr := password.QueryAdditionalFields().All(db.Context)
		if afErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
			return
		}

		urls, uErr := password.QueryUrls().All(db.Context)
		if uErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
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
			"emoji":            password.Emoji,
			"additionalFields": jsonAdditionalFields,
			"urls":             jsonUrls,
		}

		c.JSON(200, gin.H{"password": jsonPassword})
	}
}

func Post(c *gin.Context) {
	var input PostInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.Name == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingName", "message": "Required 'name' was not found."}})
		return
	}

	if input.NameIv == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingNameIv", "message": "Required 'nameIv' was not found."}})
		return
	}

	if input.Username == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingUsername", "message": "Required 'username' was not found."}})
		return
	}

	if input.UsernameIv == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingUsernameIv", "message": "Required 'usernameIv' was not found."}})
		return
	}

	if input.Password == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingPassword", "message": "Required 'password' was not found."}})
		return
	}

	if input.PasswordIv == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingPasswordIv", "message": "Required 'passwordIv' was not found."}})
		return
	}

	if input.Emoji == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingEmoji", "message": "Required 'emoji' was not found."}})
		return
	}

	decodedName, dnErr := base64.StdEncoding.DecodeString(input.Name)
	if dnErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingUsername", "message": "Unable to parse 'username', expected base64 encoding."}})
		return
	}

	decodedNameIv, dniErr := base64.StdEncoding.DecodeString(input.NameIv)
	if dniErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingUsername", "message": "Unable to parse 'username', expected base64 encoding."}})
		return
	}

	decodedUsername, duErr := base64.StdEncoding.DecodeString(input.Username)
	if duErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingUsername", "message": "Unable to parse 'username', expected base64 encoding."}})
		return
	}

	decodedUsernameIv, duiErr := base64.StdEncoding.DecodeString(input.UsernameIv)
	if duiErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingUsernameIv", "message": "Unable to parse 'usernameIv', expected base64 encoding."}})
		return
	}

	decodedPassword, dpErr := base64.StdEncoding.DecodeString(input.Password)
	if dpErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingPassword", "message": "Unable to parse 'password', expected base64 encoding."}})
		return
	}

	decodedPasswordIv, dpiErr := base64.StdEncoding.DecodeString(input.PasswordIv)
	if dpiErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingPasswordIv", "message": "Unable to parse 'passwordIv', expected base64 encoding."}})
		return
	}

	if !emoji.IsEmoji(input.Emoji) {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingEmoji", "message": "Unable to parse 'emoji', expected unicode emoji."}})
		return
	}

	additionalFields := make([]*ent.AdditionalField, len(input.AdditionalFields))
	for index, additionalField := range input.AdditionalFields {
		decodedKey, dkErr := base64.StdEncoding.DecodeString(additionalField.Key)
		if dkErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errPassingAdditionalFieldKey", "message": "Unable to parse 'additionalFields['" + fmt.Sprint(index) + "].key', expected base64 encoding."}})
			return
		}

		decodedKeyIv, dkiErr := base64.StdEncoding.DecodeString(additionalField.KeyIv)
		if dkiErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errPassingAdditionalFieldKeyIv", "message": "Unable to parse 'additionalFields['" + fmt.Sprint(index) + "].keyIv', expected base64 encoding."}})
			return
		}

		decodedValue, dvErr := base64.StdEncoding.DecodeString(additionalField.Value)
		if dvErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errPassingAdditionalFieldValue", "message": "Unable to parse 'additionalFields['" + fmt.Sprint(index) + "].value', expected base64 encoding."}})
			return
		}

		decodedValueIv, dviErr := base64.StdEncoding.DecodeString(additionalField.ValueIv)
		if dviErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errPassingAdditionalFieldValueIv", "message": "Unable to parse 'additionalFields['" + fmt.Sprint(index) + "].valueIv', expected base64 encoding."}})
			return
		}

		newAdditionalField, nafErr := db.Client.AdditionalField.Create().
			SetKey(decodedKey).
			SetKeyIv(decodedKeyIv).
			SetValue(decodedValue).
			SetValueIv(decodedValueIv).
			Save(db.Context)

		if nafErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
			return
		}

		additionalFields[index] = newAdditionalField
	}

	urls := make([]*ent.Url, len(input.Urls))
	for index, url := range input.Urls {
		decodedUrl, duErr := base64.StdEncoding.DecodeString(url.Url)
		if duErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errPassingUrl", "message": "Unable to parse 'urls['" + fmt.Sprint(index) + "].url', expected base64 encoding."}})
			return
		}

		decodedUrlIv, duiErr := base64.StdEncoding.DecodeString(url.UrlIv)
		if duiErr != nil {
			c.JSON(400, gin.H{"error": gin.H{"code": "errPassingUrlIv", "message": "Unable to parse 'urls['" + fmt.Sprint(index) + "].urlIv', expected base64 encoding."}})
			return
		}

		newUrl, nuErr := db.Client.Url.Create().
			SetURL(decodedUrl).
			SetUrlIv(decodedUrlIv).
			Save(db.Context)

		if nuErr != nil {
			c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
			return
		}

		urls[index] = newUrl
	}

	password, passwordErr := db.Client.Password.Create().
		SetName(decodedName).
		SetNameIv(decodedNameIv).
		SetUsername(decodedUsername).
		SetUsernameIv(decodedUsernameIv).
		SetPassword(decodedPassword).
		SetPasswordIv(decodedPasswordIv).
		SetEmoji(input.Emoji).
		AddAdditionalFields(additionalFields...).
		AddUrls(urls...).
		Save(db.Context)

	if passwordErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	c.JSON(200, gin.H{"passwordId": password.ID.String()})
}

func Delete(c *gin.Context) {
	var input DeleteInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.PasswordId == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingPasswordId", "message": "Required 'passwordId' was not found."}})
		return
	}

	decodedPasswordId, dpiErr := uuid.Parse(input.PasswordId)
	if dpiErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingPasswordId", "message": "Unable to parse 'passwordId', expected uuid."}})
		return
	}

	password, passwordErr := db.Client.Password.Get(db.Context, decodedPasswordId)
	if passwordErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errPasswordNotFound", "message": "Unable to find valid password using 'passwordId'."}})
		return
	}

	dpErr := db.Client.Password.DeleteOne(password).Exec(db.Context)
	if dpErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	c.JSON(200, gin.H{})
}
