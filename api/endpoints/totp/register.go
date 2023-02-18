package totp

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/ent"
)

type GetRegisterInput struct{}

type PostRegisterInput struct {
	TotpCredentialId string `form:"totpCredentialId" json:"totpCredentialId" xml:"totpCredentialId"`
	Code             string `form:"code" json:"code" xml:"code"`
}

func GetRegister(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	totpKey, totpErr := totp.Generate(totp.GenerateOpts{
		Issuer:      "Password Manager",
		AccountName: authedUser.Email,
		Algorithm:   otp.AlgorithmSHA512,
	})

	if totpErr != nil {
		c.JSON(500, exceptions.Builder("totpCredential", exceptions.Creating, exceptions.TryAgain))
		return
	}

	secret := totpKey.Secret()

	credential, credentialErr := db.Client.TotpCredential.Create().
		SetSecret(secret).
		SetUser(authedUser).
		Save(db.Context)

	if credentialErr != nil {
		c.JSON(500, exceptions.Builder("totpCredential", exceptions.Creating, exceptions.TryAgain))
		return
	}

	secretQr, sqErr := totpKey.Image(512, 512)
	if sqErr != nil {
		c.JSON(500, exceptions.Builder("totpSecretQr", exceptions.Issuing, exceptions.TryAgain))
		return
	}

	var buffer bytes.Buffer
	imageErr := jpeg.Encode(&buffer, secretQr, nil)
	if imageErr != nil {
		c.JSON(500, exceptions.Builder("totpSecretQr", exceptions.Issuing, exceptions.TryAgain))
		return
	}

	base64Image := base64.StdEncoding.EncodeToString(buffer.Bytes())

	c.JSON(200, gin.H{"totpCredentialId": credential.ID.String(), "totpSecret": secret, "totpSecretQr": "data:image/jpeg;base64," + base64Image})
}

func PostRegister(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	var input PostRegisterInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, exceptions.Builder("body", exceptions.Invalid, exceptions.JsonOrXml))
		return
	}

	if input.TotpCredentialId == "" {
		c.JSON(400, exceptions.Builder("totpCredentialId", exceptions.MissingParam))
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.TotpCredentialId)
	if dciErr != nil {
		c.JSON(400, exceptions.Builder("totpCredentialId", exceptions.ParsingParam, exceptions.Uuid))
		return
	}

	credential, credentialErr := db.GetUserTotpCredential(authedUser)
	if credentialErr != nil {
		c.JSON(400, exceptions.Builder("totpCredential", exceptions.Invalid))
		return
	}

	if credential.ID != decodedChallengeId {
		c.JSON(400, exceptions.Builder("totpCredential", exceptions.Invalid))
		return
	}

	valid := totp.Validate(input.Code, credential.Secret)
	if valid {
		credential.Update().SetValidated(true).Exec(db.Context)
		authedUser.Update().SetTotpEnabled(true).Exec(db.Context)
		c.JSON(200, gin.H{"totpCredentialId": credential.ID.String()})
	} else if !valid {
		db.Client.TotpCredential.DeleteOne(credential).Exec(db.Context)
		c.JSON(400, exceptions.Builder("code", exceptions.IncorrectChallenge))
	}
}
