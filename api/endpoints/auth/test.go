package auth

import "github.com/gin-gonic/gin"

type GetTestInput struct{}

func GetTest(c *gin.Context) {
	c.JSON(200, gin.H{})
}
