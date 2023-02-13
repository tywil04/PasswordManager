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
	"PasswordManager/ent"
	"PasswordManager/ent/user"
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
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	secret := totpKey.Secret()

	credential, credentialErr := db.Client.TotpCredential.Create().
		SetSecret(secret).
		SetUser(authedUser).
		Save(db.Context)

	if credentialErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	secretQr, sqErr := totpKey.Image(512, 512)
	if sqErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	var buffer bytes.Buffer
	imageErr := jpeg.Encode(&buffer, secretQr, nil)
	if imageErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
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
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.TotpCredentialId == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingTotpCredentialId", "message": "Required 'totpCredentialId' was not found."}})
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.TotpCredentialId)
	if dciErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingTotpCredentialId", "message": "Unable to parse 'totpCredentialId', expected uuid."}})
		return
	}

	credential, credentialErr := db.GetUserTotpCredential(authedUser)
	if credentialErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errTotpCredentialNotFound", "message": "Unable to find valid totp credential using 'totpCredentialId'."}})
		return
	}

	if credential.ID != decodedChallengeId {
		c.JSON(400, gin.H{"error": gin.H{"code": "errTotpCredentialNotFound", "message": "Unable to find valid totp credential using 'totpCredentialId'."}})
		return
	}

	valid := totp.Validate(input.Code, credential.Secret)
	if valid {
		credential.Update().SetValidated(true).Exec(db.Context)

		if authedUser.Default2FA != user.Default2FAWebauthn {
			authedUser.Update().SetDefault2FA(user.Default2FATotp).Exec(db.Context)
		}

		c.JSON(200, gin.H{"totpCredentialId": credential.ID.String()})
	} else if !valid {
		db.Client.TotpCredential.DeleteOne(credential).Exec(db.Context)

		if credential.ID != decodedChallengeId {
			c.JSON(400, gin.H{"error": gin.H{"code": "errIncorrectTotpChallengeCode", "message": "Incorrect code for totpChallenge."}})
			return
		}
	}
}
