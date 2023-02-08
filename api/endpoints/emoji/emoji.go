package emoji

import (
	"github.com/gin-gonic/gin"

	"PasswordManager/api/lib/emoji"
)

type GetInput struct{}

func Get(c *gin.Context) {
	c.JSON(200, gin.H{"emojis": emoji.Emojis})
}
