package email

import (
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/smtp"
	"PasswordManager/ent/emailchallenge"
)

type PostSigninEmailChallengeInput struct {
	EmailChallengeId string `form:"emailChallengeId" json:"emailChallengeId" xml:"emailChallengeId"`
	Code             string `form:"code" json:"code" xml:"code"`
}

func PostSigninEmailChallenge(c *gin.Context) {
	var input PostSigninEmailChallengeInput

	bindingErr := c.Bind(&input)
	if bindingErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingBody", "message": "Unable to parse body, expected json structure."}})
		return
	}

	if input.EmailChallengeId == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingEmailChallengeId", "message": "Required 'emailChallengeId' was not found."}})
		return
	}

	if input.Code == "" {
		c.JSON(400, gin.H{"error": gin.H{"code": "errMissingCode", "message": "Required 'code' was not found."}})
		return
	}

	decodedChallengeId, dciErr := uuid.Parse(input.EmailChallengeId)
	if dciErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errParsingEmailChallengeId", "message": "Unable to parse 'emailChallengeId', expected uuid."}})
		return
	}

	foundChallenge, foundChallengeErr := db.Client.EmailChallenge.Get(db.Context, decodedChallengeId)
	if foundChallengeErr != nil || foundChallenge.For != emailchallenge.ForSignin {
		c.JSON(400, gin.H{"error": gin.H{"code": "errEmailChallengeNotFound", "message": "Unable to find valid challenge using 'emailChallengeId'."}})
		return
	}

	foundUser, foundUserErr := foundChallenge.QueryUser().Unique(true).First(db.Context)
	if foundUserErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errEmailChallengeNotFound", "message": "Unable to find valid challenge using 'emailChallengeId'."}})
		return
	}

	sameCode := cryptography.ConstantTimeCompare([]byte(foundChallenge.Code), []byte(input.Code))
	if !sameCode {
		c.JSON(400, gin.H{"error": gin.H{"code": "errIncorrectEmailChallengeCode", "message": "Incorrect code for emailChallenge."}})
		go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Unsuccessful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: false})
		return
	}

	go smtp.SendTemplate(foundUser.Email, "PasswordManager5 Successful Sign In Notification", smtp.SigninNotificationTemplate, smtp.SigninNotificationTemplateData{Successful: true})

	salt := cryptography.RandomBytes(32)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	publicKey, signature := cryptography.GenerateSignature(foundUser.Email + base64.StdEncoding.EncodeToString(foundUser.StrengthenedMasterHash) + encodedSalt)

	session, sessionErr := db.Client.Session.Create().
		SetUser(foundUser).
		SetN(publicKey.N.Bytes()).
		SetE(publicKey.E).
		SetExpiry(time.Now().Add(time.Hour)).
		Save(db.Context)

	if sessionErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	token := base64.StdEncoding.EncodeToString([]byte(session.ID.String())) + ";" + encodedSalt + ";" + base64.StdEncoding.EncodeToString(signature)

	db.Client.EmailChallenge.DeleteOne(foundChallenge).Exec(db.Context)

	encodedProtectedDatabaseKey := base64.StdEncoding.EncodeToString(foundUser.ProtectedDatabaseKey)
	encodedProtectedDatabaseKeyIv := base64.StdEncoding.EncodeToString(foundUser.ProtectedDatabaseKeyIv)

	c.JSON(200, gin.H{"authToken": token, "protectedDatabaseKey": encodedProtectedDatabaseKey, "protectedDatabaseKeyIv": encodedProtectedDatabaseKeyIv})
}
