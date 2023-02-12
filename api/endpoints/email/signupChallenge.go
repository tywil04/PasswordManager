package email

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
	"PasswordManager/ent/emailchallenge"
)

type PostSignupEmailChallengeInput struct {
	EmailChallengeId string `form:"emailChallengeId" json:"emailChallengeId" xml:"emailChallengeId"`
	Code             string `form:"code" json:"code" xml:"code"`
}

func PostSignupEmailChallenge(c *gin.Context) {
	var input PostSignupEmailChallengeInput

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

	foundChallenge, foundChallengeErr := db.GetEmailChallenge(decodedChallengeId)
	if foundChallengeErr != nil || foundChallenge.For != emailchallenge.ForSignup {
		c.JSON(400, gin.H{"error": gin.H{"code": "errEmailChallengeNotFound", "message": "Unable to find valid challenge using 'emailChallengeId'."}})
		return
	}

	foundUser, foundUserErr := db.GetEmailChallengeUser(foundChallenge)
	if foundUserErr != nil {
		c.JSON(400, gin.H{"error": gin.H{"code": "errEmailChallengeNotFound", "message": "Unable to find valid challenge using 'emailChallengeId'."}})
		return
	}

	sameCode := cryptography.ConstantTimeCompare([]byte(foundChallenge.Code), []byte(input.Code))
	if !sameCode {
		c.JSON(400, gin.H{"error": gin.H{"code": "errIncorrectEmailChallengeCode", "message": "Incorrect code for emailChallenge."}})
		return
	}

	foundUser.Update().SetVerified(true).Save(db.Context)
	db.Client.EmailChallenge.DeleteOne(foundChallenge).Exec(db.Context)

	c.JSON(200, gin.H{})
}
