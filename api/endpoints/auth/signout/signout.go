package signout

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/db"
	"PasswordManager/ent"
)

type DeleteInput struct{}

func Delete(c *gin.Context) {
	authedSession := c.MustGet("authedSession").(*ent.Session)

	sessionErr := db.Client.Session.DeleteOne(authedSession).Exec(db.Context)
	if sessionErr != nil {
		c.JSON(500, gin.H{"error": gin.H{"code": "errUnknown", "message": "An unknown error has occured. Please try again later."}})
		return
	}

	c.JSON(200, gin.H{})
}
