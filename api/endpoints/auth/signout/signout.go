package signout

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/helpers"
	"PasswordManager/ent"
)

type DeleteInput struct{}

func Delete(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	authedSession := c.MustGet("authedSession").(*ent.Session)

	sessionErr := db.DeleteUserSession(authedUser, authedSession)
	if sessionErr != nil {
		c.JSON(500, helpers.ErrorDeleting("session"))
		return
	}

	c.JSON(200, gin.H{})
}
