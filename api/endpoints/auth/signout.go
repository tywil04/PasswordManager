package auth

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/ent"
)

type DeleteSignoutInput struct{}

func DeleteSignout(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	authedSession := c.MustGet("authedSession").(*ent.Session)

	sessionErr := db.DeleteUserSession(authedUser, authedSession)
	if sessionErr != nil {
		c.JSON(500, exceptions.Builder("session", exceptions.Deleting, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{})
}
