package test

import "github.com/gin-gonic/gin"

type GetInput struct{}

func Get(c *gin.Context) {
	c.JSON(200, gin.H{})
}
