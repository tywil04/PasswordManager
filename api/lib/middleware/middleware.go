package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/cryptography"
	"PasswordManager/api/lib/db"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		parts := strings.Split(authToken, ";")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		sessionId := parts[0]
		encodedSalt := parts[1]
		signature := parts[2]

		decodedSessionId, dsiErr := base64.StdEncoding.DecodeString(sessionId)
		if dsiErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		decodedSignature, dsErr := base64.StdEncoding.DecodeString(signature)
		if dsErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		parsedDecodedSessionId, pdsiErr := uuid.Parse(string(decodedSessionId[:]))
		if pdsiErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		session, sessionErr := db.Client.Session.Get(db.Context, parsedDecodedSessionId)
		if sessionErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		if session.Expiry.Before(time.Now()) {
			db.Client.Session.DeleteOne(session).Exec(db.Context)
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		user, userErr := session.QueryUser().First(db.Context)
		if userErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
			c.Abort()
			return
		}

		publicKey := cryptography.ImportPublicKey(session.N, session.E)
		valid := cryptography.VerifySignature(publicKey, decodedSignature, user.Email+base64.StdEncoding.EncodeToString(user.StrengthenedMasterHash)+encodedSalt)

		if valid {
			c.Set("authedUser", user)
			c.Set("authedSession", session)
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "errAuthRequired", "message": "Failed to authenticate user, invalid 'authToken' found in 'Authorization' header."}})
		c.Abort()
	}
}
